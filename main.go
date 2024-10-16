package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
	"strings"
)

func main() {
	args := os.Args
	filePath := args[1]
	content, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	ymls := strings.Split(string(content), "---")
	for _, v := range ymls {
		obj := make(map[string]any)
		err := yaml.Unmarshal([]byte(v), &obj)
		if err != nil {
			panic(err)
		}
		kindAny := obj["kind"]
		kind := kindAny.(string)
		file, err := os.Create(fmt.Sprintf("%s.yaml", kind))
		if err != nil {
			panic(err)
		}
		_, err = file.WriteString(v)
		if err != nil {
			panic(err)
		}
	}
}
