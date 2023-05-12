package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if strings.HasPrefix(url, "http://") == false {
			url = "http://" + url
		}
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
