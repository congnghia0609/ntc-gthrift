/**
 *
 * @author nghiatc
 * @since May 31, 2020
 * @thrift 0.13.0
 */

package ntserver

import (
	"crypto/tls"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/congnghia0609/ntc-gconf/nconf"
	log "github.com/sirupsen/logrus"
)

const NTSPrefix = ".ntserver."

type NTServer struct {
	Name string
	Address string
	IsSSL bool
	Server *thrift.TSimpleServer
}

func NewNTServer(name string, processor thrift.TProcessor) *NTServer {
	if len(name) == 0 {
		return nil
	}
	c := nconf.GetConfig()
	addr := c.GetString(name + NTSPrefix + "address")
	isSSL := c.GetBool(name + NTSPrefix + "is_ssl")
	fmt.Printf("NewNTServer[%s] isSSL: %v\n", name, isSSL)

	// BinaryProtocol & FramedTransport
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	transportFactory := thrift.NewTTransportFactory()
	transportFactory = thrift.NewTFramedTransportFactory(transportFactory)
	// Transport
	var transport thrift.TServerTransport
	var err error
	if isSSL {
		cfg := new(tls.Config)
		certFile := c.GetString(name + NTSPrefix + "cert_file") //"ssl/server.crt"
		keyFile := c.GetString(name + NTSPrefix + "key_file")   //"ssl/server.pem"
		if cert, err := tls.LoadX509KeyPair(certFile, keyFile); err == nil {
			cfg.Certificates = append(cfg.Certificates, cert)
		} else {
			log.Fatalf("NewNTServer LoadX509KeyPair fail: %v\n", err)
			return nil
		}
		transport, err = thrift.NewTSSLServerSocket(addr, cfg)
	} else {
		transport, err = thrift.NewTServerSocket(addr)
	}

	if err != nil {
		log.Fatalf("NewNTServer init transport fail: %v\n", err)
		return nil
	}
	fmt.Printf("%T\n", transport)

	server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)
	return &NTServer{Name: name, Address: addr, IsSSL: isSSL, Server: server}
}

func (nts *NTServer) Start() error {
	fmt.Println("Starting the simple server on:", nts.Address)
	return nts.Server.Serve()
}
