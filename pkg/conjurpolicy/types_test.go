package conjurpolicy

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestResourceMarshalUnmarshal(t *testing.T) {
	testCases := []struct {
		name     string
		policy   Resource
		expected string
	}{
		{
			name: "empty-policy",
			policy: Policy{
				Id: "empty-policy",
			},
			expected: `!policy
id: empty-policy
`,
		},
		{
			name: "delete-policy",
			policy: Delete{
				Record: HostRef("test-host"),
			},
			expected: `!delete
record: !host test-host
`,
		},
		{
			name: "permit",
			policy: Permit{
				Role:       HostRef("/test-host"),
				Privileges: []Privilege{PrivilegeRead},
				Resources:  VariableRef("test-variable"),
			},
			expected: `!permit
role: !host /test-host
privileges: [read]
resource: !variable test-variable
`,
		},
		{
			name: "deny",
			policy: Deny{
				Role:       HostRef("/test-host"),
				Privileges: []Privilege{PrivilegeCreate, PrivilegeExecute},
				Resources:  VariableRef("test-variable"),
			},
			expected: `!deny
role: !host /test-host
privileges: [create, execute]
resource: !variable test-variable
`,
		},
		{
			name: "policy-with-annotations",
			policy: Policy{
				Id: "policy-with-annotations",
				Annotations: Annotations{
					"description": "this is a test policy",
				},
			},
			expected: `!policy
id: policy-with-annotations
annotations:
    description: this is a test policy
`,
		},
		{
			name: "policy-with-owner",
			policy: Policy{
				Id: "policy-with-owner",
				Owner: ResourceRef{
					Id:   "test-owner",
					Kind: KindUser,
				},
			},
			expected: `!policy
id: policy-with-owner
owner: !user test-owner
`,
		},
		{
			name: "policy-with-body",
			policy: Policy{
				Id: "policy-with-body",
				Body: PolicyStatements{
					User{
						Id: "test-user",
					},
					Group{
						Id: "test-group",
					},
					Variable{
						Id:   "test-variable",
						Kind: "text",
					},
				},
			},
			expected: `!policy
id: policy-with-body
body:
    - !user
      id: test-user
    - !group
      id: test-group
    - !variable
      id: test-variable
      kind: text
`,
		}, {
			name: "policy-with-layer",
			policy: Policy{
				Id: "policy-with-body",
				Body: PolicyStatements{
					Layer{},
					Grant{
						Role:   LayerRef(""),
						Member: LayerRef("test-layer"),
					},
				},
			},
			expected: `!policy
id: policy-with-body
body:
    - !layer
    - !grant
      role: !layer
      member: !layer test-layer
`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Marshal
			actual, err := yaml.Marshal(tc.policy)
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, string(actual))

			// Unmarshal
			policy := reflect.New(reflect.TypeOf(tc.policy)).Elem()
			// https://github.com/go-yaml/yaml/issues/769
			err = yaml.Unmarshal([]byte(tc.expected), policy.Addr().Interface())
			assert.NoError(t, err)
			assert.Equal(t, tc.policy, policy.Interface())
		})
	}
}
