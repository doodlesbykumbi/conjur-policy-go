package conjurpolicy

import "gopkg.in/yaml.v3"

type ResourceRef struct {
	Id   string `yaml:"id"`
	Type Type
}

func UserRef(id string) ResourceRef {
	return ResourceRef{
		Id:   id,
		Type: TypeUser,
	}
}

func LayerRef(id string) ResourceRef {
	return ResourceRef{
		Id:   id,
		Type: TypeLayer,
	}
}

func HostRef(id string) ResourceRef {
	return ResourceRef{
		Id:   id,
		Type: TypeHost,
	}
}

func VariableRef(id string) ResourceRef {
	return ResourceRef{
		Id:   id,
		Type: TypeVariable,
	}
}

func (r *ResourceRef) UnmarshalYAML(value *yaml.Node) (err error) {
	var id string
	if err = value.Decode(&id); err != nil {
		return
	}

	r.Id = id
	r.Type, err = ParseType(value.Tag[1:])
	if err != nil {
		return
	}
	return
}

func (r ResourceRef) MarshalYAML() (interface{}, error) {
	return &yaml.Node{
		Kind:  yaml.ScalarNode,
		Value: r.Id,
		Tag:   r.Type.Tag(),
		Style: yaml.TaggedStyle,
	}, nil
}
