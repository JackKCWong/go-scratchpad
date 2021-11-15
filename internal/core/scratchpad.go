package core

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/go-cmd/cmd"
)

type Scratchpad struct {
	Snippet string
}

func boilerplate(snippet string) string {
	draft := fmt.Sprintf(`
		package main

		func main() {
			%s
		}
	`, snippet)

	return draft
}

func fixImports(goCode string) (string, error) {
	goImports := cmd.NewCmdOptions(cmd.Options{
		Buffered: true,
	}, "goimports")

	status := <-goImports.StartWithStdin(strings.NewReader(goCode))

	if status.Error != nil {
		return "", status.Error
	}

	return strings.Join(status.Stdout, "\n"), nil
}

func compileToWasm(code string) ([]byte, error) {
	scratchFile, err := os.CreateTemp(os.TempDir(), "go-scratchpad-*.go")
	if err != nil {
		return nil, fmt.Errorf("failed to create temp file: %w", err)
	}

	defer os.Remove(scratchFile.Name())
	defer os.Remove(scratchFile.Name() + ".wasm")
	defer scratchFile.Close()

	_, err = io.WriteString(scratchFile, code)

	if err != nil {
		return nil, fmt.Errorf("failed to write to temp file: %w", err)
	}

	goCompile := cmd.NewCmdOptions(cmd.Options{
		Buffered: true,
	}, "go", "build", "-o", scratchFile.Name()+".wasm", scratchFile.Name())

	goCompile.Dir = os.TempDir()
	goCompile.Env = append(os.Environ(),
		"GOOS=js", "GOARCH=wasm")

	status := <-goCompile.Start()

	if status.Error != nil {
		return nil, fmt.Errorf("failed to compile: %w", status.Error)
	}

	if len(status.Stderr) > 0 {
		return nil, fmt.Errorf("failed to compile: %v", status.Stderr)
	}

	return ioutil.ReadFile(scratchFile.Name() + ".wasm")
}
