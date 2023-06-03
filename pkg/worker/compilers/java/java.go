package java

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/pkg/errors"
	"os"
	"os/exec"
	"path/filepath"
)

//go:embed TestRunner.java
var JavaTemplateFile string

func RunCompile(code string) (string, string, string, error) {
	f, tmpDir, err := createTempJavaFile(code)
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

func createTempJavaFile(code string) (*os.File, string, error) {
	// code = string(JavaTemplateFile) + code

	tmpDir, err := os.MkdirTemp("", "*-aec-compiler-java")
	if err != nil {
		return nil, "", errors.Wrap(err, "error creating the temp dir for compilationn targets")
	}

	f, err := os.Create(filepath.Join(tmpDir, "RunnerTest.java"))
	if err != nil {
		return nil, "", errors.Wrap(err, "there was a problem creating the temporary java file")
	}

	_, err = f.WriteString(JavaTemplateFile)
	if err != nil {
		return nil, "", errors.Wrap(err, "error writing code contents to temporary java file")
	}
	return f, tmpDir, nil
}
