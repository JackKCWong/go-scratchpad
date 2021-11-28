# go-scratchpad

A go-playground like webapp to run Go SNIPPETs in YOUR BROWSER using WASM.

Basically it wraps your code inside a `main()` and pipe it thru `goimports`, then compile it to WASM to run in your browser.


## installation

The host machine must have goimports installed and available in the PATH.

```
go get golang.org/x/tools/cmd/goimports
```
