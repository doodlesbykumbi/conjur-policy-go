package conjurpolicy

import (
	"reflect"

	"gopkg.in/yaml.v3"
)

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
	if allFieldsEmpty(v) {
		node.Kind = yaml.ScalarNode
	} else {
		node.Kind = yaml.MappingNode
		if err := node.Encode(&data); err != nil {
			return nil, err
		}
	}
	node.Tag = kind.Tag()
	node.Style = yaml.TaggedStyle
	return node, nil
}

func allFieldsEmpty(r interface{}) bool {
	v := reflect.ValueOf(r)
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		if !reflect.DeepEqual(f.Interface(), reflect.Zero(f.Type()).Interface()) {
			return false
		}
	}
	return true
}
