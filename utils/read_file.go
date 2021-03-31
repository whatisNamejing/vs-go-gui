package utils

import (
	"io/ioutil"
	"strings"
)

func ReadFile(fileName string) map[int]string {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	var t = strings.Split(string(data), "\n")
	var stringMap map[int]string
	stringMap = make(map[int]string)
	for i := 0; i < len(t); i++ {
		if t[i] != "" {
			stringMap[i] = t[i]
		}
	}
	return stringMap
}
