package conjurpolicy

import (
	"reflect"

	"gopkg.in/yaml.v3"
)

type Resource interface {
	unused() // to prevent Resource from being used as a type
}

type Kind string

const (
	PolicyKind   Kind = "policy"
	VariableKind Kind = "variable"
	UserKind     Kind = "user"
	GroupKind    Kind = "group"
	LayerKind    Kind = "layer"
	GrantKind    Kind = "grant"
)

func (t Kind) String() string {
	return string(t)
}

func (t Kind) Tag() string {
	return "!" + t.String()
}

func UserRef(id string) ResourceRef {
	return ResourceRef{
		Id:   id,
		Kind: UserKind,
	}
}

func LayerRef(id string) ResourceRef {
	return ResourceRef{
		Id:   id,
		Kind: LayerKind,
	}
}

// copyStructWithoutMethods avoids infinite recursion when marshaling
func copyStructWithoutMethods(in interface{}) interface{} {
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

func MarshalYAMLWithTag[T Resources](v T, kind Kind) (interface{}, error) {
	data := copyStructWithoutMethods(v)

	node := &yaml.Node{}
	switch kind {
	case LayerKind:
		node.Kind = yaml.ScalarNode
	default:
		node.Kind = yaml.MappingNode
		if err := node.Encode(&data); err != nil {
			return nil, err
		}
	}
	node.Tag = kind.Tag()
	node.Style = yaml.TaggedStyle
	return node, nil
}

type Resources interface {
	Policy | Variable | User | Group | Layer | Grant
}

func (p Policy) MarshalYAML() (interface{}, error) {
	return MarshalYAMLWithTag(p, PolicyKind)
}

func (v Variable) MarshalYAML() (interface{}, error) {
	return MarshalYAMLWithTag(v, VariableKind)
}

func (u User) MarshalYAML() (interface{}, error) {
	return MarshalYAMLWithTag(u, UserKind)
}

func (g Group) MarshalYAML() (interface{}, error) {
	return MarshalYAMLWithTag(g, GroupKind)
}

func (l Layer) MarshalYAML() (interface{}, error) {
	return MarshalYAMLWithTag(l, LayerKind)
}

func (g Grant) MarshalYAML() (interface{}, error) {
	return MarshalYAMLWithTag(g, GrantKind)
}

type ResourceRef struct {
	Id   string `yaml:"id"`
	Kind Kind
}

func (r *ResourceRef) UnmarshalYAML(value *yaml.Node) error {
	var id string
	if err := value.Decode(&id); err != nil {
		return err
	}

	r.Id = id
	r.Kind = Kind(value.Tag[1:])

	return nil
}

func (r ResourceRef) MarshalYAML() (interface{}, error) {
	return &yaml.Node{
		Kind:  yaml.ScalarNode,
		Value: r.Id,
		Tag:   r.Kind.Tag(),
		Style: yaml.TaggedStyle,
	}, nil
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

type PolicyStatements []Resource

func (s *PolicyStatements) UnmarshalYAML(value *yaml.Node) error {
	statements := []Resource{}
	for _, node := range value.Content {
		var statement Resource

		switch node.Tag {
		case PolicyKind.Tag():
			var policy Policy
			if err := node.Decode(&policy); err != nil {
				return err
			}
			statement = policy
		case GroupKind.Tag():
			var group Group
			if err := node.Decode(&group); err != nil {
				return err
			}

			statement = group
		case UserKind.Tag():
			var user User
			if err := node.Decode(&user); err != nil {
				return err
			}
			statement = user
		case VariableKind.Tag():
			var variable Variable
			if err := node.Decode(&variable); err != nil {
				return err
			}
			statement = variable
		case LayerKind.Tag():
			var layer Layer
			statement = layer
		case GrantKind.Tag():
			var grant Grant
			if err := node.Decode(&grant); err != nil {
				return err
			}
			statement = grant
		}
		statements = append(statements, statement)
	}

	*s = statements

	return nil
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
