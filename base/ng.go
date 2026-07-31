package base

import (
	"fmt"
	"io"
	h "net/http"
)

func main() {
	resp, err := h.Get("https://httpbin.org/uuid")
	if err != nil {
		fmt.Println("HTTP error:", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("Status:", resp.Status)
	fmt.Println("Content-Type:", resp.Header.Get("Content-Type"))

	body, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		fmt.Println("Read error:", readErr)
		return
	}
	previewLen := 120
	if len(body) < previewLen {
		previewLen = len(body)
	}
	fmt.Printf("Body preview (%d bytes): %s\n",
		previewLen, string(body[:previewLen]))
}
