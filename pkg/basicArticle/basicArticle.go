package basicArticle

import (
	"errors"
	"strings"
)

func GetTitle(fileName string) (string, error) {
	dotCount := strings.Count(fileName, ".")
	if dotCount != 1 {
		return "", errors.New("filename must contain a single dot")
	}
	return strings.Split(fileName, ".")[0], nil
}

func GetContent(fileName string) (string, error) {
	return "", nil
}
