package worker

// ExecutionResult is the struct that contains the result of the execution
type ExecutionResult struct {
	Stdout          string
	Stderr          string
	RawTestsResults string
	TestsResults    []TestResult
}

// TestResult is the struct that contains the result of an assignment test
type TestResult struct {
	TestName string
	Expected string
	Actual   string
	Passed   bool
}
