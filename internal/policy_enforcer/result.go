package policy_enforcer

type ResultReport struct {
	Description string
	// Field contains the checked policy field
	Field string
	// Sections holds section (provider, variable or resource) where the success occurred
	Section string
}

type ResultError struct {
	ResultReport
}
type ResultSuccess struct {
	ResultReport
}

type Result struct {
	// errors contains policies that have not been followed
	errors []ResultError
	// success contains policies that have been successfully followed
	success []ResultSuccess
}

func (v *Result) Valid() bool {
	return len(v.errors) == 0
}

func (v *Result) Errors() []ResultError {
	return v.errors
}
func (v *Result) AddError(err ResultError) {
	v.errors = append(v.errors, err)
}

func (v *Result) Success() []ResultSuccess {
	return v.success
}
func (v *Result) AddSuccess(err ResultSuccess) {
	v.success = append(v.success, err)
}
