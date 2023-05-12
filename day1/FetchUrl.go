package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			err.Error()
			os.Exit(1)
		}
		b, err := io.ReadAll(resp.Body) // 这里的b是字节数组，在下面输出的时候记得使用string函数转成string
		if err != nil {
			err.Error()
		}
		fmt.Println(string(b))
		fmt.Println(resp.TransferEncoding)
		resp.Body.Close()
	}
}
