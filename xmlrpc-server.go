package main

import (
	"flag"
	"fmt"
	"github.com/divan/gorilla-xmlrpc/xml"
	"github.com/gorilla/rpc"
	"log"
	"net/http"
	"os"
)

type MessageArgs struct {
	MessageBody string
}

type MessageReply struct {
	ResponseBody string
}

type MessageService struct{}

func (h *MessageService) Send(r *http.Request, args *MessageArgs, reply *MessageReply) error {
	log.Printf("Received: %s\n", args.MessageBody)
	reply.ResponseBody = "Thank you for the message."
	return nil
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s -b [address] -p [port]\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	bind := flag.String("b", "0.0.0.0", "bind address")
	port := flag.Int("p", 8000, "port")
	flag.Usage = usage
	flag.Parse()

	RPC := rpc.NewServer()
	xmlrpcCodec := xml.NewCodec()
	RPC.RegisterCodec(xmlrpcCodec, "text/xml")
	RPC.RegisterService(new(MessageService), "")
	http.Handle("/", RPC)

	svradr := fmt.Sprintf("%s:%d", *bind, *port)
	log.Printf("Starting XML-RPC server on %s\n", svradr)
	err := http.ListenAndServe(svradr, nil)
	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}
