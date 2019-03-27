package bexpr

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateFieldConfigurations(t *testing.T) {
	t.Parallel()

	t.Run("Flat Struct", func(t *testing.T) {
		t.Parallel()

		expected := FieldConfigurations{
			"Int":     &FieldConfiguration{StructFieldName: "Int", CoerceFn: CoerceInt, SupportedOperations: []MatchOperator{MatchEqual, MatchNotEqual}},
			"Int8":    &FieldConfiguration{StructFieldName: "Int8", CoerceFn: CoerceInt8, SupportedOperations: []MatchOperator{MatchEqual, MatchNotEqual}},
			"Int16":   &FieldConfiguration{StructFieldName: "Int16", CoerceFn: CoerceInt16, SupportedOperations: []MatchOperator{MatchEqual, MatchNotEqual}},
			"Int32":   &FieldConfiguration{StructFieldName: "Int32", CoerceFn: CoerceInt32, SupportedOperations: []MatchOperator{MatchEqual, MatchNotEqual}},
			"Int64":   &FieldConfiguration{StructFieldName: "Int64", CoerceFn: CoerceInt64, SupportedOperations: []MatchOperator{MatchEqual, MatchNotEqual}},
			"Uint":    &FieldConfiguration{StructFieldName: "Uint", CoerceFn: CoerceUint, SupportedOperations: []MatchOperator{MatchEqual, MatchNotEqual}},
			"Uint8":   &FieldConfiguration{StructFieldName: "Uint8", CoerceFn: CoerceUint8, SupportedOperations: []MatchOperator{MatchEqual, MatchNotEqual}},
			"Uint16":  &FieldConfiguration{StructFieldName: "Uint16", CoerceFn: CoerceUint16, SupportedOperations: []MatchOperator{MatchEqual, MatchNotEqual}},
			"Uint32":  &FieldConfiguration{StructFieldName: "Uint32", CoerceFn: CoerceUint32, SupportedOperations: []MatchOperator{MatchEqual, MatchNotEqual}},
			"Uint64":  &FieldConfiguration{StructFieldName: "Uint64", CoerceFn: CoerceUint64, SupportedOperations: []MatchOperator{MatchEqual, MatchNotEqual}},
			"Float32": &FieldConfiguration{StructFieldName: "Float32", CoerceFn: CoerceFloat32, SupportedOperations: []MatchOperator{MatchEqual, MatchNotEqual}},
			"Float64": &FieldConfiguration{StructFieldName: "Float64", CoerceFn: CoerceFloat64, SupportedOperations: []MatchOperator{MatchEqual, MatchNotEqual}},
			"Bool":    &FieldConfiguration{StructFieldName: "Bool", CoerceFn: CoerceBool, SupportedOperations: []MatchOperator{MatchEqual, MatchNotEqual}},
			"String":  &FieldConfiguration{StructFieldName: "String", CoerceFn: CoerceString, SupportedOperations: []MatchOperator{MatchEqual, MatchNotEqual}},
		}

		var ttype *testFlatStruct
		fields, err := GenerateFieldConfigurations(ttype)
		require.NoError(t, err)
		validateFieldConfigurations(t, expected, fields)
	})

	t.Run("map[string]bool", func(t *testing.T) {
		t.Parallel()

		expected := FieldConfigurations{
			FieldNameAny: &FieldConfiguration{CoerceFn: CoerceBool, SupportedOperations: []MatchOperator{MatchEqual, MatchNotEqual}},
		}

		var ttype map[string]bool
		fields, err := GenerateFieldConfigurations(ttype)
		require.NoError(t, err)
		validateFieldConfigurations(t, expected, fields)
	})

	t.Run("map[string]interface{}", func(t *testing.T) {
		t.Parallel()

		var ttype map[string]interface{}
		fields, err := GenerateFieldConfigurations(ttype)
		require.NoError(t, err)
		require.Len(t, fields, 0)
	})

	t.Run("map[interface{}]interface{}", func(t *testing.T) {
		t.Parallel()

		var ttype map[interface{}]interface{}
		fields, err := GenerateFieldConfigurations(ttype)
		require.Len(t, fields, 0)
		require.Error(t, err)
		require.EqualError(t, err, "Cannot generate FieldConfigurations for maps with keys that are not strings")
	})

	t.Run("[]map[string]string", func(t *testing.T) {
		t.Parallel()

		var ttype []map[string]string
		fields, err := GenerateFieldConfigurations(ttype)
		require.Len(t, fields, 0)
		require.Error(t, err)
		require.EqualError(t, err, "Invalid top level type - can only use structs, map[string]* or an MatchExpressionEvaluator")
	})

	t.Run("Nested Structs And Maps", func(t *testing.T) {
		t.Parallel()

		expected := FieldConfigurations{
			"Nested": &FieldConfiguration{StructFieldName: "Nested", SubFields: FieldConfigurations{
				"Map": &FieldConfiguration{StructFieldName: "Map", SupportedOperations: []MatchOperator{MatchIn, MatchNotIn, MatchIsEmpty, MatchIsNotEmpty}, SubFields: FieldConfigurations{
					FieldNameAny: &FieldConfiguration{StructFieldName: "", SupportedOperations: []MatchOperator{MatchEqual, MatchNotEqual}},
				}},
				"MapOfStructs": &FieldConfiguration{StructFieldName: "MapOfStructs", SupportedOperations: []MatchOperator{MatchIsEmpty, MatchIsNotEmpty, MatchIn, MatchNotIn}, SubFields: FieldConfigurations{
					FieldNameAny: &FieldConfiguration{StructFieldName: "", SubFields: FieldConfigurations{
						"Foo": &FieldConfiguration{StructFieldName: "Foo", CoerceFn: CoerceInt, SupportedOperations: []MatchOperator{MatchEqual, MatchNotEqual}},
						"Baz": &FieldConfiguration{StructFieldName: "Baz", SupportedOperations: []MatchOperator{MatchEqual, MatchNotEqual}},
					}},
				}},
				"MapInfInf":   &FieldConfiguration{StructFieldName: "MapInfInf", SupportedOperations: []MatchOperator{MatchIsEmpty, MatchIsNotEmpty}},
				"SliceOfInts": &FieldConfiguration{StructFieldName: "SliceOfInts", CoerceFn: CoerceInt, SupportedOperations: []MatchOperator{MatchIn, MatchNotIn, MatchIsEmpty, MatchIsNotEmpty}},
				"SliceOfStructs": &FieldConfiguration{StructFieldName: "SliceOfStructs", SupportedOperations: []MatchOperator{MatchIsEmpty, MatchIsNotEmpty}, SubFields: FieldConfigurations{
					"X": &FieldConfiguration{StructFieldName: "X", CoerceFn: CoerceInt, SupportedOperations: []MatchOperator{MatchEqual, MatchNotEqual}},
					"Y": &FieldConfiguration{StructFieldName: "Y", CoerceFn: CoerceInt, SupportedOperations: []MatchOperator{MatchEqual, MatchNotEqual}},
				}},
				"SliceOfMapInfInf": &FieldConfiguration{StructFieldName: "SliceOfMapInfInf", SupportedOperations: []MatchOperator{MatchIsEmpty, MatchIsNotEmpty}},
			}},
			"TopInt": &FieldConfiguration{StructFieldName: "TopInt", CoerceFn: CoerceInt, SupportedOperations: []MatchOperator{MatchEqual, MatchNotEqual}},
		}

		var ttype *testNestedTypes
		fields, err := GenerateFieldConfigurations(ttype)
		require.NoError(t, err)
		validateFieldConfigurations(t, expected, fields)

	})
}

/*
func TestReflectEvaluation(t *testing.T) {
	t.Parallel()

	type expTest struct {
		expression string
		result     bool
		err        string
	}

	type testCase struct {
		value       interface{}
		expressions []expTest
	}

	tests := map[string]testCase{
		"Flat Struct": {
			testFlatStruct{
				Int:        -1,
				Int8:       -2,
				Int16:      -3,
				Int32:      -4,
				Int64:      -5,
				Uint:       6,
				Uint8:      7,
				Uint16:     8,
				Uint32:     9,
				Uint64:     10,
				Float32:    1.1,
				Float64:    1.2,
				Bool:       true,
				String:     "exported",
				unexported: "unexported",
			},
			[]expTest{
				{expression: "Int == -1", result: true},
				{expression: "Int == -99", result: false},
				{expression: "Int != -1", result: false},
				{expression: "Int != -99", result: true},
				{expression: "Int8 == -2", result: true},
				{expression: "Int8 == -99", result: false},
				{expression: "Int8 != -2", result: false},
				{expression: "Int8 != -99", result: true},
				{expression: "Int16 == -3", result: true},
				{expression: "Int16 == -99", result: false},
				{expression: "Int16 != -3", result: false},
				{expression: "Int16 != -99", result: true},
				{expression: "Int32 == -4", result: true},
				{expression: "Int32 == -99", result: false},
				{expression: "Int32 != -4", result: false},
				{expression: "Int32 != -99", result: true},
				{expression: "Int64 == -5", result: true},
				{expression: "Int64 == -99", result: false},
				{expression: "Int64 != -5", result: false},
				{expression: "Int64 != -99", result: true},
				{expression: "Uint == 6", result: true},
				{expression: "Uint == 99", result: false},
				{expression: "Uint != 6", result: false},
				{expression: "Uint != 99", result: true},
				{expression: "Uint8 == 7", result: true},
				{expression: "Uint8 == 99", result: false},
				{expression: "Uint8 != 7", result: false},
				{expression: "Uint8 != 99", result: true},
				{expression: "Uint16 == 8", result: true},
				{expression: "Uint16 == 99", result: false},
				{expression: "Uint16 != 8", result: false},
				{expression: "Uint16 != 99", result: true},
				{expression: "Uint32 == 9", result: true},
				{expression: "Uint32 == 99", result: false},
				{expression: "Uint32 != 9", result: false},
				{expression: "Uint32 != 99", result: true},
				{expression: "Uint64 == 10", result: true},
				{expression: "Uint64 == 99", result: false},
				{expression: "Uint64 != 10", result: false},
				{expression: "Uint64 != 99", result: true},
				{expression: "Float32 == 1.1", result: true},
				{expression: "Float32 == 9.9", result: false},
				{expression: "Float32 != 1.1", result: false},
				{expression: "Float32 != 9.9", result: true},
				{expression: "Float64 == 1.2", result: true},
				{expression: "Float64 == 9.9", result: false},
				{expression: "Float64 != 1.2", result: false},
				{expression: "Float64 != 9.9", result: true},
				{expression: "Bool == true", result: true},
				{expression: "Bool == false", result: false},
				{expression: "Bool != true", result: false},
				{expression: "Bool != false", result: true},
				{expression: "String == `exported`", result: true},
				{expression: "String == `not-it`", result: false},
				{expression: "String != `exported`", result: false},
				{expression: "String != `not-it`", result: true},
				{expression: "exp in String", result: true},
				{expression: "foo in String", result: false},
				{expression: "`not` not in String", result: true},
				{expression: "`port` not in String", result: false},
				{expression: "unexported == `unexported`", result: false, err: "Invalid selector: \"unexported\""},
			},
		},
		"map[string]map[string]bool": {
			map[string]map[string]bool{
				"foo": {
					"bar": true,
					"baz": false,
				},
				"abc": nil,
			},
			[]expTest{
				{expression: "bar in foo", result: true},
				{expression: "arg in foo", result: false},
				{expression: "arg not in foo", result: true},
				{expression: "baz not in foo", result: false},
				{expression: "foo is empty", result: false},
				{expression: "foo is not empty", result: true},
				{expression: "abc is empty", result: true},
				{expression: "abc is not empty", result: false},
				{expression: "foo in abc", result: false},
				{expression: "foo not in abc", result: true},
				{expression: "foo.bar == true", result: true},
				{expression: "foo.bar == false", result: false},
				{expression: "foo.baz == false", result: true},
				{expression: "foo.baz == true", result: false},
				{expression: "foo.bar != true", result: false},
				{expression: "foo.bar != false", result: true},
				{expression: "foo.baz != false", result: false},
				{expression: "foo.baz != true", result: true},
				{expression: "foo.bar.baz == 3", result: false, err: "Value at selector \"foo.bar\" with type bool does not support nested field selection"},
			},
		},
		"Nested Structs and Maps": {
			testNestedTypes{
				Nested: testNestedLevel1{
					Map: map[string]string{
						"foo": "bar",
						"bar": "baz",
						"abc": "123",
					},
					MapOfStructs: map[string]testNestedLevel2_1{
						"one": testNestedLevel2_1{
							Foo: 42,
							bar: "unexported",
							Baz: "exported",
						},
						"two": testNestedLevel2_1{
							Foo: 77,
							bar: "unexported",
							Baz: "consul",
						},
					},
					SliceOfInts: []int{1, 3, 5, 7, 9},
					SliceOfStructs: []testNestedLevel2_2{
						testNestedLevel2_2{
							X: 1,
							Y: 2,
							z: 10,
						},
						testNestedLevel2_2{
							X: 3,
							Y: 5,
							z: 10,
						},
					},
					SliceOfMapInfInf: []map[interface{}]interface{}{
						map[interface{}]interface{}{
							1: 2,
						},
					},
				},
				TopInt: 5,
			},
			[]expTest{
				{expression: "Nested.Map.foo == bar", result: true},
				{expression: "Nested.Map.foo == baz", result: false},
				{expression: "Nested.Map is not empty", result: true},
				{expression: "Nested.Map is not empty", result: true},
				{expression: "Nested.Map contains foo and Nested.Map contains bar", result: true},
				{expression: "Nested.Map contains nope", result: false},
				{expression: "Nested.Map contains bar", result: true},
				{expression: "Nested.Map.bar == `bazel`", result: false},
				{expression: "TopInt != 0", result: true},
				{expression: "Nested.Map contains nope or (Nested.Map contains bar and Nested.Map.bar == `bazel`) or TopInt != 0", result: true},
				{expression: "Nested.MapOfStructs.one.Foo == 42", result: true},
				{expression: "Nested.MapOfStructs is empty or (Nested.SliceOfInts contains 7 and 9 in Nested.SliceOfInts)", result: true},
				{expression: "Nested.SliceOfStructs.X == 1", result: true},
				{expression: "Nested.SliceOfStructs.Y == 4", result: false},
				{expression: "Nested.Map.notfound == 4", result: false, err: "Invalid selector - Nested.Map.notfound: key not found in map"},
				{expression: "Map in Nested", result: false, err: "Cannot perform in/contains operations on type struct for selector: \"Nested\""},
				{expression: "Nested.MapInfInf.foo == 4", result: false, err: "Invalid map key type for selector: \"Nested.MapInfInf.foo\" - interface"},
				{expression: "Nested.SliceOfMapInfInf.foo == 4", result: false, err: "Invalid map key type for selector: \"Nested.SliceOfMapInfInf.foo\" - interface"},
			},
		},
	}

	for name, tcase := range tests {
		// capture these values in the closure
		name := name
		tcase := tcase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			for i, expressionTest := range tcase.expressions {
				// capture these values in the closure
				expressionTest := expressionTest
				t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
					t.Parallel()

					expr, err := Create(expressionTest.expression, nil)
					require.NoError(t, err)

					match, err := expr.Evaluate(tcase.value)
					if expressionTest.err != "" {
						require.Error(t, err)
						require.EqualError(t, err, expressionTest.err)
					} else {
						require.NoError(t, err)
					}
					require.Equal(t, expressionTest.result, match)
				})
			}
		})
	}
}
*/