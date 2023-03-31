package main

import (
	"fmt"

	conjurpolicy "github.com/doodlesbykumbi/conjur-policy-go/pkg/conjurpolicy"

	"gopkg.in/yaml.v3"
)

func marshal() {
	policy := conjurpolicy.PolicyBody{
		conjurpolicy.Policy{
			Id:    "dev",
			Owner: conjurpolicy.UserRef("admin"),
			Annotations: conjurpolicy.Annotations{
				"foo": "bar",
			},
			Body: []interface{}{
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
			Body: []interface{}{
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
		panic(err)
	}
	fmt.Printf("unmarshaled:\n%+v\n", string(data))
}

func unmarshal() {
	var p conjurpolicy.PolicyBody
	err := yaml.Unmarshal([]byte(`
- !policy
  id: dev
  owner: !user admin
  annotations:
    foo: bar
  body:
    - !group
      id: bar
      owner: !user foo
    - !user
      id: foo
      owner: !user admin
`), &p)
	if err != nil {
		panic(err)
	}
	fmt.Printf("unmarshaled:\n%+v\n", p)
}

func main() {
	marshal()
	unmarshal()
}
