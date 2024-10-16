package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/neovim/go-client/nvim/plugin"
	"gopkg.in/yaml.v2"
)

func Split(args []string) {
	content, err := os.ReadFile(args[0])
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

func main() {
	plugin.Main(func(p *plugin.Plugin) error {
		p.HandleFunction(&plugin.FunctionOptions{Name: "KubeSplit"}, Split)
		return nil
	})
}
