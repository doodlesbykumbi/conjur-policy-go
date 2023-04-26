package conjurpolicy

//go:generate go run github.com/abice/go-enum --marshal

// ENUM(read, execute, update, create)
type Privilege int
