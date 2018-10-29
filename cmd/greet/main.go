package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/rugwirobaker/tutorial/greeter/proto"
	"google.golang.org/grpc"
)

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		fmt.Fprintln(os.Stderr, "missing subcommand \"name\"")
		os.Exit(1)
	}

	conn, err := grpc.Dial(":4444", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to the backend:%v", err)
	}
	client := greeter.NewGreeterServiceClient(conn)

	switch cmd := flag.Arg(0); cmd {
	case "name":
		name := strings.Join(flag.Args()[1:], " ")
		err = Greet(context.Background(), client, name)
	default:
		err = fmt.Errorf("unknown subcommand: %s", cmd)
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

//Greet takes a string and returns an error
func Greet(ctx context.Context, client greeter.GreeterServiceClient, name string) error {
	greeting := &greeter.GreetRequest{
		Name: name,
	}
	response, err := client.Greet(ctx, greeting)
	if err != nil {
		return err
	}
	fmt.Println(response.String())
	return nil
}
