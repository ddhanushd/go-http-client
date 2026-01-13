package client_test

import (
	"path/filepath"

	"github.com/therewardstore/httpmatter"
)

func init() {
	_ = httpmatter.Init(&httpmatter.Config{
		BaseDir:       filepath.Join("..", "testdata"),
		FileExtension: ".http",
	})
}
