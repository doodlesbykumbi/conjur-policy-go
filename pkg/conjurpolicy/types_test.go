package conjurpolicy

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestPolicyMarshalUnmarshal(t *testing.T) {
	testCases := []struct {
		name     string
		policy   Policy
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
					Kind: "user",
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
				Body: PolicyBody{
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
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Marshal
			actual, err := yaml.Marshal(tc.policy)
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, string(actual))

			// Unmarsha;
			var policy Policy
			err = yaml.Unmarshal([]byte(tc.expected), &policy)
			assert.NoError(t, err)
			assert.Equal(t, tc.policy, policy)
		})
	}
}
