package configparser

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"regexp"

	"github.com/graytonio/chamber/pkg/chambererrors"
	"github.com/graytonio/chamber/pkg/utils"
)

type RepoInfo struct {
	Site   string
	User   string
	Name   string
	Ref    string
	Url    string
	Ssh    string
	Branch string
}

type Template struct {
	template string
	repo     *RepoInfo
}

type Config struct {
	Templates []Template
}

type config_struct struct {
	Repos []string `json:"repos"`
}

type manifest_struct struct {
	Templates []string `json:"templates"`
}

func ReadConfigFile(f fs.FS, path string) (*Config, *chambererrors.ChamberError) {
	content, err := fs.ReadFile(f, path)
	if err != nil {
		fmt.Println(err)
		return nil, &chambererrors.BadConfigError
	}

	var config_file config_struct
	json.Unmarshal(content, &config_file)

	var config Config
	for _, repo := range config_file.Repos {
		info, err := parse_repo_url(repo)
		if err != nil {
			return nil, err
		}

		templates, err := read_repo_templates(f, info)
		if err != nil {
			return nil, err
		}

		config.Templates = append(config.Templates, templates...)
	}

	return &config, nil
}

func read_repo_templates(f fs.FS, repo *RepoInfo) ([]Template, *chambererrors.ChamberError) {
	var manifest_url string = fmt.Sprintf("https://raw.githubusercontent.com/%s/%s/%s/manifest.json", repo.User, repo.Name, repo.Branch)
	content, err := utils.DownloadFile(manifest_url)
	if err != nil {
		fmt.Println(err.Error())
		return nil, &chambererrors.DownloadError
	}

	var manifest_file manifest_struct
	json.Unmarshal(content, &manifest_file)

	var templates []Template
	for _, template := range manifest_file.Templates {
		templates = append(templates, Template{template: template, repo: repo})
	}

	return templates, nil
}

var supported = []string{"github.com"}

func parse_repo_url(repo string) (*RepoInfo, *chambererrors.ChamberError) {
	regex := regexp.MustCompile(`^(?:(?:https:\/\/)?(?P<Domain1>[^:/]+\.[^:/]+)\/|git@(?P<Domain2>[^:/]+)[:/]|(?P<Domain3>[^/]+):)?(?P<User>[^/\s]+)\/(?P<Name>[^/\s#]+)(?:(?P<Subdir>(?:\/[^/\s#]+)+))?(?:\/)?(?:#(?P<Ref>.+))?`)
	remove_ext := regexp.MustCompile(`\.git$`)

	matches := regex.FindStringSubmatch(repo)
	var parseMap = make(map[string]string)
	for i, param := range regex.SubexpNames() {
		if i > 0 && i <= len(matches) {
			parseMap[param] = matches[i]
		}
	}

	var site string
	if parseMap["Domain1"] != "" {
		site = parseMap["Domain1"]
	} else if parseMap["Domain2"] != "" {
		site = parseMap["Domain2"]
	} else if parseMap["Domain3"] != "" {
		site = parseMap["Domain3"]
	} else {
		site = "github.com"
	}

	if !utils.SliceContains(supported, site) {
		return nil, &chambererrors.UnsupportedHostError
	}

	user := parseMap["User"]
	name := remove_ext.ReplaceAllString(parseMap["Name"], "")

	var ref string
	if ref = parseMap["Ref"]; parseMap["Ref"] != "" {
		ref = "HEAD"
	}

	url := fmt.Sprintf("https://%s/%s/%s", site, user, name)
	ssh := fmt.Sprintf("git@%s:%s/%s", site, user, name)

	return &RepoInfo{
		Site:   site,
		User:   user,
		Name:   name,
		Ref:    ref,
		Url:    url,
		Ssh:    ssh,
		Branch: "main",
	}, nil
}
