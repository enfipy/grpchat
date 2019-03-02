package chat

import (
	"log"
	"reflect"

	"google.golang.org/grpc"
)

type test struct{}

func ExecMethod(arg string, conn *grpc.ClientConn) {
	log.Printf("Execute %s method", arg)

	t := test{}

	inputs := make([]reflect.Value, 1)
	inputs[0] = reflect.ValueOf(conn)

	reflect.ValueOf(&t).MethodByName(arg).Call(inputs)
}
