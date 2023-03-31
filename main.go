package main

import (
	"fmt"
	"reflect"

	"gopkg.in/yaml.v3"
)

// copyWithoutMethods avoids infinite recursion when marshaling
func copyWithoutMethods(in interface{}) interface{} {
	t := reflect.TypeOf(in)
	if t.Kind() != reflect.Struct {
		return nil
	}

	// Create a new type that embeds the original struct type
	// but with no methods.
	fields := make([]reflect.StructField, 0, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Type.Kind() == reflect.Func {
			continue // skip methods
		}
		fields = append(fields, field)
	}
	newType := reflect.StructOf(fields)

	// Create a new value of the new type and set its fields to the
	// values of the original value.
	inValue := reflect.ValueOf(in)
	newValue := reflect.New(newType).Elem()
	for i := 0; i < newType.NumField(); i++ {
		newValue.Field(i).Set(inValue.FieldByName(newType.Field(i).Name))
	}

	return newValue.Interface()
}

type UserRef string

func (u UserRef) MarshalYAML() (interface{}, error) {
	return &yaml.Node{
		Kind:  yaml.ScalarNode,
		Value: string(u),
		Tag:   "!user",
		Style: yaml.TaggedStyle,
	}, nil
}

type Group struct {
	Id    string  `yaml:"id"`
	Owner UserRef `yaml:"owner"`
}

func (g Group) MarshalYAML() (interface{}, error) {
	data := copyWithoutMethods(g)

	node := &yaml.Node{Kind: yaml.MappingNode}
	if err := node.Encode(&data); err != nil {
		return nil, err
	}
	node.Tag = "!group"
	node.Style = yaml.TaggedStyle

	return node, nil
}

type User struct {
	Id    string  `yaml:"id"`
	Owner UserRef `yaml:"owner"`
}

func (u User) MarshalYAML() (interface{}, error) {
	data := copyWithoutMethods(u)

	node := &yaml.Node{Kind: yaml.MappingNode}
	if err := node.Encode(&data); err != nil {
		return nil, err
	}
	node.Tag = "!user"
	node.Style = yaml.TaggedStyle

	return node, nil
}

type PolicyDocument []interface{}

type Policy struct {
	Id    string        `yaml:"id"`
	Owner UserRef       `yaml:"owner"`
	Body  []interface{} `yaml:"body"`
}

func (p Policy) MarshalYAML() (interface{}, error) {
	data := copyWithoutMethods(p)

	node := &yaml.Node{Kind: yaml.MappingNode}
	if err := node.Encode(&data); err != nil {
		return nil, err
	}
	node.Tag = "!policy"
	node.Style = yaml.TaggedStyle

	return node, nil
}

func main() {
	policy := PolicyDocument{
		Policy{
			Id:    "dev",
			Owner: UserRef("admin"),
			Body: []interface{}{
				Group{
					Id:    "bar",
					Owner: UserRef("foo"),
				},
				User{
					Id:    "foo",
					Owner: UserRef("admin"),
				},
			},
		},
		Policy{
			Owner: UserRef("admin"),
			Id:    "pcf/prod",
			Body: []interface{}{
				Group{
					Id:    "bar",
					Owner: UserRef("foo"),
				},
				User{
					Id:    "foo",
					Owner: UserRef("admin"),
				},
			},
		},
	}

	data, err := yaml.Marshal(policy)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
