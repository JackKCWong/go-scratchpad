package core

import (
	"testing"

	expect "github.com/stretchr/testify/require"
)

func TestScratchpad_FixImport(t *testing.T) {
	draft := boilerplate(`fmt.Println("hello world")`)
	code, err := fixImports(draft)

	expect.Nil(t, err)
	expect.Contains(t, code, `import "fmt"`)
}

func TestScratchpad_CompileToWasm(t *testing.T) {
	draft := boilerplate(`fmt.Println("hello world")`)
	code, err := fixImports(draft)
	expect.Nil(t, err)

	wasm, err := compileToWasm(code)
	expect.Nil(t, err)
	expect.NotEmpty(t, wasm)
}
