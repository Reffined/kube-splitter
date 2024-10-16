package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

func main() {
	filePath := flag.String("p", "", "kube-splitter -p <path to yaml file>")
	flag.Parse()
	content, err := os.ReadFile(*filePath)
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
		fileName := fmt.Sprintf("%s.yaml", kind)

		file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
		if err != nil {
			panic(err)
		}

		_, err = file.WriteString("---\n")
		if err != nil {
			file.Close()
			panic(err)
		}
		_, err = file.WriteString(v)
		if err != nil {
			file.Close()
			panic(err)
		}
		file.Close()
	}
}
