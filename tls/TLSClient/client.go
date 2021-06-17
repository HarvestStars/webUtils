package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
)

var rootCertPath = "../rootCA/root.crt"
var clientCertPath = "./client.crt"
var clientKeyPath = "./client.key"

func main() {
	log.SetFlags(log.Lshortfile)
	//这里读取的是根证书
	buf, err := ioutil.ReadFile(rootCertPath)
	if err != nil {
		return
	}

	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(buf)

	//加载客户端证书
	cert, err := tls.LoadX509KeyPair(clientCertPath, clientKeyPath)
	if err != nil {
		return
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      pool,
	}

	conn, err := tls.Dial("tcp", "localhost:8000", tlsConfig)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()
	n, err := conn.Write([]byte("hello\n"))
	if err != nil {
		log.Println(n, err)
		return
	}
	bufRec := make([]byte, 100)
	n, err = conn.Read(bufRec)
	if err != nil {
		log.Println(n, err)
		return
	}
	println(string(bufRec[:n]))
}
