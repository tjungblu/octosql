package functions

import (
	"time"

	"github.com/cube2222/octosql/octosql"
	"github.com/cube2222/octosql/physical"
)

var Equals = []physical.FunctionDescriptor{
	// TODO: Specializations for concrete primitive types.
	{
		ArgumentTypes: []octosql.Type{octosql.Any, octosql.Any},
		OutputType:    octosql.Boolean,
		Function: func(values []octosql.Value) (octosql.Value, error) {
			return octosql.NewBoolean(values[0].Compare(values[1]) == 0), nil
		},
	},
}
var FromUnix = []physical.FunctionDescriptor{
	// TODO: Specializations for concrete primitive types.
	{
		ArgumentTypes: []octosql.Type{octosql.Int},
		OutputType:    octosql.Time,
		Function: func(values []octosql.Value) (octosql.Value, error) {
			return octosql.NewTime(time.Unix(int64(values[0].Int), 0)), nil
		},
	},
	{
		ArgumentTypes: []octosql.Type{octosql.Float},
		OutputType:    octosql.Time,
		Function: func(values []octosql.Value) (octosql.Value, error) {
			return octosql.NewTime(time.Unix(int64(values[0].Float), 0)), nil
		},
	},
}
