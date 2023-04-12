package conjurpolicy

//go:generate enumer -type Kind -trimprefix Kind -transform lower -text -output kind.gen.go

type Kind int

const (
	KindPolicy Kind = iota
	KindVariable
	KindUser
	KindGroup
	KindLayer
	KindGrant
	KindHost
	KindDelete
	KindPermit
	KindDeny
)

func (t Kind) Tag() string {
	return "!" + t.String()
}
