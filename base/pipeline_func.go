package base

import "github.com/patrykjadamczyk/go-utils/utils"

// Pipeline Function
// Call like Pipe(1, f1, f2, f3) should work similiarly to 1 |> f1 |> f2 |> f3
func Pipe(initialArgument any, functions ...any) any {
	var result = initialArgument
	for _, function := range functions {
		tempResult := utils.CallFunc(function, result)
		tempResult2 := utils.ReflectValueToValue(tempResult...)
		if len(tempResult2) == 1 {
			result = tempResult2[0]
		} else {
			result = tempResult2
		}
	}
	return result
}

// Pipeline Function
// This variant is have enforced return type
func PipeTypedReturn[RT any](initialArgument any, functions ...any) RT {
	return EnsureType[RT](Pipe(initialArgument, functions...))
}

// Pipeline Function
// This variant is have enforced return type and initialArgument
func PipeTypedArgReturn[IAT any, RT any](initialArgument IAT, functions ...any) RT {
	return EnsureType[RT](Pipe(initialArgument, functions...))
}

// Pipeline function
// This variant is fully typed. Note this function don't let you change type of pipeline value at any point
func PipeTyped[VT any](
	initialArgument VT,
	functions ...func(VT) VT,
) VT {
	var result = initialArgument
	for _, function := range functions {
		result = function(result)
	}
	return result
}

// Pipeline function
// This variant is fully typed for return type and initial argument and uses functions fun(any) -> any
// There in these functions you can use EnsureType for correcting types of arguments
func PipeTypedAny[IAT any, RT any](initialArgument IAT, functions ...AnyFunc) RT {
	var result = ToAny(initialArgument)
	for _, function := range functions {
		result = function(result)
	}
	return EnsureType[RT](result)
}
