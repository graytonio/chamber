package main

import (
	"fmt"
	"os"

	"github.com/graytonio/chamber/pkg/configparser"
)

var home, _ = os.UserHomeDir()

// var path = fmt.Sprintf("%s/.chambercfg", home)
var path = "home/graytonio/.chambercfg"

func main() {
	config, err := configparser.ReadConfigFile(os.DirFS("/"), path)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	for _, template := range config.Templates {
		fmt.Printf("%s/%s\n", template.Repo.Url, template.Template)
	}
}
