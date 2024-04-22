package main

import "github.com/begonia-org/go-sdk/example"

func main() {
	go example.RunPlugins("127.0.0.1:21216")
	go example.RunPlugins("127.0.0.1:21217")
	example.Run("0.0.0.0:29527")
}