package gremlin

import (
	"github.com/qasaur/gremgo"
	"fmt"
)

func Prueba (){

	dialer := gremgo.NewDialer("127.0.0.1:8182") // Returns a WebSocket dialer to connect to Gremlin Server
	g, err := gremgo.Dial(dialer) // Returns a gremgo client to interact with
	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := g.Execute( // Sends a query to Gremlin Server with bindings
		"g.V().values('name')",
		map[string]string{},
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}