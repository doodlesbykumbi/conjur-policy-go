package conjurpolicy

//go:generate go run github.com/abice/go-enum --marshal

// Privilege is an enum representing Conjur type of transactions that are permitted
// ENUM(read, execute, update, create)
type Privilege int
