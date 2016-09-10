// temp
package main

import (
	"crypto/tls"
	//"crypto/x509"
	"fmt"
	//"io/ioutil"
	//"net"

	"flag"
	"log"
	"net/http"
	"net/url"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:9200", "http service address")

func main() {

	//u := url.URL{Scheme: "wss", Host: *addr, Path: "/ws"}

	//	URL := "wss://localhost:9200/ws"
	URL := "wss://127.0.0.1:9200/ws"
	u, err := url.Parse(URL)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//	rawConn, err := net.Dial("tcp", u.Host)
	//	if err != nil {
	//		fmt.Println()
	//		return
	//	}

	wsHeaders := http.Header{
		"Origin": {u.Host},
		// your milage may differ
		"Sec-WebSocket-Extensions": {"permessage-deflate; client_max_window_bits, x-webkit-deflate-frame"},
	}

	//	CA_Pool := x509.NewCertPool()
	//	severCert, err := ioutil.ReadFile("./algo.crt")
	//	if err != nil {
	//		fmt.Println("Could not load server certificate!")
	//		return
	//	}
	//	if ok := CA_Pool.AppendCertsFromPEM(severCert); !ok {
	//		fmt.Println("Not ok!")
	//	}

	//config := tls.Config{RootCAs: CA_Pool /*, InsecureSkipVerify: true*/}
	config := tls.Config{RootCAs: nil, InsecureSkipVerify: true}
	c, err := tls.Dial("tcp", u.Host, &config)
	//c, err := net.Dial("tcp", u.Host)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//	fmt.Println(URL)
	ws, _, err := websocket.NewClient(c, u, wsHeaders, 1024, 1024)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	log.Printf("connecting to %s", u.String())

	//	ws, _, err := websocket.DefaultDialer.Dial(u.String(), wsHeaders)
	//	if err != nil {
	//		log.Fatal("dial:", err)
	//	}

	con := &connection{
		send: make(chan []byte, 256),
		ws:   ws} //300315 Simon - #865
	fmt.Println("new connection %#v", con)

	go con.writePump()
	go con.readPump()
	//	if err := con.LoginToProxy(); err != nil {
	//		printForDebug("error Login to proxy " + err.Error())
	//	} else {
	//		go con.sendEventsToProxy()
	//	}
	//return nil

}
