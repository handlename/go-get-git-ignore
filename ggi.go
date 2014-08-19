package ggi

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const urlBase = "https://raw.githubusercontent.com/github/gitignore/master/%s.gitignore"

func fetch(lang string) ([]byte, error) {
	resp, err := http.Get(fmt.Sprintf(urlBase, lang))

	if err != nil {
		return nil, fmt.Errorf("failed to request: %s", err)
	}

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

// FetchAndSave saves .gitignore file for targeted language to given path.
func FetchAndSave(lang, out string) error {
	body, err := fetch(lang)

	if err != nil {
		return err
	}

	save(out, body)

	return nil
}
