package types

var (
	errWrongFunctionParameters = `function "%s" was called with incorrect parameters. The expected parameter are %s, but the provided parameter were %s.\n` +
		`Double-check the function's documentation or usage guide to ensure that the input parameters match the expected types`
)
