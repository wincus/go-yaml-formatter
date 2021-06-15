package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

func main() {

	in, err := ioutil.ReadAll(os.Stdin)

	if err != nil {
		log.Panic(err)
	}

	out := make(map[string]interface{})

	err = yaml.Unmarshal(in, out)

	if err != nil {
		log.Panic(err)
	}

	y, err := yaml.Marshal(out)

	if err != nil {
		log.Panic(err)
	}

	fmt.Fprintf(os.Stdout, "%s", y)
}
