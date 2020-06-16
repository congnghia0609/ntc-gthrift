/**
 *
 * @author nghiatc
 * @since Jun 16, 2020
 */

package main

import (
	"context"
	"fmt"
	"github.com/congnghia0609/ntc-gconf/nconf"
	"ntc-gthrift/ntclient"
	"ntc-gthrift/thrift/gen-go/tutorial"
	"path/filepath"
	"runtime"
)

func InitNConf2() {
	_, b, _, _ := runtime.Caller(0)
	wdir := filepath.Dir(b)
	fmt.Println("wdir:", wdir)
	nconf.Init(wdir)
}

func main() {
	// Init NConf
	InitNConf2()

	// Start Simple NTClient
	StartSimpleNTClient()
}

var defaultCtx = context.Background()

func StartSimpleNTClient() error {
	// Init Calculator Client
	name := "tutorial"
	ntc := ntclient.NewNTClient(name)
	defer ntc.Close()
	client := tutorial.NewCalculatorClient(ntc.Client)

	// Call Methods RPC Thrift Server
	client.Ping(defaultCtx)
	fmt.Println("ping()")

	// Add
	sum, _ := client.Add(defaultCtx, 1, 1)
	fmt.Print("1+1=", sum, "\n")

	// Calculate
	work := tutorial.NewWork()
	work.Op = tutorial.Operation_DIVIDE
	work.Num1 = 1
	work.Num2 = 0
	quotiel, err := client.Calculate(defaultCtx, 1, work)
	if err != nil {
		switch v := err.(type) {
		case *tutorial.InvalidOperation:
			fmt.Println("Invalid operation:", v)
		default:
			fmt.Println("Error during operaTION:", err)
		}
	} else {
		fmt.Println("Whoa we can divide by 0 with new value:", quotiel)
	}

	work.Op = tutorial.Operation_SUBTRACT
	work.Num1 = 15
	work.Num2 = 10
	diff, err := client.Calculate(defaultCtx, 1, work)
	if err != nil {
		switch v := err.(type) {
		case *tutorial.InvalidOperation:
			fmt.Println("Invalid operation:", v)
		default:
			fmt.Println("Error during operation:", err)
		}
		return err
	} else {
		fmt.Print("15-10=", diff, "\n")
	}

	// GetStruct
	log, err := client.GetStruct(defaultCtx, 1)
	if err != nil {
		fmt.Println("Unable to get struct:", err)
		return err
	} else {
		fmt.Println("Check log:", log)
	}
	return err
}
