package testdata

type testStruct struct {
	// some random comment

	ExportedRequired int

	unexportedRequired int

	ExportedOptional   int `exhaustruct:"optional"`
	unexportedOptional int `exhaustruct:"optional"`

	ExportedOptionalType   OptionalType
	unexportedOptionalType OptionalType
}

type OptionalType int

var (
	_unnamed = testStruct{1, 2, 3, 4, 5, 6}
	_named   = testStruct{
		ExportedRequired:       1,
		unexportedRequired:     2,
		ExportedOptional:       3,
		unexportedOptional:     4,
		ExportedOptionalType:   5,
		unexportedOptionalType: 6,
	}
	_unnamedIncomplete = testStruct{1}
	_namedIncomplete1  = testStruct{
		ExportedRequired:     1,
		ExportedOptional:     3,
		ExportedOptionalType: 5,
	}
	_namedIncomplete2 = testStruct{
		ExportedOptional:       3,
		unexportedOptional:     4,
		unexportedOptionalType: 6,
	}
)
