package main

import (
	"bufio"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net"
)

var rootCertPath = "../rootCA/root.crt"
var serverCertPath = "./server.crt"
var serverKeyPath = "./server.key"

func main() {
	log.SetFlags(log.Lshortfile)

	//这里读取的是根证书
	buf, err := ioutil.ReadFile(rootCertPath)
	if err != nil {
		return
	}

	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(buf)

	//加载服务端证书
	cert, err := tls.LoadX509KeyPair(serverCertPath, serverKeyPath)
	if err != nil {
		return
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    pool,
	}

	ln, err := tls.Listen("tcp", ":8000", tlsConfig)
	if err != nil {
		log.Println(err)
		return
	}
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConnection(conn)
	}
}
func handleConnection(conn net.Conn) {
	defer conn.Close()
	r := bufio.NewReader(conn)
	for {
		msg, err := r.ReadString('\n')
		if err != nil {
			log.Println(err)
			return
		}
		println(msg)
		n, err := conn.Write([]byte("world\n"))
		if err != nil {
			log.Println(n, err)
			return
		}
	}
}
