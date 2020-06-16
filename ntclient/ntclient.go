/**
 *
 * @author nghiatc
 * @since May 31, 2020
 * @thrift 0.13.0
 */

package ntclient

import (
	"crypto/tls"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/congnghia0609/ntc-gconf/nconf"
)

const NTCPrefix = ".ntclient."

type NTClient struct {
	Name      string
	Address   string
	IsSSL     bool
	Transport thrift.TTransport
	Client    *thrift.TStandardClient
}

func NewNTClient(name string) *NTClient {
	if len(name) == 0 {
		return nil
	}
	c := nconf.GetConfig()
	addr := c.GetString(name + NTCPrefix + "address")
	isSSL := c.GetBool(name + NTCPrefix + "is_ssl")
	fmt.Printf("NewNTClient[%s] isSSL: %v\n", name, isSSL)

	// BinaryProtocol & FramedTransport
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	transportFactory := thrift.NewTTransportFactory()
	transportFactory = thrift.NewTFramedTransportFactory(transportFactory)
	// Transport
	var transport thrift.TTransport
	var err error
	if isSSL {
		cfg := new(tls.Config)
		cfg.InsecureSkipVerify = true
		transport, err = thrift.NewTSSLSocket(addr, cfg)
	} else {
		transport, err = thrift.NewTSocket(addr)
	}
	if err != nil {
		fmt.Println("Error opening socket:", err)
		return nil
	}
	transport, err = transportFactory.GetTransport(transport)
	if err != nil {
		fmt.Println("NewNTClient init transport Error:", err)
		return nil
	}
	//defer transport.Close()
	if err := transport.Open(); err != nil {
		fmt.Println("NewNTClient Open transport Error:", err)
		return nil
	}
	iprot := protocolFactory.GetProtocol(transport)
	oprot := protocolFactory.GetProtocol(transport)
	client := thrift.NewTStandardClient(iprot, oprot)
	return &NTClient{Name: name, Address: addr, IsSSL: isSSL, Transport: transport, Client: client}
}

func (ntc *NTClient) Close() {
	if ntc != nil && ntc.Transport != nil {
		ntc.Transport.Close()
	}
}

func (ntc *NTClient) IsOpen() bool {
	return ntc.Transport.IsOpen()
}
