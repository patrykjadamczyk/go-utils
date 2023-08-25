package ignore

// Ignore package imported and not used
func IKnowThatThisPackageIsNotUsed(x ...interface{}) {}

// Ignore variable declared and not used
func IKnowThatThisVariableIsNotUsed(x ...interface{}) {}

// Ignore package imported and not used for debugging packages
func PackagesUsedInDebugging(x ...interface{}) {}

