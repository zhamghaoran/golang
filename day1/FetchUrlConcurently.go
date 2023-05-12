package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		if strings.HasPrefix(url, "http://") == false {
			url = "http://" + url
		}
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch) // 遍历通道里面的数据
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) { // 在函数的参数里面指定了通道的方向
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		err.Error()
	}
	nbytes, err := io.Copy(io.Discard, resp.Body) // 因为我们不想要具体的内容，我们将数据复制到Discard输出流当中并返回一个字节数
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s : %v", url, err) // 向通道里面写入数据
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)

}
