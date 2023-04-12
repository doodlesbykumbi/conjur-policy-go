package conjurpolicy

//go:generate enumer -type Privilege -trimprefix Privilege -transform lower -yaml -output privilege.gen.go
type Privilege int

const (
	PrivilegeRead Privilege = iota
	PrivilegeExecute
	PrivilegeUpdate
	PrivilegeCreate
)
