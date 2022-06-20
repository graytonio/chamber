package configparser_test

import (
	"reflect"
	"testing"
	"testing/fstest"

	"github.com/graytonio/chamber/pkg/configparser"
)

var parse_want = &configparser.RepoInfo{
	Site:   "github.com",
	User:   "graytonio",
	Name:   "anteroom-templates",
	Ref:    "",
	Url:    "https://github.com/graytonio/anteroom-templates",
	Ssh:    "git@github.com:graytonio/anteroom-templates",
	Branch: "main",
}

type parse_test struct {
	Input string
	Want  *configparser.RepoInfo
}

var parse_tests = []parse_test{
	{"graytonio/anteroom-templates", parse_want},
	{"graytonio/anteroom-templates/solid-ts", parse_want},
	{"https://github.com/graytonio/anteroom-templates", parse_want},
	{"graytonio/anteroom-templates", parse_want},
	{"graytonio/anteroom-templates", parse_want},
}

func TestParseRepoUrl(t *testing.T) {
	for _, test := range parse_tests {
		if output, err := configparser.ParseRepoUrl(test.Input); !reflect.DeepEqual(output, test.Want) || err != nil {
			t.Errorf("FAILED\nWanted:\n%v\nGot:\n%v", test.Want, output)
		}
	}
}

var template_want = []configparser.Template{
	{Repo: parse_want, Template: "solid-ts"},
	{Repo: parse_want, Template: "next-ts"},
	{Repo: parse_want, Template: "vite-react"},
}

type template_test struct {
	RepoInfo *configparser.RepoInfo
	Want     []configparser.Template
}

var template_tests = []template_test{
	{RepoInfo: parse_want, Want: template_want},
}

func TestReadRepoTemplates(t *testing.T) {
	tmp_fs := fstest.MapFS{}
	for _, test := range template_tests {
		if output, err := configparser.ReadRepoTemplates(tmp_fs, test.RepoInfo); !reflect.DeepEqual(output, test.Want) || err != nil {
			t.Errorf("FAILED\nWanted:\n%v\nGot:\n%v", test.Want, output)
		}
	}
}

// TODO Write Unit Test for ReadConfigFile
