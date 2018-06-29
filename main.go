package main

import (
	"io"
	"log"
	"os"

	yaml "gopkg.in/yaml.v2"
)

func main() {
	e := yaml.NewEncoder(os.Stdout)

	defer e.Close()

	for _, arg := range os.Args[1:] {
		f, err := os.Open(arg)
		if err != nil {
			log.Fatalf("Open error: %s", err)
		}

		defer f.Close()

		d := yaml.NewDecoder(f)

		for {
			var data interface{}
			if dErr := d.Decode(&data); dErr != nil {
				if dErr == io.EOF {
					break
				}

				log.Fatalf("Decode error: %s", dErr)
			}
			e.Encode(data)
		}
	}
}
