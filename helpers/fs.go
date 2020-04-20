package helpers

import (
	"io/ioutil"
	"strings"
)

func GetFile(path string) []string {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	values := strings.Split(string(file), "\n")
	return values
}
