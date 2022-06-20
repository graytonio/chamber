package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/graytonio/chamber/pkg/configparser"
	"github.com/graytonio/chamber/pkg/utils"
)

var home, _ = os.UserHomeDir()

// var path = fmt.Sprintf("%s/.chambercfg", home)
var path = ".chambercfg"

func main() {
	home_dir, _ := os.UserHomeDir()
	home_fs := os.DirFS(home_dir)
	config, err := configparser.ReadConfigFile(home_fs, path)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	tar_url := fmt.Sprintf("%s/tarball/%s", config.Templates[0].Repo.Url, config.Templates[0].Repo.Branch)
	tar_ball, download_err := utils.DownloadFile(tar_url, "*.tar.gz")
	if download_err != nil {
		fmt.Println(download_err.Error())
		os.Exit(1)
	}

	tmp_fs := os.DirFS(os.TempDir())
	templates, tar_error := utils.ListTarContents(tmp_fs, strings.ReplaceAll(tar_ball.Name(), os.TempDir(), "")[1:])
	if tar_error != nil {
		fmt.Println(tar_error.Error())
		os.Exit(1)
	}

	fmt.Println(templates)
}
