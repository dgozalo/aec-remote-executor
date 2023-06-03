package worker

type ExecutionResult struct {
	Stdout          string
	Stderr          string
	RawTestsResults string
	TestsResults    []TestResult
}

type TestResult struct {
	TestName string
	Expected string
	Actual   string
	Passed   bool
}
