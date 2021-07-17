# ntc-gthrift
ntc-gthrift is a example golang thrift  

## Thrift Server
Example Calculator Server  
```go
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
```

## Thrift Client
Example Calculator Client  
```go
func main() {
	// Init NConf
	InitNConf()

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
```

## Run project
```bash
// Thrift 0.13.0 gen source code for Golang not support Go Module
export GO111MODULE=off

// Gen Source Code
make gen

// Run Server
make server

// Run client
make client
```

## HAProxy Config Load Balancer for Thrift Server
```bash
frontend thrift_fe
	bind *:9000
	mode tcp
	option tcplog
	default_backend thrift_be

backend thrift_be
	mode tcp
	balance roundrobin
	option tcp-check
	server thift-go 127.0.0.1:9090 check
	server thrift-java 127.0.0.1:9091 check
```

