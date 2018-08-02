package main

import (
	"log"
	"os"

	"github.com/imdario/mergo"
	yaml "gopkg.in/yaml.v2"
)

func loadYAML(filePath string) (map[string]interface{}, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	x := make(map[string]interface{})
	if err := yaml.NewDecoder(f).Decode(&x); err != nil {
		return nil, err
	}
	return x, nil
}

func main() {
	v := make(map[string]interface{})
	for _, p := range os.Args[1:] {
		x, err := loadYAML(p)
		if err != nil {
			log.Fatal(err)
		}
		if err := mergo.Map(&v, x, mergo.WithAppendSlice); err != nil {
			log.Fatal(err)
		}
	}
	if err := yaml.NewEncoder(os.Stdout).Encode(v); err != nil {
		log.Fatal(err)
	}
}
