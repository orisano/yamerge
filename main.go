package main

import (
	"fmt"
	"log"
	"os"

	"github.com/goccy/go-yaml"
	"github.com/imdario/mergo"
)

func main() {
	log.SetPrefix("yamerge: ")
	log.SetFlags(0)
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	v := make(map[string]interface{})
	for _, p := range os.Args[1:] {
		x, err := loadYAML(p)
		if err != nil {
			return fmt.Errorf("load yaml: %w", err)
		}
		if err := mergo.Map(&v, x, mergo.WithAppendSlice); err != nil {
			return fmt.Errorf("merge yaml: %w", err)
		}
	}
	if err := yaml.NewEncoder(os.Stdout).Encode(v); err != nil {
		return fmt.Errorf("encode yaml: %w", err)
	}
	return nil
}

func loadYAML(filePath string) (map[string]interface{}, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("open yaml: %w", err)
	}
	defer f.Close()

	x := make(map[string]interface{})
	if err := yaml.NewDecoder(f).Decode(&x); err != nil {
		return nil, fmt.Errorf("decode yaml: %w", err)
	}
	return x, nil
}
