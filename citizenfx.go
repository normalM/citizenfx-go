package citizenfx

import "syscall/js"

type _client struct {
}

var (
	Global = js.Global()
	Server = new(_server)
	Client = new(_client)
)

func init() {
	if err := BindGlobals(Server); err != nil {
		panic(err)
	}
}
func AddRawEventListener(eventName string, callback js.Func) {

	Global.Call("AddRawEventListener", eventName, callback)
}

func AddEventListener(eventName string, callback js.Func, netSafe bool) {
	Global.Call("AddEventListener", eventName, callback, netSafe)
}
func On(eventName string, callback js.Func) {
	Global.Call("On", eventName, callback)

}
func AddEventHandler(eventName string, callback js.Func) {
	Global.Call("AddEventHandler", eventName, callback)

}

func AddNetEventListener(eventName string, callback js.Func) {
	Global.Call("AddNetEventListener", eventName, callback)

}
func OnNet(eventName string, callback js.Func) {
	Global.Call("OnNet", eventName, callback)

}

func Emit(args ...interface{}) {
	Global.Call("Emit", args)

}

func TriggerEvent(args ...interface{}) {
	Global.Call("TriggerEvent", args)

}

// func EmitNet(args ...interface{}) {
// }
func EmitNet(args ...interface{}) {
	Global.Call("EmitNet", args...)

}
func TriggerServerEvent(args ...interface{}) {
	Global.Call("TriggerServerEvent", args)

}
func TriggerLatentServerEvent(args ...interface{}) {
	Global.Call("TriggerLatentServerEvent", args...)

}

func GetPlayerIdentifiers(player interface{}) js.Value {
	return Global.Call("GetPlayerIdentifiers", player)

}
func GetPlayerTokens(player interface{}) js.Value {
	return Global.Call("GetPlayerTokens", player)
}
func GetPlayers() js.Value {
	return Global.Call("GetPlayers")
}

func SendNUIMessage(data interface{}) {
	Global.Call("SendNUIMessage", data)
}

func TriggerClientEvent(args ...interface{}) {
	Global.Call("TriggerClientEvent", args...)
}
func TriggerLatentClientEvent(args ...interface{}) {
	Global.Call("TriggerLatentClientEvent", args...)
}

func RemoveEventListener(eventName string, callback js.Func) {
	Global.Call("RemoveEventListener", eventName, callback)
}

func SetTick(callback js.Func) float64 {
	return Global.Call("SetTick", callback).Float()

}
func ClearTick(callback float64) {
	Global.Call("ClearTick", callback)
}

func NewStateBag(name string) js.Value {
	return Global.Call("NewStateBag", name)
}
func Entity(entity float64) js.Value {
	return Global.Call("Entity", entity)
}
func Player(entity interface{}) js.Value {
	return Global.Call("Player", entity)
}
func Print(arg ...interface{}) {
	Global.Get("console").Call("log", arg...)
}
