package java

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/pkg/errors"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"
)

//go:embed TestRunner.java.tmpl
var JavaTemplateFile string

// TemplateData is the struct that contains the data to be used in the code template
type TemplateData struct {
	SubmittedCode  string
	TestRunnerCode string
}

// RunCompile receives the code submitted by the user and the test runner code defined for this assignment
// it then processes the Go template containing the Java file with the given code snippets.
// It then compiles the Java code and runs the test runner code, which creates a file with the test results
// that is then read and returned
func RunCompile(submittedCode string, testRunnerCode string) (string, string, string, error) {
	f, tmpDir, err := createTempJavaFile(submittedCode, testRunnerCode)
	if err != nil {
		return "", "", "", errors.Wrap(err, "there was a problem with the temp compilation file")
	}

	testFile, err := createTestReportFile(tmpDir)
	if err != nil {
		return "", "", "", errors.Wrap(err, "there was a problem creating the test report file")
	}

	cmd := exec.Command("java", f.Name(), testFile.Name())
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()
	if err != nil {
		errors.Wrap(err, "cmd.Run() failed")
	}
	testFileResults, err := readTestReportFile(testFile)
	if err != nil {
		return "", "", "", errors.Wrap(err, "there was a problem reading the test report file")
	}
	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	fmt.Printf("out:\n%s\nerr:\n%s\n", outStr, errStr)

	return string(stdout.Bytes()), string(stderr.Bytes()), testFileResults, nil
}

func readTestReportFile(testFile *os.File) (string, error) {
	testResults, err := os.ReadFile(testFile.Name())
	if err != nil {
		return "", errors.Wrap(err, "there was a problem reading the test report file")
	}
	return string(testResults), nil
}

func createTestReportFile(tmpDir string) (*os.File, error) {
	testFile, err := os.Create(filepath.Join(tmpDir, "TEST-TestRunner.txt"))
	if err != nil {
		return nil, errors.Wrap(err, "there was a problem creating the test report file")
	}
	return testFile, nil
}

func createTempJavaFile(submittedCode string, testRunnerCode string) (*os.File, string, error) {
	tmpDir, err := os.MkdirTemp("", "*-aec-compiler-java")
	if err != nil {
		return nil, "", errors.Wrap(err, "error creating the temp dir for compilationn targets")
	}

	tmpl := template.Must(template.New("java").Parse(JavaTemplateFile))
	if err != nil {
		return nil, "", errors.Wrap(err, "error parsing the java template")
	}

	f, err := os.Create(filepath.Join(tmpDir, "RunnerTest.java"))
	if err != nil {
		return nil, "", errors.Wrap(err, "there was a problem creating the temporary java file")
	}

	err = tmpl.Execute(f, TemplateData{SubmittedCode: submittedCode, TestRunnerCode: testRunnerCode})
	if err != nil {
		return nil, "", errors.Wrap(err, "error executing the java template")
	}

	return f, tmpDir, nil
}
