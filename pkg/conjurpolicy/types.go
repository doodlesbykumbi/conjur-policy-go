package conjurpolicy

import "gopkg.in/yaml.v3"

func UserRef(id string) ResourceRef {
	return ResourceRef{
		Id:   id,
		Kind: "user",
	}
}

type ResourceRef struct {
	Id   string
	Kind string
}

func (r *ResourceRef) UnmarshalYAML(value *yaml.Node) error {
	var id string
	if err := value.Decode(&id); err != nil {
		return err
	}

	r.Id = id
	r.Kind = value.Tag[1:]

	return nil
}

func (r ResourceRef) MarshalYAML() (interface{}, error) {
	return &yaml.Node{
		Kind:  yaml.ScalarNode,
		Value: r.Id,
		Tag:   "!" + r.Kind,
		Style: yaml.TaggedStyle,
	}, nil
}

type Group struct {
	Id          string      `yaml:"id"`
	Annotations Annotations `yaml:"annotations,omitempty"`
	Owner       ResourceRef `yaml:"owner,omitempty"`
}

func (g Group) MarshalYAML() (interface{}, error) {
	type aliasType Group
	data := (aliasType)(g)

	node := &yaml.Node{Kind: yaml.MappingNode}
	if err := node.Encode(&data); err != nil {
		return nil, err
	}
	node.Tag = "!group"
	node.Style = yaml.TaggedStyle

	return node, nil
}

type Annotations map[string]string

type Variable struct {
	Id          string      `yaml:"id"`
	Annotations Annotations `yaml:"annotations,omitempty"`
	Kind        string      `yaml:"kind,omitempty"`
}

func (v Variable) MarshalYAML() (interface{}, error) {
	type aliasType Variable
	data := (aliasType)(v)

	node := &yaml.Node{Kind: yaml.MappingNode}
	if err := node.Encode(&data); err != nil {
		return nil, err
	}
	node.Tag = "!variable"
	node.Style = yaml.TaggedStyle

	return node, nil
}

type User struct {
	Id          string      `yaml:"id"`
	Owner       ResourceRef `yaml:"owner,omitempty"`
	Annotations Annotations `yaml:"annotations,omitempty"`
}

func (u User) MarshalYAML() (interface{}, error) {
	type aliasType User
	data := (aliasType)(u)

	node := &yaml.Node{Kind: yaml.MappingNode}
	if err := node.Encode(&data); err != nil {
		return nil, err
	}
	node.Tag = "!user"
	node.Style = yaml.TaggedStyle

	return node, nil
}

type PolicyBody []interface{}

func (pb *PolicyBody) UnmarshalYAML(value *yaml.Node) error {
	things := []interface{}{}
	for _, innerNode := range value.Content {
		var thing interface{}
		switch innerNode.Tag {
		case "!policy":
			var policy Policy
			if err := innerNode.Decode(&policy); err != nil {
				return err
			}
			thing = policy
		case "!group":
			var group Group
			if err := innerNode.Decode(&group); err != nil {
				return err
			}
			thing = group
		case "!user":
			var user User
			if err := innerNode.Decode(&user); err != nil {
				return err
			}
			thing = user
		case "!variable":
			var variable Variable
			if err := innerNode.Decode(&variable); err != nil {
				return err
			}
			thing = variable
		}

		things = append(things, thing)
	}

	*pb = PolicyBody(things)

	return nil
}

type Policy struct {
	Id          string      `yaml:"id"`
	Annotations Annotations `yaml:"annotations,omitempty"`
	Owner       ResourceRef `yaml:"owner,omitempty"`
	Body        PolicyBody  `yaml:"body,omitempty"`
}

func (p Policy) MarshalYAML() (interface{}, error) {
	type aliasType Policy
	data := (aliasType)(p)

	node := &yaml.Node{Kind: yaml.MappingNode}
	if err := node.Encode(&data); err != nil {
		return nil, err
	}
	node.Tag = "!policy"
	node.Style = yaml.TaggedStyle

	return node, nil
}
