package ggi

import (
	"os"
	"testing"
)

func TestFetchAndSaveSuccess(t *testing.T) {
	out := "_test/Go.gitignore"

	os.Mkdir("_test", 0755)
	os.Remove(out)

	err := FetchAndSave("Go", out)

	if err != nil {
		t.Errorf("failed to fetch: %s", err)
	}

	_, err = os.Stat(out)

	if err != nil {
		t.Errorf("failed to save: %s", err)
	}
}

func TestFetchNotFound(t *testing.T) {
	err := FetchAndSave("ThereAreNoLanguageLikeThis", "_test/Foo.gitignore")

	if err == nil {
		t.Errorf("fetch must fail.")
	}
}
