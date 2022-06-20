package utils

import (
	"archive/tar"
	"io"
	"io/fs"
)

func ListTarContents(fs fs.FS, path string) ([]string, error) {
	tarFile, err := fs.Open(path)
	if err != nil {
		return nil, err
	}
	defer tarFile.Close()

	tr := tar.NewReader(tarFile)
	results := []string{}
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		results = append(results, hdr.Name)
	}
	return results, nil
}
