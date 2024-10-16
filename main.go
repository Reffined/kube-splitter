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
		fileName := fmt.Sprintf("%s.yaml", kind)
		//_, err = os.Stat(fileName)
		//if err != nil {
		//	if os.IsNotExist(err) {
		//		file, err := os.Create(fileName)
		//		if err != nil {
		//			panic(err)
		//		}
		//		_, err = file.WriteString(v)
		//		if err != nil {
		//			panic(err)
		//		}
		//		file.Close()
		//		continue
		//	}
		//}
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
