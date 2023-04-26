package conjurpolicy

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestResourceMarshalUnmarshal(t *testing.T) {
	testCases := []struct {
		name     string
		policy   PolicyStatements
		expected string
	}{
		{
			name:   "empty-policy",
			policy: PolicyStatements{Policy{Id: "empty-policy"}},
			expected: `- !policy
  id: empty-policy
`,
		},
		{
			name:   "delete-policy",
			policy: PolicyStatements{Delete{Record: HostRef("test-host")}},
			expected: `- !delete
  record: !host test-host
`,
		},
		{
			name: "permit",
			policy: PolicyStatements{Permit{
				Role:       HostRef("/test-host"),
				Privileges: []Privilege{PrivilegeRead},
				Resources:  VariableRef("test-variable"),
			}},
			expected: `- !permit
  role: !host /test-host
  privileges: [read]
  resource: !variable test-variable
`,
		},
		{
			name: "deny",
			policy: PolicyStatements{Deny{
				Role:       HostRef("/test-host"),
				Privileges: []Privilege{PrivilegeCreate, PrivilegeExecute},
				Resources:  VariableRef("test-variable"),
			}},
			expected: `- !deny
  role: !host /test-host
  privileges: [create, execute]
  resource: !variable test-variable
`,
		},
		{
			name:   "empty-host",
			policy: PolicyStatements{Host{}},
			expected: `- !host
`,
		},
		{
			name: "policy-with-annotations",
			policy: PolicyStatements{Policy{
				Id: "policy-with-annotations",
				Annotations: Annotations{
					"description": "this is a test policy",
				},
			}},
			expected: `- !policy
  id: policy-with-annotations
  annotations:
    description: this is a test policy
`,
		},
		{
			name: "policy-with-owner",
			policy: PolicyStatements{Policy{
				Id: "policy-with-owner",
				Owner: ResourceRef{
					Id:   "test-owner",
					Type: TypeUser,
				},
			}},
			expected: `- !policy
  id: policy-with-owner
  owner: !user test-owner
`,
		},
		{
			name: "policy-with-body",
			policy: PolicyStatements{Policy{
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
			}},
			expected: `- !policy
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
			policy: PolicyStatements{Policy{
				Id: "policy-with-body",
				Body: PolicyStatements{
					Layer{},
					Grant{
						Role:   LayerRef(""),
						Member: LayerRef("test-layer"),
					},
				},
			}},
			expected: `- !policy
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
			var policy PolicyStatements
			err = yaml.Unmarshal([]byte(tc.expected), &policy)
			assert.NoError(t, err)
			assert.Equal(t, tc.policy, policy)
		})
	}
}

func BenchmarkResourceMarshal(b *testing.B) {
	policy := PolicyStatements{
		Policy{
			Id:    "dev",
			Owner: UserRef("admin"),
			Annotations: Annotations{
				"foo": "bar",
			},
			Body: PolicyStatements{
				Group{
					Id:    "bar",
					Owner: UserRef("foo"),
				},
				User{
					Id:    "foo",
					Owner: UserRef("admin"),
				},
			},
		},
		Policy{
			Owner: UserRef("admin"),
			Id:    "pcf/prod",
			Body: PolicyStatements{
				Group{
					Id:    "bar",
					Owner: UserRef("foo"),
				},
				User{
					Id:    "foo",
					Owner: UserRef("admin"),
				},
			},
		},
	}
	for i := 0; i < b.N; i++ {
		_, _ = yaml.Marshal(policy)
	}
}

func BenchmarkResourceUnmarshal(b *testing.B) {
	var p PolicyStatements
	policy := []byte(`
- !policy
  id: dev
  owner: !user /admin
  annotations:
    foo: bar
  body:
    - !policy
      id: /inner
      body:
      - !group
        id: bar
        owner: !user foo
    - !group
      id: bar
      owner: !user foo
    - !user
      id: foo
      owner: !user admin
`)
	for i := 0; i < b.N; i++ {
		_ = yaml.Unmarshal(policy, &p)
	}
}
