package conjurpolicy

import (
	"errors"
	"os"

	"gopkg.in/yaml.v3"
)

func include(node *yaml.Node) ([]Resource, error) {
	fileName, err := fileName(node)
	if err != nil {
		return nil, err
	}
	file, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	var res PolicyStatements
	err = yaml.Unmarshal(file, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func fileName(node *yaml.Node) (string, error) {
	switch {
	case len(node.Value) > 0:
		return node.Value, nil
	case len(node.Content) == 2:
		file := node.Content[1]
		return file.Value, nil
	default:
		return "", errors.New("failed to extract file name for include tag")
	}

}
