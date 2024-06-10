package base

type IgnoreStruct struct {}

// Ignore package imported and not used
func IKnowThatThisPackageIsNotUsed(x ...any) (r IgnoreStruct) { return }

// Ignore variable declared and not used
func IKnowThatThisVariableIsNotUsed(x ...any) (r IgnoreStruct) { return }

// Ignore package imported and not used for debugging packages
func PackagesUsedInDebugging(x ...any) (r IgnoreStruct) { return }

// Ignore package imported and not used
func PossiblyUsingThis(x ...any) (r IgnoreStruct) { return }
