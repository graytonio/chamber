package configparser

import (
	"reflect"
	"testing"
)

var parse_want = &RepoInfo{
	Site: "github.com",
	User: "graytonio",
	Name: "anteroom-templates",
	Ref:  "",
	Url:  "https://github.com/graytonio/anteroom-templates",
	Ssh:  "git@github.com:graytonio/anteroom-templates",
}

type parse_test struct {
	input string
	want  *RepoInfo
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
		if output, err := parse_repo_url(test.input); !reflect.DeepEqual(output, test.want) || err != nil {
			t.Errorf("FAILED\n\tWanted:\n%v\nGot:\n%v", test.want, output)
		}
	}
}
