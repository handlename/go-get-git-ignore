package ggi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const urlList = "https://api.github.com/repos/github/gitignore/contents"
const urlBase = "https://raw.githubusercontent.com/github/gitignore/master/%s.gitignore"

type file struct {
	Name string `json:"name"`
}

func fetch(lang string) ([]byte, error) {
	resp, err := http.Get(fmt.Sprintf(urlBase, lang))

	if err != nil {
		return nil, fmt.Errorf("failed to request: %s", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("failed to request: status code = %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}

func save(out string, body []byte) error {
	err := ioutil.WriteFile(out, body, 0644)

	if err != nil {
		return fmt.Errorf("failed to write to %s: %s", out, err)
	}

	return nil
}

// ListLangs lists .gitignore files.
func ListLangs() ([]string, error) {
	resp, err := http.Get(urlList)

	if err != nil {
		return nil, fmt.Errorf("failed to request: %s", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("failed to request: status code = %d", resp.StatusCode)
	}

	var files []file

	err = json.NewDecoder(resp.Body).Decode(&files)

	var names []string
	for _, f := range files {
		if strings.HasSuffix(f.Name, ".gitignore") {
			names = append(names, f.Name[:len(f.Name)-10])
		}
	}

	return names, err
}

// FetchAndSave saves .gitignore file for targeted language to given path.
func FetchAndSave(lang, out string) error {
	body, err := fetch(lang)

	if err != nil {
		return err
	}

	save(out, body)

	return nil
}
