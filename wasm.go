package main

import (
	"fmt"
	"syscall/js"

	"github.com/kaz/go-wrapped-wasm/md"
)

func md2html(args ...js.Value) (interface{}, error) {
	html, _, err := md.Render(args[0].String())
	if err != nil {
		return nil, fmt.Errorf("md.Render failed: %w", err)
	}
	return html, nil
}

func errortest(args ...js.Value) (interface{}, error) {
	return nil, fmt.Errorf("error occured")
}
func panictest(args ...js.Value) (interface{}, error) {
	panic("panic occured")
}

func wrappedFn(inernalFn func(args ...js.Value) (interface{}, error)) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		raise := func(err interface{}) {
			if callback := this.Get("__go_error_callback"); callback.Type() == js.TypeFunction {
				callback.Invoke(js.ValueOf(fmt.Sprintf("%v", err)))
			}
		}

		defer func() {
			if err := recover(); err != nil {
				raise(err)
			}
		}()

		res, err := inernalFn(args...)
		if err != nil {
			raise(err)
			return nil
		}
		return js.ValueOf(res)
	})
}

func main() {
	js.Global().Set("__go_exports", js.ValueOf(map[string]interface{}{
		"md2html":   wrappedFn(md2html),
		"errortest": wrappedFn(errortest),
		"panictest": wrappedFn(panictest),
	}))
	if callback := js.Global().Get("__go_startup_callback"); callback.Type() == js.TypeFunction {
		callback.Invoke()
	}
	<-make(chan interface{}, 0)
}
