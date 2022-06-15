package utils

import (
	"io"
	"net/http"
	"os"
)

func DownloadFile(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	out, err := os.CreateTemp("", "*")
	if err != nil {
		return nil, err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return nil, err
	}

	content, err := os.ReadFile(out.Name())
	return content, err
}
