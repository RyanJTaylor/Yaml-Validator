package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

func main() {
	dir := flag.String("dir", ".", "Directory containing YAML files")
	flag.Parse()

	fmt.Println("Checking directory:", *dir)

	files, err := os.ReadDir(*dir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		ext := strings.ToLower(filepath.Ext(file.Name()))
		if ext == ".yaml" || ext == ".yml" {
			path := filepath.Join(*dir, file.Name())

			data, err := os.ReadFile(path)
			if err != nil {
				fmt.Printf("%s - error reading file: %v\n", file.Name(), err)
				continue
			}

			var out interface{}
			if err := yaml.Unmarshal(data, &out); err != nil {
				fmt.Printf("%s - invalid YAML: %v\n", file.Name(), err)
			} else {
				fmt.Printf("%s - valid YAML\n", file.Name())
			}

		}
	}
}
