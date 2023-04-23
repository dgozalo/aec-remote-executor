package compilers

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func RunCompile(code string) (string, error) {
	f, err := createTempJavaFile(code)
	if err != nil {
		return "", errors.Wrap(err, "there was a problem with the temp compilation file")
	}

	cmd := exec.Command("java", f.Name())
	var outbuf, errbuf strings.Builder // or bytes.Buffer

	b, err := cmd.CombinedOutput()
	if err != nil {
		errors.Wrap(err, "cmd.Run() failed")
	}
	fmt.Println("Stdout", outbuf.String())
	fmt.Println("Stderr", errbuf.String())

	return string(b), nil
}

func createTempJavaFile(code string) (*os.File, error) {
	tmpDir, err := os.MkdirTemp("", "*-aec-compiler-java")
	if err != nil {
		return nil, errors.Wrap(err, "error creating the temp dir for compilationn targets")
	}
	f, err := os.Create(filepath.Join(tmpDir, "Main.java"))
	if err != nil {
		return nil, errors.Wrap(err, "there was a problem creating the temporary java file")
	}
	_, err = f.WriteString(code)
	if err != nil {
		return nil, errors.Wrap(err, "error writing code contents to temporary java file")
	}
	return f, nil
}
