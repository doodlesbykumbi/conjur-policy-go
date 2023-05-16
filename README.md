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

### Roadmap

What is needed to achieve feature parity with ruby implementation?

- [x] marshal models to YAML compliant with conjur syntax
- [x] unmarshal conjur syntax YAML to models
- [x] handle edge-cases (like [empty policy](https://github.com/cyberark/conjur-policy-parser/blob/master/spec/round-trip/yaml/empty.expected.yml))
- [x] special handling of inlined id on objects with single attribute
- [x] validate attribute duplications
- [x] provide negative test-cases (e.g. invalid syntax)
- [x] add license
- [ ] provide docs
- [ ] support for all entities [policy statements](https://docs.conjur.org/Latest/en/Content/Operations/Policy/policy-statement-ref.htm?tocpath=Fundamentals%7CPolicy%7CPolicy%20statement%20reference%7C_____0)
- [ ] verify fields available on model entities
- [ ] tag reference [strong typing](https://docs.conjur.org/Latest/en/Content/Operations/Policy/statement-ref-permit.htm?tocpath=Fundamentals%7CPolicy%7CPolicy%20statement%20reference%7C_____8#Attributes)
- [ ] annotations stronger typing - support for conjur based annotations
- [ ] support [inclusion of other yaml](https://github.com/cyberark/conjur-policy-parser/blob/master/spec/round-trip/yaml/include.yml) files

On another layer build on top of the parser
- [ ] restricted_to support with CIDR validation
- [ ] validate relative and absolute paths
- [ ] dependency [order resolution](https://github.com/cyberark/conjur-policy-parser/blob/master/spec/resolver-fixtures/dependency-order.yml)

To be confirmed
- [ ] support for [gidnumber / uidnumber](https://github.com/cyberark/conjur-policy-parser/blob/master/spec/round-trip/yaml/org.yml)
- [ ] https://github.com/cyberark/conjur-policy-parser/blob/master/spec/resolver-fixtures/nop.yml
