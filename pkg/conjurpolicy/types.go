package conjurpolicy

import (
	"gopkg.in/yaml.v3"
)

//go:generate go run github.com/abice/go-enum -t yaml.tmpl

// Type is an enum representing conjur policy types
// ENUM(policy, variable, user, group, layer, grant, host, delete, permit, deny)
type Type int

type Resource interface {
	unused() // to prevent Resource from being used as a type
}

type Group struct {
	Resource    `yaml:"-"`
	Id          string      `yaml:"id"`
	Annotations Annotations `yaml:"annotations,omitempty"`
	Owner       ResourceRef `yaml:"owner,omitempty"`
}

type Annotations map[string]string

type Variable struct {
	Resource    `yaml:"-"`
	Id          string      `yaml:"id"`
	Annotations Annotations `yaml:"annotations,omitempty"`
	Kind        string      `yaml:"kind,omitempty"`
}

type User struct {
	Resource    `yaml:"-"`
	Id          string      `yaml:"id"`
	Owner       ResourceRef `yaml:"owner,omitempty"`
	Annotations Annotations `yaml:"annotations,omitempty"`
}

type Policy struct {
	Resource    `yaml:"-"`
	Id          string           `yaml:"id"`
	Annotations Annotations      `yaml:"annotations,omitempty"`
	Owner       ResourceRef      `yaml:"owner,omitempty"`
	Body        PolicyStatements `yaml:"body,omitempty"`
}

type Layer struct {
	Resource `yaml:"-"`
	Id       string `yaml:"id,omitempty"`
}

type Grant struct {
	Resource `yaml:"-"`
	Role     ResourceRef `yaml:"role"`
	Member   ResourceRef `yaml:"member"`
}

type Host struct {
	Resource    `yaml:"-"`
	Id          string            `yaml:"id,omitempty"`
	Owner       ResourceRef       `yaml:"owner,omitempty"`
	Body        PolicyStatements  `yaml:"body,omitempty"`
	Annotations map[string]string `yaml:"annotations,omitempty"`
}

type Delete struct {
	Resource `yaml:"-"`
	Record   ResourceRef `yaml:"record"`
}

type Permit struct {
	Resource   `yaml:"-"`
	Role       ResourceRef `yaml:"role"`
	Privileges []Privilege `yaml:"privileges,flow"`
	Resources  ResourceRef `yaml:"resource"`
}

type Deny struct {
	Resource   `yaml:"-"`
	Role       ResourceRef `yaml:"role"`
	Privileges []Privilege `yaml:"privileges,flow"`
	Resources  ResourceRef `yaml:"resource"`
}

type PolicyStatements []Resource

func (p PolicyStatements) MarshalYAML() (interface{}, error) {
	if len(p) == 0 { // empty policy statements should result in empty yaml
		return &yaml.Node{Kind: yaml.ScalarNode}, nil
	}
	type aliasType PolicyStatements
	return aliasType(p), nil
}
