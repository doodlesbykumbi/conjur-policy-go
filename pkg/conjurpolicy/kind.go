package conjurpolicy

//go:generate go run github.com/abice/go-enum -t kind.tmpl

// ENUM(policy, variable, user, group, layer, grant, host, delete, permit, deny)
type Kind int

func (t Kind) Tag() string {
	return "!" + t.String()
}
