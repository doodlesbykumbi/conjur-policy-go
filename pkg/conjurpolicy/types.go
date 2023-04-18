package conjurpolicy

import (
	"gopkg.in/yaml.v3"
)

type Resource interface {
	unused() // to prevent Resource from being used as a type
}

type Resources interface {
	Policy | Variable | User | Group | Layer | Grant | Host | Delete | Permit | Deny
}

//go:generate go run .\cmd\typegenerator.go -type Resources2
type Resources2 struct {
	Policy
	Variable
	User
	Group
	Layer
	Grant
	Host
	Delete
	Permit
	Deny
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

func (s *PolicyStatements) UnmarshalYAML(value *yaml.Node) error {
	var statements []Resource
	for _, node := range value.Content {
		var statement Resource

		switch node.Tag {
		case KindPolicy.Tag():
			var policy Policy
			if err := node.Decode(&policy); err != nil {
				return err
			}
			statement = policy
		case KindGroup.Tag():
			var group Group
			if err := node.Decode(&group); err != nil {
				return err
			}
			statement = group
		case KindUser.Tag():
			var user User
			if err := node.Decode(&user); err != nil {
				return err
			}
			statement = user
		case KindVariable.Tag():
			var variable Variable
			if err := node.Decode(&variable); err != nil {
				return err
			}
			statement = variable
		case KindLayer.Tag():
			var layer Layer
			statement = layer
		case KindGrant.Tag():
			var grant Grant
			if err := node.Decode(&grant); err != nil {
				return err
			}
			statement = grant
		case KindHost.Tag():
			var host Host
			if len(node.Value) > 0 || len(node.Content) > 0 {
				if err := node.Decode(&host); err != nil {
					return err
				}
			}
			statement = host
		case KindDelete.Tag():
			var delete Delete
			if err := node.Decode(&delete); err != nil {
				return err
			}
			statement = delete
		case KindPermit.Tag():
			var permit Permit
			if err := node.Decode(&permit); err != nil {
				return err
			}
			statement = permit
		case KindDeny.Tag():
			var deny Deny
			if err := node.Decode(&deny); err != nil {
				return err
			}
			statement = deny
		}
		statements = append(statements, statement)
	}

	*s = statements

	return nil
}

func (p Policy) MarshalYAML() (interface{}, error) {
	return MarshalYAMLWithTag(p, KindPolicy)
}

func (v Variable) MarshalYAML() (interface{}, error) {
	return MarshalYAMLWithTag(v, KindVariable)
}

func (u User) MarshalYAML() (interface{}, error) {
	return MarshalYAMLWithTag(u, KindUser)
}

func (g Group) MarshalYAML() (interface{}, error) {
	return MarshalYAMLWithTag(g, KindGroup)
}

func (l Layer) MarshalYAML() (interface{}, error) {
	return MarshalYAMLWithTag(l, KindLayer)
}

func (g Grant) MarshalYAML() (interface{}, error) {
	return MarshalYAMLWithTag(g, KindGrant)
}

func (h Host) MarshalYAML() (interface{}, error) {
	return MarshalYAMLWithTag(h, KindHost)
}

func (d Delete) MarshalYAML() (interface{}, error) {
	return MarshalYAMLWithTag(d, KindDelete)
}

func (p Permit) MarshalYAML() (interface{}, error) {
	return MarshalYAMLWithTag(p, KindPermit)
}

func (d Deny) MarshalYAML() (interface{}, error) {
	return MarshalYAMLWithTag(d, KindDeny)
}
