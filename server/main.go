package main

import (
	"context"
	"fmt"
	"math/rand"
	"net"
	"time"

	greeter "github.com/rugwirobaker/tutorial/greeter/proto"
	"google.golang.org/grpc"
)

//Greeter implements GreeterServiceServer interface
type Greeter struct {
	Greetings []string
}

//NewGreeter takes in a number of greetings and
//returns an new Greeter Service Server
func NewGreeter(greetings ...string) *Greeter {
	return &Greeter{
		Greetings: greetings,
	}
}

//Greet returns a greeting response or an error given a greeting request.
//The function always responds with a random response among those in Greeter.Greetings.
func (g *Greeter) Greet(ctx context.Context, req *greeter.GreetRequest) (*greeter.GreetResponse, error) {
	if len(req.Name) < 1 {
		return nil, fmt.Errorf("the name is short")
	}

	if len(g.Greetings) < 1 {
		return nil, fmt.Errorf("We have not greetings in store")
	}

	//throw dice
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	pick := g.Greetings[r1.Intn(len(g.Greetings))]
	response := &greeter.GreetResponse{
		Response: fmt.Sprintf("%s, %s", pick, req.GetName()),
	}
	return response, nil
}

func main() {
	grpcServer := grpc.NewServer()

	g := NewGreeter("Hello", "Greetings", "Saluto", "Mwiriwe", "Greetings Infidel")
	greeter.RegisterGreeterServiceServer(grpcServer, g)
	lis, err := net.Listen("tcp", ":4444")
	if err != nil {
		panic(err)
	}
	fmt.Println("Listening on port :4444")
	grpcServer.Serve(lis)
}
