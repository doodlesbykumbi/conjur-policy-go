# conjur-policy-go

The general goal is to be able to represent Conjur policy as objects in Go, and to be able to emit valid Conjur policy from objects in Go.

For this Go representation:

```go
policy := PolicyBody{
  Policy{
   Id:    "dev",
   Owner: UserRef("admin"),
   Body: PolicyBody{
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
   Body: PolicyBody{
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
```

Outputs this YAML (and vice-versa):

```yaml
- !policy
  id: dev
  owner: !user admin
  body:
    - !group
      id: bar
      owner: !user foo
    - !user
      id: foo
      owner: !user admin
- !policy
  id: pcf/prod
  owner: !user admin
  body:
    - !group
      id: bar
      owner: !user foo
    - !user
      id: foo
      owner: !user admin
```
