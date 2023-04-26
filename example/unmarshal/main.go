package main

import (
	"fmt"
	"github.com/doodlesbykumbi/conjur-policy-go/pkg/conjurpolicy"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

func main() {
	data, err := os.ReadFile("./policy.yaml")
	if err != nil {
		log.Fatal(err)
	}

	var p conjurpolicy.PolicyStatements
	err = yaml.Unmarshal(data, &p)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("unmarshaled:\n%+v\n", p)
}
