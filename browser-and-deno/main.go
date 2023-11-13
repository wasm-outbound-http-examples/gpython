package main

import (
	"io"
	"net/http"

	"github.com/go-python/gpython/py"
	_ "github.com/go-python/gpython/stdlib"
)

func main() {
	py.RegisterModule(&py.ModuleImpl{
		Info: py.ModuleInfo{
			Name: "httpgetlib_go",
			Doc:  "Simple http client for GPython",
		},
		Methods: []*py.Method{
			py.MustNewMethod("HttpGet", HttpGet, 0, ""),
		},
	})

	ctx := py.NewContext(py.DefaultContextOpts())
	defer ctx.Close()

	var err error
	const moduleDescr = "Simplest GPython HTTP client module. Just a wrapper for Golang implementation."
	moduleCode := `
import httpgetlib_go

def Get(url):
  return httpgetlib_go.HttpGet(url)
`

	codeModCompiled, err := py.Compile(moduleCode, moduleDescr, py.ExecMode, 0, true)
	if err != nil {
		py.TracebackDump(err)
	}

	_, err = py.RunCode(ctx, codeModCompiled, moduleDescr, "httpclient")
	if err != nil {
		py.TracebackDump(err)
	}

	const runtimeDescr = "<runtime>"
	mainCode := `
import httpclient

print(httpclient.Get('https://httpbin.org/anything'))
`

	codeCompiled, err := py.Compile(mainCode, runtimeDescr, py.ExecMode, 0, true)
	if err != nil {
		py.TracebackDump(err)
	}

	_, err = py.RunCode(ctx, codeCompiled, runtimeDescr, nil)
	if err != nil {
		py.TracebackDump(err)
	}
}

func HttpGet(self py.Object, urlObj py.Object) (py.Object, error) {
	pyStr, _ := urlObj.(py.String)
	url := string(pyStr)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	return py.String(string(body)), err
}
