module nlp

go 1.25

// Risks of using 3rd party packages:
// - Security (ignorance, or intentional).
// - Bugs.
// - Compatibility (API changes).
// - Legal (license).
// - Might be gone/deprecated/removed from the Internet.

require (
	github.com/BurntSushi/toml v1.5.0
	github.com/stretchr/testify v1.11.1
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/exp/typeparams v0.0.0-20231108232855-2478ac86f678 // indirect
	golang.org/x/mod v0.23.0 // indirect
	golang.org/x/sync v0.11.0 // indirect
	golang.org/x/tools v0.30.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	honnef.co/go/tools v0.6.1 // indirect
)

// Run this tool by running the following command in the terminal:
// "go tool staticcheck ."
tool honnef.co/go/tools/cmd/staticcheck
