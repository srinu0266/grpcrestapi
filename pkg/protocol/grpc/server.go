package grpc

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"

	v1 "github.com/srinu0266/grpcrestapi/pkg/service/v1"
	"google.golang.org/grpc"
)

func RunServer(ctx context.Context,v1API v1.ToDoServiceServer,port string ) error {
	listen,err:=net.Listen("tcp",":"+port)
	if err!=nil{
		return err
	}

	server:=grpc.NewServer()

	v1.RegisterToDoServiceServer(server,v1API)

	c:=make(chan os.Signal,1)

	signal.Notify(c,os.Interrupt)

	go func ()  {
		for range c{
			log.Println("shutting down grpc server.........")

			server.GracefulStop()

			<-ctx.Done()
		}
	}()

	log.Panicln("Starting GRPC Server...........")

	return server.Serve(listen)
}
