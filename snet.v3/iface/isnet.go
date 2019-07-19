package iface

type ISnet interface {
	Serve()
	Stop()
	AddConnection(IConnection)
}
