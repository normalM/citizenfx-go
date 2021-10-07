package main

import (
	"fmt"
	"syscall/js"

	cfx "github.com/normalM/citizenfx-go"
)

var (
	c  = make(chan bool)
	fn = js.FuncOf
)

func init() {
	cfx.Print(fmt.Sprintf("wasm %s loaded.", cfx.Server.GetCurrentResourceName()))
}
func main() {

	cfx.Server.RegisterCommand("asd", fn(func(this js.Value, inputs []js.Value) interface{} {

		cfx.TriggerClientEvent(
			"chat:addMessage",
			-1,
			"TEST!!",
		)
		cfx.Print("DO IT!!!!!!!")

		return nil
	}), false)

	<-c
}
