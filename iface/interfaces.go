package iface

type ISnet interface {
	GetConfig() map[string]string
	NewServer() IServer
}

type IServer interface {
	ServerStart()
	ServerStop()
}

// type IClient interface {
// 	ClientStart()
// 	ClientStop()
// }

// type IConnections interface {
// 	AddClient(*net.TCPConn) error
// 	DeleteClient(string) error
// }
