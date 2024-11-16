package converter

// goverter:converter
// goverter:output:file ./converter/converter.gen.go
// goverter:output:package :converter
// goverter:useZeroValueOnPointerInconsistency
// goverter:ignoreUnexported
// goverter:matchIgnoreCase
//
//go:generate goverter gen ./
type RepoConverter interface {
}
