package main

import (
	"github.com/doodlesbykumbi/conjur-policy-go/pkg/conjurpolicy"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

func main() {
	policy := conjurpolicy.PolicyStatements{
		conjurpolicy.Policy{
			Id:    "dev",
			Owner: conjurpolicy.UserRef("admin"),
			Annotations: conjurpolicy.Annotations{
				"foo": "bar",
			},
			Body: conjurpolicy.PolicyStatements{
				conjurpolicy.Group{
					Id:    "bar",
					Owner: conjurpolicy.UserRef("foo"),
				},
				conjurpolicy.User{
					Id:    "foo",
					Owner: conjurpolicy.UserRef("admin"),
				},
			},
		},
		conjurpolicy.Policy{
			Owner: conjurpolicy.UserRef("admin"),
			Id:    "pcf/prod",
			Body: conjurpolicy.PolicyStatements{
				conjurpolicy.Group{
					Id:    "bar",
					Owner: conjurpolicy.UserRef("foo"),
				},
				conjurpolicy.User{
					Id:    "foo",
					Owner: conjurpolicy.UserRef("admin"),
				},
			},
		},
	}

	data, err := yaml.Marshal(policy)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("marshaled:\n%+v\n", string(data))

	err = os.WriteFile("policy.yaml", data, 0666)
	if err != nil {
		log.Fatal(err)
	}
}
