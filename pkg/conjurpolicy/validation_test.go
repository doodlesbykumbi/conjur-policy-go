package conjurpolicy

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestValidateDuplicatesInPolicy(t *testing.T) {
	tests := []struct {
		name    string
		policy  string
		wantErr string
	}{{
		"duplicate attribute member",
		`  - !host host1
  - !host host2

  - !variable secret
  - !group secret-consumers

  - !permit
    role: !group secret-consumers
    resource: !variable secret
    privilege: [ read, execute ]

  - !grant
    role: !group secret-consumers
    member: !host host1
    member: !host host2`,
		`yaml: unmarshal errors:
  line 15: mapping key "member" already defined at line 14`,
	}, {
		"duplicate attribute description",
		`  - !host
    id: test-app
    annotations:
      description: value1
      description: value2
`,
		`yaml: unmarshal errors:
  line 5: mapping key "description" already defined at line 4`,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var p PolicyStatements
			err := yaml.Unmarshal([]byte(tt.policy), &p)
			assert.EqualError(t, err, tt.wantErr)
		})
	}
}
