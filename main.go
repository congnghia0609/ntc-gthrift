/**
 *
 * @author nghiatc
 * @since Jun 16, 2020
 */

package main

import (
	"fmt"
	"github.com/congnghia0609/ntc-gconf/nconf"
	"ntc-gthrift/example/handler"
	"ntc-gthrift/example/thrift/gen-go/tutorial"
	"ntc-gthrift/ntserver"
	"path/filepath"
	"runtime"
)

func InitNConf() {
	_, b, _, _ := runtime.Caller(0)
	wdir := filepath.Dir(b)
	fmt.Println("wdir:", wdir)
	nconf.Init(wdir)
}

func main() {
	// Init NConf
	InitNConf()

	// Start Simple NTServer
	StartSimpleNTServer()
}

func StartSimpleNTServer() {
	name := "tutorial"
	handler := handler.NewCalculatorHandle()
	processor := tutorial.NewCalculatorProcessor(handler)
	nts := ntserver.NewNTServer(name, processor)
	if err := nts.Start(); err != nil {
		fmt.Println("error running server:", err)
	}
}
