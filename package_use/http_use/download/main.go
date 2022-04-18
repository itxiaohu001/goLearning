package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	download("https://p3-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/41eae12520064636ba3ccc82a77e3fc8~tplv-k3u1fbpfcp-zoom-crop-mark:1304:1304:1304:734.awebp?", "picture.png")
}
func download(url, filename string) {
	r, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer func() { _ = r.Body.Close() }()
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer func() { _ = f.Close() }()
	n, err := io.Copy(f, r.Body)
	fmt.Println(n, err)
}
