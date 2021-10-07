package citizenfx

import "syscall/js"

type _server struct {
}

var (
	Global = js.Global()
	Server = new(_server)
)
