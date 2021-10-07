package citizenfx

import "syscall/js"

type _server struct {
	GetCurrentResourceNamex func() string `js:"GetCurrentResourceName()"`
}
type _client struct {
}

var (
	Global = js.Global()
	Server = new(_xxx)
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

func Emit(eventName string, args ...interface{}) {
	Global.Call("Emit", eventName, args)

}

func TriggerEvent(eventName string, args ...interface{}) {
	Global.Call("TriggerEvent", eventName, args)

}

// func EmitNet(eventName string, args ...interface{}) {
// }
func EmitNet(eventName string, target interface{}, args ...interface{}) {
	if target == nil {
		Global.Call("EmitNet", eventName, target, args)
	} else {
		Global.Call("EmitNet", eventName, args)
	}

}
func TriggerServerEvent(eventName string, args ...interface{}) {
	Global.Call("TriggerServerEvent", eventName, args)

}
func TriggerLatentServerEvent(eventName string, bps float64, args ...interface{}) {
	Global.Call("TriggerLatentServerEvent", eventName, bps, args)

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

func TriggerClientEvent(eventName string, target interface{}, args ...interface{}) {
	Global.Call("TriggerClientEvent", eventName, target, args)
}
func TriggerLatentClientEvent(eventName string, target interface{}, bps float64, args ...interface{}) {
	Global.Call("TriggerLatentClientEvent", eventName, target, bps, args)
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
