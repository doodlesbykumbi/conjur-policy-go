// Code generated by go-enum DO NOT EDIT.
// Version:
// Revision:
// Build Date:
// Built By:

package conjurpolicy

import (
	"errors"
	"fmt"

	"gopkg.in/yaml.v3"
)

const (
	// TypePolicy is a Type of type Policy.
	TypePolicy Type = iota
	// TypeVariable is a Type of type Variable.
	TypeVariable
	// TypeUser is a Type of type User.
	TypeUser
	// TypeGroup is a Type of type Group.
	TypeGroup
	// TypeLayer is a Type of type Layer.
	TypeLayer
	// TypeGrant is a Type of type Grant.
	TypeGrant
	// TypeHost is a Type of type Host.
	TypeHost
	// TypeDelete is a Type of type Delete.
	TypeDelete
	// TypePermit is a Type of type Permit.
	TypePermit
	// TypeDeny is a Type of type Deny.
	TypeDeny
)

var ErrInvalidType = errors.New("not a valid Type")

const _TypeName = "policyvariableusergrouplayergranthostdeletepermitdeny"

var _TypeMap = map[Type]string{
	TypePolicy:   _TypeName[0:6],
	TypeVariable: _TypeName[6:14],
	TypeUser:     _TypeName[14:18],
	TypeGroup:    _TypeName[18:23],
	TypeLayer:    _TypeName[23:28],
	TypeGrant:    _TypeName[28:33],
	TypeHost:     _TypeName[33:37],
	TypeDelete:   _TypeName[37:43],
	TypePermit:   _TypeName[43:49],
	TypeDeny:     _TypeName[49:53],
}

// String implements the Stringer interface.
func (x Type) String() string {
	if str, ok := _TypeMap[x]; ok {
		return str
	}
	return fmt.Sprintf("Type(%d)", x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x Type) IsValid() bool {
	_, ok := _TypeMap[x]
	return ok
}

var _TypeValue = map[string]Type{
	_TypeName[0:6]:   TypePolicy,
	_TypeName[6:14]:  TypeVariable,
	_TypeName[14:18]: TypeUser,
	_TypeName[18:23]: TypeGroup,
	_TypeName[23:28]: TypeLayer,
	_TypeName[28:33]: TypeGrant,
	_TypeName[33:37]: TypeHost,
	_TypeName[37:43]: TypeDelete,
	_TypeName[43:49]: TypePermit,
	_TypeName[49:53]: TypeDeny,
}

// ParseType attempts to convert a string to a Type.
func ParseType(name string) (Type, error) {
	if x, ok := _TypeValue[name]; ok {
		return x, nil
	}
	return Type(0), fmt.Errorf("%s is %w", name, ErrInvalidType)
}

func (p Policy) MarshalYAML() (interface{}, error) {
	type aliasPolicy Policy
	data := aliasPolicy(p)
	node := &yaml.Node{}
	node.Kind = yaml.MappingNode
	if err := node.Encode(data); err != nil {
		return nil, err
	}
	// Avoid emitting strings like `- !variable {}` and instead emit `- !variable` by setting Kind to ScalarNode
	// when the resource struct is empty!
	if len(node.Content) == 0 {
		node.Kind = yaml.ScalarNode
	}
	node.Tag = TypePolicy.Tag()
	node.Style = yaml.TaggedStyle
	return node, nil
}

func (v Variable) MarshalYAML() (interface{}, error) {
	type aliasVariable Variable
	data := aliasVariable(v)
	node := &yaml.Node{}
	node.Kind = yaml.MappingNode
	if err := node.Encode(data); err != nil {
		return nil, err
	}
	// Avoid emitting strings like `- !variable {}` and instead emit `- !variable` by setting Kind to ScalarNode
	// when the resource struct is empty!
	if len(node.Content) == 0 {
		node.Kind = yaml.ScalarNode
	}
	node.Tag = TypeVariable.Tag()
	node.Style = yaml.TaggedStyle
	return node, nil
}

func (u User) MarshalYAML() (interface{}, error) {
	type aliasUser User
	data := aliasUser(u)
	node := &yaml.Node{}
	node.Kind = yaml.MappingNode
	if err := node.Encode(data); err != nil {
		return nil, err
	}
	// Avoid emitting strings like `- !variable {}` and instead emit `- !variable` by setting Kind to ScalarNode
	// when the resource struct is empty!
	if len(node.Content) == 0 {
		node.Kind = yaml.ScalarNode
	}
	node.Tag = TypeUser.Tag()
	node.Style = yaml.TaggedStyle
	return node, nil
}

func (g Group) MarshalYAML() (interface{}, error) {
	type aliasGroup Group
	data := aliasGroup(g)
	node := &yaml.Node{}
	node.Kind = yaml.MappingNode
	if err := node.Encode(data); err != nil {
		return nil, err
	}
	// Avoid emitting strings like `- !variable {}` and instead emit `- !variable` by setting Kind to ScalarNode
	// when the resource struct is empty!
	if len(node.Content) == 0 {
		node.Kind = yaml.ScalarNode
	}
	node.Tag = TypeGroup.Tag()
	node.Style = yaml.TaggedStyle
	return node, nil
}

func (l Layer) MarshalYAML() (interface{}, error) {
	type aliasLayer Layer
	data := aliasLayer(l)
	node := &yaml.Node{}
	node.Kind = yaml.MappingNode
	if err := node.Encode(data); err != nil {
		return nil, err
	}
	// Avoid emitting strings like `- !variable {}` and instead emit `- !variable` by setting Kind to ScalarNode
	// when the resource struct is empty!
	if len(node.Content) == 0 {
		node.Kind = yaml.ScalarNode
	}
	node.Tag = TypeLayer.Tag()
	node.Style = yaml.TaggedStyle
	return node, nil
}

func (g Grant) MarshalYAML() (interface{}, error) {
	type aliasGrant Grant
	data := aliasGrant(g)
	node := &yaml.Node{}
	node.Kind = yaml.MappingNode
	if err := node.Encode(data); err != nil {
		return nil, err
	}
	// Avoid emitting strings like `- !variable {}` and instead emit `- !variable` by setting Kind to ScalarNode
	// when the resource struct is empty!
	if len(node.Content) == 0 {
		node.Kind = yaml.ScalarNode
	}
	node.Tag = TypeGrant.Tag()
	node.Style = yaml.TaggedStyle
	return node, nil
}

func (h Host) MarshalYAML() (interface{}, error) {
	type aliasHost Host
	data := aliasHost(h)
	node := &yaml.Node{}
	node.Kind = yaml.MappingNode
	if err := node.Encode(data); err != nil {
		return nil, err
	}
	// Avoid emitting strings like `- !variable {}` and instead emit `- !variable` by setting Kind to ScalarNode
	// when the resource struct is empty!
	if len(node.Content) == 0 {
		node.Kind = yaml.ScalarNode
	}
	node.Tag = TypeHost.Tag()
	node.Style = yaml.TaggedStyle
	return node, nil
}

func (d Delete) MarshalYAML() (interface{}, error) {
	type aliasDelete Delete
	data := aliasDelete(d)
	node := &yaml.Node{}
	node.Kind = yaml.MappingNode
	if err := node.Encode(data); err != nil {
		return nil, err
	}
	// Avoid emitting strings like `- !variable {}` and instead emit `- !variable` by setting Kind to ScalarNode
	// when the resource struct is empty!
	if len(node.Content) == 0 {
		node.Kind = yaml.ScalarNode
	}
	node.Tag = TypeDelete.Tag()
	node.Style = yaml.TaggedStyle
	return node, nil
}

func (p Permit) MarshalYAML() (interface{}, error) {
	type aliasPermit Permit
	data := aliasPermit(p)
	node := &yaml.Node{}
	node.Kind = yaml.MappingNode
	if err := node.Encode(data); err != nil {
		return nil, err
	}
	// Avoid emitting strings like `- !variable {}` and instead emit `- !variable` by setting Kind to ScalarNode
	// when the resource struct is empty!
	if len(node.Content) == 0 {
		node.Kind = yaml.ScalarNode
	}
	node.Tag = TypePermit.Tag()
	node.Style = yaml.TaggedStyle
	return node, nil
}

func (d Deny) MarshalYAML() (interface{}, error) {
	type aliasDeny Deny
	data := aliasDeny(d)
	node := &yaml.Node{}
	node.Kind = yaml.MappingNode
	if err := node.Encode(data); err != nil {
		return nil, err
	}
	// Avoid emitting strings like `- !variable {}` and instead emit `- !variable` by setting Kind to ScalarNode
	// when the resource struct is empty!
	if len(node.Content) == 0 {
		node.Kind = yaml.ScalarNode
	}
	node.Tag = TypeDeny.Tag()
	node.Style = yaml.TaggedStyle
	return node, nil
}

func toID(node *yaml.Node) *yaml.Node {
	node.Kind = yaml.MappingNode
	node.Content = []*yaml.Node{
		&yaml.Node{
			Value: "id",
			Kind:  yaml.ScalarNode,
		},
		&yaml.Node{
			Value: node.Value,
			Kind:  yaml.ScalarNode,
		},
	}
	return node
}

func (s *PolicyStatements) UnmarshalYAML(value *yaml.Node) error {
	var statements []Resource
	for _, node := range value.Content {
		var statement Resource
		switch node.Tag {

		case TypePolicy.Tag():
			var p Policy
			if len(node.Value) > 0 || len(node.Content) > 0 {
				if len(node.Content) == 0 && len(node.Value) > 0 {
					node = toID(node)
				}
				node.KnownFields(true)
				if err := node.Decode(&p); err != nil {
					return err
				}
			}
			statement = p

		case TypeVariable.Tag():
			var v Variable
			if len(node.Value) > 0 || len(node.Content) > 0 {
				if len(node.Content) == 0 && len(node.Value) > 0 {
					node = toID(node)
				}
				node.KnownFields(true)
				if err := node.Decode(&v); err != nil {
					return err
				}
			}
			statement = v

		case TypeUser.Tag():
			var u User
			if len(node.Value) > 0 || len(node.Content) > 0 {
				if len(node.Content) == 0 && len(node.Value) > 0 {
					node = toID(node)
				}
				node.KnownFields(true)
				if err := node.Decode(&u); err != nil {
					return err
				}
			}
			statement = u

		case TypeGroup.Tag():
			var g Group
			if len(node.Value) > 0 || len(node.Content) > 0 {
				if len(node.Content) == 0 && len(node.Value) > 0 {
					node = toID(node)
				}
				node.KnownFields(true)
				if err := node.Decode(&g); err != nil {
					return err
				}
			}
			statement = g

		case TypeLayer.Tag():
			var l Layer
			if len(node.Value) > 0 || len(node.Content) > 0 {
				if len(node.Content) == 0 && len(node.Value) > 0 {
					node = toID(node)
				}
				node.KnownFields(true)
				if err := node.Decode(&l); err != nil {
					return err
				}
			}
			statement = l

		case TypeGrant.Tag():
			var g Grant
			if len(node.Value) > 0 || len(node.Content) > 0 {
				if len(node.Content) == 0 && len(node.Value) > 0 {
					node = toID(node)
				}
				node.KnownFields(true)
				if err := node.Decode(&g); err != nil {
					return err
				}
			}
			statement = g

		case TypeHost.Tag():
			var h Host
			if len(node.Value) > 0 || len(node.Content) > 0 {
				if len(node.Content) == 0 && len(node.Value) > 0 {
					node = toID(node)
				}
				node.KnownFields(true)
				if err := node.Decode(&h); err != nil {
					return err
				}
			}
			statement = h

		case TypeDelete.Tag():
			var d Delete
			if len(node.Value) > 0 || len(node.Content) > 0 {
				if len(node.Content) == 0 && len(node.Value) > 0 {
					node = toID(node)
				}
				node.KnownFields(true)
				if err := node.Decode(&d); err != nil {
					return err
				}
			}
			statement = d

		case TypePermit.Tag():
			var p Permit
			if len(node.Value) > 0 || len(node.Content) > 0 {
				if len(node.Content) == 0 && len(node.Value) > 0 {
					node = toID(node)
				}
				node.KnownFields(true)
				if err := node.Decode(&p); err != nil {
					return err
				}
			}
			statement = p

		case TypeDeny.Tag():
			var d Deny
			if len(node.Value) > 0 || len(node.Content) > 0 {
				if len(node.Content) == 0 && len(node.Value) > 0 {
					node = toID(node)
				}
				node.KnownFields(true)
				if err := node.Decode(&d); err != nil {
					return err
				}
			}
			statement = d
		}
		statements = append(statements, statement)
	}
	*s = statements
	return nil
}

// Tag is a method that returns a YAML tag from entity kind
func (t Type) Tag() string {
	return "!" + t.String()
}
