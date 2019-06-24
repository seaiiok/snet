package debugger

import (
	"github.com/seaiiok/utils/fmtx"
)

var DebugOn = false

func Debug(a ...interface{}) {
	if DebugOn == true && a != nil {
		fmtx.Println(fmtx.Info, a...)
	}

}
