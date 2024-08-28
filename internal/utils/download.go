package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// DownloadPackage downloads a package from a repository.
func DownloadPackage(packageName string) error {
	url := fmt.Sprintf("http://localhost:8080/packages/%s.tar.gz", packageName)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(packageName + ".tar.gz")
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}
