package snet

import (
	"utils/snet/iface"
	"utils/snet/pkg"

	"github.com/seaiiok/utils/files"
)

//NewSnet ...path:configfile
func NewSnet(path string) iface.ISnet {
	if files.IsExists(path) {
		return nil
	}

	return &pkg.snet{
		Conf: loadConfigFile(confPath),
	}
}
