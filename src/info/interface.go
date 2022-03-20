package info

// Intended to use the mocks to build some tests, but honestly I just got a bit busy, and it's
// not a difficult refactor to incorporate the interface here.

// Additionally, this interface allows us to determine if it's an IP Address, Hash, OR Domain, should this code
// be intended to be extended later.
type Info interface {
	ValidateJsonBody() bool
	GenerateResults() (string, error)
	ObtainResults() error
}
