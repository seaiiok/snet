package main

import (
	"bufio"
	"gcom/gcmd"
	"net"
	"os"

	"snet/snet.v3"
)

func main() {
	CMDClient()
}

func CMDClient() {
	conn, err := net.Dial("tcp", "localhost:496")
	if err != nil {
		gcmd.Println(gcmd.Warn, "client dial err, exit!")
		return
	}

	p := &snet.Package{}
	//避免控制台显示乱码，临时采用UTF-8
	gcmd.ExecCommand("chcp", "65001")

	gcmd.Println(13, "open client command...")
	for {
		gcmd.Println(11, "please input command...")
		input := bufio.NewScanner(os.Stdin)
		input.Scan()

		p.Msg = []byte(input.Text())
		b := p.Pack()
		_, err := conn.Write(b)
		if err != nil {
			gcmd.Println(gcmd.Err, err)
		}
	}
}
