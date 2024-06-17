package client

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"sync"
	"time"

	api "github.com/begonia-org/go-sdk/api/file/v1"
	common "github.com/begonia-org/go-sdk/common/api/v1"
)

type FilesAPI struct {
	*BaseAPI
	engine api.FileEngine
}

type UploadFileAPIResponse struct {
	*Response
	*api.UploadFileResponse
}
type AbortUploadAPIResponse struct {
	*Response
	*api.AbortMultipartUploadResponse
}
type UploadPartAPIResponse struct {
	*Response
	*api.UploadMultipartFileResponse
}
type UploadCompleteAPIResponse struct {
	*Response
	*api.CompleteMultipartUploadResponse
}
type FileMetadataAPIResponse struct {
	*Response
	*api.FileMetadataResponse
}
type FileListAPIResponse struct {
	*Response
	*api.ListFilesResponse
}
type CreateBucketResponse struct {
	*Response
	*api.MakeBucketResponse
}

func NewFilesAPI(addr, accessKey, secretKey string, engine api.FileEngine) *FilesAPI {
	return &FilesAPI{
		BaseAPI: NewAPIClient(addr, accessKey, secretKey),
		engine:  engine,
	}
}

// UploadFile upload file to server
func (f *FilesAPI) UploadFile(ctx context.Context, srcPath string, dst string, bucket string, useVersion bool) (*UploadFileAPIResponse, error) {

	data, err := os.ReadFile(srcPath)
	if err != nil {
		return nil, err

	}
	shaer := sha256.New()
	shaer.Write(data)
	sha := shaer.Sum(nil)
	hashStr := fmt.Sprintf("%x", sha)
	content := api.UploadFileRequest{
		Key:         dst,
		Content:     data,
		Sha256:      hashStr,
		ContentType: http.DetectContentType(data),
		Engine:      f.engine.String(),
		Bucket:      bucket,
	}
	payload, err := json.Marshal(&content)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal data to JSON: %w", err)

	}

	rsp, err := f.Post(ctx, UPLOAD_API, nil, bytes.NewReader(payload))
	if err != nil {
		return nil, fmt.Errorf("Failed to send request: %w", err)

	}
	if rsp != nil {
		apiRsp := &api.UploadFileResponse{}
		resp, err := f.unmarshal(rsp, apiRsp)
		if err != nil {
			return nil, err
		}
		return &UploadFileAPIResponse{
			Response:           resp,
			UploadFileResponse: apiRsp,
		}, nil

	}
	return nil, nil
}
func (f *FilesAPI) AbortUpload(ctx context.Context, uploadId string) (*AbortUploadAPIResponse, error) {
	content := &api.AbortMultipartUploadRequest{
		UploadId: uploadId,
		Engine:   f.engine.String(),
	}
	payload, err := json.Marshal(&content)
	if err != nil {
		return nil, err
	}
	rsp, err := f.Post(ctx, ABORT_PART_API, nil, bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}
	if rsp != nil {
		apiRsp := &api.AbortMultipartUploadResponse{}
		resp, err := f.unmarshal(rsp, apiRsp)
		if err != nil {
			return nil, err
		}
		return &AbortUploadAPIResponse{Response: resp, AbortMultipartUploadResponse: apiRsp}, nil
	}
	return nil, nil
}
func (f *FilesAPI) UploadPart(ctx context.Context, data []byte, key string, partNumber int, uploadId string) (*UploadPartAPIResponse, error) {
	shaer := sha256.New()
	shaer.Write(data)
	sha := shaer.Sum(nil)
	hashStr := fmt.Sprintf("%x", sha)
	content := &api.UploadMultipartFileRequest{
		Key:        key,
		PartNumber: int64(partNumber),
		UploadId:   uploadId,
		Content:    data,
		Sha256:     hashStr,
		Engine:     f.engine.String(),
	}
	var err error
	defer func(ctx context.Context) {
		if err != nil {
			_, _ = f.AbortUpload(ctx, uploadId)
		}
	}(ctx)
	payload, err := json.Marshal(&content)
	if err != nil {
		return nil, err
	}
	rsp, err := f.Put(ctx, UPLOAD_PART_API, nil, bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}
	if rsp != nil {
		apiRsp := &api.UploadMultipartFileResponse{}
		resp, err := f.unmarshal(rsp, apiRsp)
		if err != nil {
			return nil, err
		}
		return &UploadPartAPIResponse{Response: resp, UploadMultipartFileResponse: apiRsp}, nil
	}
	return nil, nil
}
func (f *FilesAPI) CompleteUpload(ctx context.Context, key string, uploadId string, sha256 string, bucket string, useVersion bool) (*UploadCompleteAPIResponse, error) {
	req := &api.CompleteMultipartUploadRequest{
		UploadId:   uploadId,
		Key:        key,
		Sha256:     sha256,
		Bucket:     bucket,
		Engine:     f.engine.String(),
	}
	payload, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	rsp, err := f.Post(ctx, COMPLETE_PART_API, nil, bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}
	if rsp != nil {
		apiRsp := &api.CompleteMultipartUploadResponse{}
		resp, err := f.unmarshal(rsp, apiRsp)
		if err != nil {
			return nil, err
		}
		return &UploadCompleteAPIResponse{Response: resp, CompleteMultipartUploadResponse: apiRsp}, nil

	}
	return nil, nil
}
func (f *FilesAPI) UploadFileWithMuiltParts(ctx context.Context, src string, key string, bucket string, useVersion bool) (*UploadCompleteAPIResponse, error) {
	initReq := &api.InitiateMultipartUploadRequest{
		Key:    key,
		Engine: f.engine.String(),
	}
	payload, err := json.Marshal(initReq)
	if err != nil {
		return nil, err
	}
	rsp, err := f.Post(ctx, INIT_PART_API, nil, bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}
	if rsp != nil {
		apiRsp := &api.InitiateMultipartUploadResponse{}
		resp, err := f.unmarshal(rsp, apiRsp)
		if err != nil {
			return nil, err
		}
		if common.Code(resp.StatusCode) != common.Code_OK {
			return nil, fmt.Errorf("Failed to initiate multipart upload: %s", resp.Err.Error())
		}
		info, err := os.Stat(src)
		if err != nil {
			return nil, fmt.Errorf("Failed to get file info: %w", err)
		}

		partSize := int64(2 * 1024 * 1024)
		partCount := math.Ceil(float64(info.Size()) / float64(partSize))
		batchSize := 0
		wg := &sync.WaitGroup{}
		file, err := os.Open(src)
		if err != nil {
			return nil, fmt.Errorf("Failed to open file: %w", err)
		}
		defer file.Close()
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		sha := sha256.New()
		for i := 0; i < int(partCount); i++ {
			buffer := make([]byte, partSize)
			n, err := file.Read(buffer)
			sha.Write(buffer)
			if err != nil && err != io.EOF {
				return nil, fmt.Errorf("Failed to read file: %w", err)
			}
			if n == 0 {
				break
			}
			wg.Add(1)
			go func(ctx context.Context, wg *sync.WaitGroup, uploadId string, partNumber int, data []byte) {
				defer wg.Done()
				rsp, err := f.UploadPart(ctx, data, key, partNumber, uploadId)
				if err != nil || rsp == nil {
					cancel()
					return
				}
				if common.Code(rsp.StatusCode) != common.Code_OK {
					cancel()
					return
				}

			}(ctx, wg, apiRsp.UploadId, i+1, buffer)
			batchSize++
			if batchSize == 10 {
				wg.Wait()
				batchSize = 0
			}

		}
		if batchSize > 0 {
			wg.Wait()
		}

		// completeReq := &api.CompleteMultipartUploadRequest{}
		// return apiRsp,nil
		hexStr := fmt.Sprintf("%x", sha.Sum(nil))
		return f.CompleteUpload(ctx, key, apiRsp.UploadId, hexStr, bucket, useVersion)
	}
	return nil, nil
}
func IsSnowflakeID(id string) bool {
	// 检查字符串长度
	if len(id) < 18 || len(id) > 19 {
		return false
	}

	// 尝试将字符串转换为整数
	_, err := strconv.ParseInt(id, 10, 64)
	return err == nil
}
func (f *FilesAPI) DownloadFile(ctx context.Context, key string, dst string, version string, bucket string) (string, error) {
	// uri := fmt.Sprintf("%s?key=%s&version=%s", Download_API, key, version)
	values := url.Values{}
	if !IsSnowflakeID(key) {
		values.Add("key", key)
	}else{
		values.Add("file_id", key)
	}
	values.Add("bucket", bucket)
	values.Add("engine", f.engine.String())
	if version != "" {
		values.Add("version", version)
	}
	uri := fmt.Sprintf("%s?%s", Download_API, values.Encode())
	headers := make(map[string]string)
	headers["accept"] = "application/octet-stream"
	rsp, err := f.Get(ctx, uri, headers)
	if err != nil {
		return "", err
	}
	if rsp != nil {
		if rsp.StatusCode != http.StatusOK {
			err := "unknown error"
			if rsp.StatusCode == http.StatusNotFound {
				err = "file not found"
			}
			return "", fmt.Errorf("Failed to download file: %s", err)
		}
		defer rsp.Body.Close()
		file, err := os.Create(dst)
		if err != nil {
			return "", err
		}
		defer file.Close()
		_, err = io.Copy(file, rsp.Body)
		if err != nil {
			return "", err
		}
		sha256Str := rsp.Header.Get("x-file-sha256")
		return sha256Str, nil
	}
	return "", fmt.Errorf("Failed to download file")
}

func (f *FilesAPI) RangeDownload(ctx context.Context, key string, version string, start, end int64, bucket string) ([]byte, error) {
	values := url.Values{}
	values.Add("key", key)
	values.Add("bucket", bucket)
	values.Add("engine", f.engine.String())
	if version != "" {
		values.Add("version", version)

	}
	uri := fmt.Sprintf("%s?%s", Download_PART_API, values.Encode())
	headers := make(map[string]string)
	headers["accept"] = "application/octet-stream"
	rangeHeader := "bytes="
	if start >= 0 {
		rangeHeader += fmt.Sprintf("%d-", start)
	}
	if end >= 0 {
		if start >= 0 {
			rangeHeader += fmt.Sprintf("%d", end)

		} else {
			rangeHeader += fmt.Sprintf("-%d", end)

		}

	}
	headers["range"] = rangeHeader
	rsp, err := f.Get(ctx, uri, headers)
	if err != nil {
		return nil, fmt.Errorf("Failed to send request: %w", err)
	}
	buf := new(bytes.Buffer)
	if rsp != nil {
		if rsp.StatusCode >= http.StatusBadRequest {
			err := "unknown error"
			if rsp.StatusCode == http.StatusNotFound {
				err = "file not found"
			}
			return nil, fmt.Errorf("Failed to download file: %s,with status code %d", err, rsp.StatusCode)
		}
		defer rsp.Body.Close()
		_, err := io.Copy(buf, rsp.Body)
		if err != nil {
			return nil, fmt.Errorf("Failed to read response: %w", err)
		}
		return buf.Bytes(), nil
	}
	return nil, fmt.Errorf("Failed to download file")
}
func (f *FilesAPI) Metadata(ctx context.Context, key string, version string, bucket string) (*FileMetadataAPIResponse, error) {
	// uri := fmt.Sprintf("%s?key=%s&version=%s", Download_API, key, version)
	values := url.Values{}
	values.Add("key", key)
	values.Add("bucket", bucket)
	values.Add("engine", f.engine.String())
	if version != "" {
		values.Add("version", version)

	}
	uri := fmt.Sprintf("%s?%s", Metadata_API, values.Encode())
	rsp, err := f.Head(ctx, uri, nil)
	if err != nil {
		return nil, err
	}
	if rsp != nil {
		if rsp.StatusCode != http.StatusOK {
			err := "unknown error"
			if rsp.StatusCode == http.StatusNotFound {
				err = "file not found"
			}
			return nil, fmt.Errorf("Failed to get file metadata: %s", err)
		}
		defer rsp.Body.Close()

		modfiyTime, _ := time.Parse(time.RFC1123, rsp.Header.Get("Last-Modified"))

		apiRsp := &api.FileMetadataResponse{
			Name:        rsp.Header.Get("X-File-Name"),
			ContentType: rsp.Header.Get("content-type"),
			Etag:        rsp.Header.Get("Etag"),
			ModifyTime:  modfiyTime.Unix(),
			Sha256:      rsp.Header.Get("X-File-Sha256"),
			Size:        rsp.ContentLength,
			Key:         key,
			Version:     rsp.Header.Get("X-File-Version"),
		}
		return &FileMetadataAPIResponse{
			Response:             &Response{StatusCode: rsp.StatusCode, RequestId: rsp.Header.Get("X-Request-Id")},
			FileMetadataResponse: apiRsp,
		}, nil
	}
	return nil, fmt.Errorf("Failed to get file metadata")
}
func (f *FilesAPI) DownloadMultiParts(ctx context.Context, key string, dst string, version string, bucket string) (*FileMetadataAPIResponse, error) {
	metadata, err := f.Metadata(ctx, key, version, bucket)
	if err != nil {
		return nil, fmt.Errorf("Failed to get file metadata: %w", err)
	}
	if metadata == nil {
		return nil, fmt.Errorf("Failed to get file metadata")
	}
	partSize := int64(2 * 1024 * 1024)
	partCount := math.Ceil(float64(metadata.Size) / float64(partSize))
	file, err := os.Create(dst)
	if err != nil {
		return nil, fmt.Errorf("Failed to create file: %w", err)
	}
	defer file.Close()
	for i := 0; i < int(partCount); i++ {
		rangeStartAt := int64(i) * partSize
		rangeEndAt := rangeStartAt + partSize - 1
		if rangeEndAt > metadata.Size {
			rangeEndAt = metadata.Size - 1
		}
		data, err := f.RangeDownload(ctx, key, version, rangeStartAt, rangeEndAt, bucket)
		if err != nil {
			return nil, fmt.Errorf("Failed to download part: %w", err)
		}
		_, err = file.WriteAt(data, rangeStartAt)
		if err != nil {
			return nil, fmt.Errorf("Failed to write data to file: %w", err)
		}

	}
	return metadata, nil
}

func (f *FilesAPI) DeleteFile(ctx context.Context, key string, bucket string) (*Response, error) {
	values := url.Values{}
	values.Add("key", key)
	values.Add("bucket", bucket)
	values.Add("engine", f.engine.String())
	uri := fmt.Sprintf("%s?%s", FILE_API, values.Encode())
	rsp, err := f.Delete(ctx, uri, nil, nil)
	if err != nil {
		return nil, err
	}
	if rsp != nil {
		apiRsp := &api.DeleteResponse{}
		resp, err := f.unmarshal(rsp, apiRsp)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
	return nil, fmt.Errorf("Failed to delete file")
}
func (f *FilesAPI) CreateBucket(ctx context.Context, bucket string, region string, objectLocking,enableVersion bool) (*CreateBucketResponse, error) {
	content := &api.MakeBucketRequest{
		Bucket:        bucket,
		Engine:        f.engine.String(),
		Region:        region,
		ObjectLocking: objectLocking,
		EnableVersion: enableVersion,
	}
	payload, err := json.Marshal(content)
	if err != nil {
		return nil, err
	}
	rsp, err := f.Post(ctx, BUCKET_API, nil, bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}
	if rsp != nil {
		apiRsp := &api.MakeBucketResponse{}
		resp, err := f.unmarshal(rsp, apiRsp)
		if err != nil {
			return nil, err
		}
		return &CreateBucketResponse{Response: resp, MakeBucketResponse: apiRsp}, nil
	}
	return nil, fmt.Errorf("Failed to delete files")
}
func (f *FilesAPI) ListFiles(ctx context.Context, engine, bucket string, page int32, pageSize int32) (*FileListAPIResponse, error) {
	values := url.Values{}
	if engine != "" {
		values.Add("engine", engine)
	}
	if bucket != "" {
		values.Add("bucket", bucket)

	}
	values.Add("page", fmt.Sprintf("%d", page))
	values.Add("page_size", fmt.Sprintf("%d", pageSize))
	uri := fmt.Sprintf("%s?%s", FILE_LIST_API, values.Encode())
	rsp, err := f.Get(ctx, uri, nil)
	if err != nil {
		return nil, err
	}
	if rsp != nil {
		apiRsp := &api.ListFilesResponse{}
		resp, err := f.unmarshal(rsp, apiRsp)
		if err != nil {
			return nil, err
		}
		return &FileListAPIResponse{
			Response:          resp,
			ListFilesResponse: apiRsp,
		}, nil
	}
	return nil, fmt.Errorf("Failed to list files")
}
