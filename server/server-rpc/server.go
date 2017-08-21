package serverrpc

import "net/rpc"

func Register(server *rpc.Server, zig shared.Arith) {
	// Registra a interface Ziguifryda com o nome de "Ziguifryda"
	server.RegisterName("Ziguifryda", zig)
}
