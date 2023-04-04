package main

import (
	"fmt"
	"os"

	"github.com/gookit/goutil/dump"
	"github.com/gookit/goutil/jsonutil"
	"github.com/hashicorp/go-bexpr"
	"github.com/hashicorp/go-bexpr/grammar"
)

var (
	canonicalDocName = "coordination-of-benefits.v4"

	keyJsonStr = `{
		"individualIdentifier": {
			"string": "cdb:4:144667964:CO:RAM0507494677209000"
		}
	}`

	valueJsonStr = `{
		"individualIdentifier": {
			"string": "cdb:4:144667964:CO:RAM0507494677209000"
		},
		"security": {
			"com.optum.exts.eligibility.model.common.Security": {
				"securityPermission": {
					"array": [
						{
							"securityPermissionValue": {
								"string": "0"
							}
						},
						{
							"securityPermissionValue": {
								"string": "1"
							}
						},
						{
							"securityPermissionValue": {
								"string": 2
							}
						}
					]
				},
				"securityPermissionAny": null,
				"securitySourceSystemCode": {
					"string": "cdb"
				},
				"securityAlt1SourceSystemCode": {
					"string": "CDB"
				},
				"securityAlt2SourceSystemCode": {
					"string": "cdb  "
				},
				"securityAlt3SourceSystemCode": {
					"string": "  cdb"
				},
				"securityAlt4SourceSystemCode": {
					"string": "  cdb  "
				}
			}
		}
	}`

	testJsonStr = `{
		"a" : {
			"string" : "1"
		}
	}`
)

func main() {
	var valueJsonMap map[string]interface{}

	// err := jsonutil.DecodeString(testJsonStr, &valueJsonMap)
	err := jsonutil.DecodeString(valueJsonStr, &valueJsonMap)
	if err != nil {
		fmt.Printf("unable to decode the json string: %v\n", err)
	}
	dump.Println(valueJsonMap)

	// var valueJsonBytes []byte
	// valueJsonBytes, err = jsonutil.Encode(valueJsonMap)
	// if err != nil {
	// 	fmt.Printf("unable to encode the json map to bytes: %v\n", err)
	// }
	// // fmt.Printf("unable to encode the json map to bytes: %v\n", err)
	// dump.Println(string(valueJsonBytes))

	// expressions := []string{
	// 	"/individualIdentifier/string == cdb:4:144667964:CO:RAM0507494677209000",
	// }

	expressions := []string{
		// "a.string == 1",
		// "individualIdentifier.string == `cdb:4:144667964:CO:RAM0507494677209000`",	//working
		// `"individualIdentifier.string" == "cdb:4:144667964:CO:RAM0507494677209000"`,
		// "individualIdentifier.string == \"cdb:4:144667964:CO:RAM0507494677209000\"",
		// `"/individualIdentifier/string == cdb:4:144667964:CO:RAM0507494677209000"`,
		// `"/individualIdentifier/string" == "cdb:4:144667964:CO:RAM0507494677209000"`, //working

		// `/security/com.optum.exts.eligibility.model.common.Security/securitySourceSystemCode/string == cdb`,
		// `"/security/com.optum.exts.eligibility.model.common.Security/securitySourceSystemCode/string" == "cdb"`, //working
		// `"/security/com.optum.exts.eligibility.model.common.Security/securitySourceSystemCode/string" == cdb`, //working
		// `"/security/com.optum.exts.eligibility.model.common.Security/securitySourceSystemCode/string == cdb"`,

		// `"/security/com.optum.exts.eligibility.model.common.Security/securityAlt1SourceSystemCode/string" == "cdb"`,
		// `"/security/com.optum.exts.eligibility.model.common.Security/securityAlt1SourceSystemCode/string" == "CDB"`, //working
		// `"/security/com.optum.exts.eligibility.model.common.Security/securityAlt2SourceSystemCode/string" == "cdb  "`, //working
		// never considers spaces either in left side or right side of string value
		// `"/security/com.optum.exts.eligibility.model.common.Security/securityAlt2SourceSystemCode/string" matches "cdb"`, //working

		// `"/security/com.optum.exts.eligibility.model.common.Security/securityAlt3SourceSystemCode/string" == "  cdb"`, //working
		// trim spaces of string value either side
		// `"/security/com.optum.exts.eligibility.model.common.Security/securityAlt3SourceSystemCode/string" matches "cdb"`, //working
		// `"/security/com.optum.exts.eligibility.model.common.Security/securityAlt3SourceSystemCode/string" matches "cdb1"`,

		// `"/security/com.optum.exts.eligibility.model.common.Security/securityAlt4SourceSystemCode/string" == "  cdb  "`, //working
		// trim spaces of string value either side
		// `"/security/com.optum.exts.eligibility.model.common.Security/securityAlt4SourceSystemCode/string" matches "cdb"`, //working

		// `
		// "/security/com.optum.exts.eligibility.model.common.Security/securityAlt1SourceSystemCode/string" matches "cdb" or //working
		// "/security/com.optum.exts.eligibility.model.common.Security/securityAlt1SourceSystemCode/string" matches "CDB"	//working
		// `,

		// string value starts with CD only one occurence
		// `
		// "/security/com.optum.exts.eligibility.model.common.Security/securityAlt1SourceSystemCode/string" matches "/cdb/" or
		// "/security/com.optum.exts.eligibility.model.common.Security/securityAlt1SourceSystemCode/string" matches "^CD{1}"
		// `,

		// string value ends with DB only one occurence
		// `"/security/com.optum.exts.eligibility.model.common.Security/securityAlt1SourceSystemCode/string" matches "DB{1}$"`,

		// `
		// "/security/com.optum.exts.eligibility.model.common.Security/securityAlt1SourceSystemCode/string" contains "cdb" or
		// "/security/com.optum.exts.eligibility.model.common.Security/securityAlt1SourceSystemCode/string" contains "CDB"
		// `,

		// `
		// "/security/com.optum.exts.eligibility.model.common.Security/securityAlt1SourceSystemCode/string" contains "cd" or
		// "/security/com.optum.exts.eligibility.model.common.Security/securityAlt1SourceSystemCode/string" contains "CD"
		// `,

		// `
		// "/security/com.optum.exts.eligibility.model.common.Security/securityAlt1SourceSystemCode/string" contains "db" or
		// "/security/com.optum.exts.eligibility.model.common.Security/securityAlt1SourceSystemCode/string" contains "DB"
		// `,

		// `"/security/com.optum.exts.eligibility.model.common.Security/securityAlt4SourceSystemCode/string" matches "cdb"`,

		// `"/security/com.optum.exts.eligibility.model.common.Security/securityPermission/array/0/securityPermissionValue/string" == "0"`, //working
		// `"/security/com.optum.exts.eligibility.model.common.Security/securityPermission/array/1/securityPermissionValue/string" == "1"`,	//working
		// `"/security/com.optum.exts.eligibility.model.common.Security/securityPermission/array/2/securityPermissionValue/string" == "2"`,	//working
		`"/security/com.optum.exts.eligibility.model.common.Security/securityPermission/array/2/securityPermissionValue/string" == 2`, //working

		// `"/security/com.optum.exts.eligibility.model.common.Security/securityPermission/array/1/securityPermissionValue/string" in "1"`,
		// `[0,1] in "/security/com.optum.exts.eligibility.model.common.Security/securityPermission/array/1/securityPermissionValue/string"`,
		// `0 in "/security/com.optum.exts.eligibility.model.common.Security/securityPermission/array/0/securityPermissionValue/string"`, //working
		// `0 in "/security/com.optum.exts.eligibility.model.common.Security/securityPermission/array/1/securityPermissionValue/string"`, //working
		// `"1" in "/security/com.optum.exts.eligibility.model.common.Security/securityPermission/array/1/securityPermissionValue/string"`, //working
		// `1 in "/security/com.optum.exts.eligibility.model.common.Security/securityPermission/array/1/securityPermissionValue/string"`, //working

		// `1 in "/security/com.optum.exts.eligibility.model.common.Security/securityPermission/array/0..1/securityPermissionValue/string"`,
		// `1 in "/security/com.optum.exts.eligibility.model.common.Security/securityPermission/array/[0..1]/securityPermissionValue/string"`,
	}
	for i, expression := range expressions {
		ast, err := grammar.Parse(fmt.Sprintf("Expression %d", i), []byte(expression))

		if err != nil {
			fmt.Println(err)
		} else {
			ast.(grammar.Expression).ExpressionDump(os.Stdout, "   ", 1)
		}

		eval, err := bexpr.CreateEvaluator(expression)

		if err != nil {
			fmt.Printf("Failed to create evaluator for expression %q: %v\n", expression, err)
			continue
		}

		result, err := eval.Evaluate(valueJsonMap)
		// result, err := eval.Evaluate(valueJsonBytes)
		if err != nil {
			fmt.Printf("Failed to run evaluation of expression %q: %v\n", expression, err)
			continue
		}

		fmt.Printf("Result of expression %q evaluation: %t\n", expression, result)
	}
}
