package snet

import (
	"fmt"
	"snet/iface"
	"snet/pkg"
)

//NewServer ...path:configfile
func NewServer(path string) iface.IServer {
	conf := pkg.LoadConfigFile(path)
	fmt.Println(conf)
	return &pkg.Server{
		Conf: pkg.LoadConfigFile(path),
	}

}
