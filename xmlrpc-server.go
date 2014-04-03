package main

import (
	"github.com/divan/gorilla-xmlrpc/xml"
	"github.com/gorilla/rpc"
	"log"
	"net/http"
)

type MessageArgs struct{
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

func main() {
	RPC := rpc.NewServer()
	xmlrpcCodec := xml.NewCodec()
	RPC.RegisterCodec(xmlrpcCodec, "text/xml")
	RPC.RegisterService(new(MessageService), "")
	http.Handle("/", RPC)

	log.Println("Starting XML-RPC server on localhost:8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}
