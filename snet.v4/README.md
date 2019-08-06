# snet
# 这是一个简易的Tcp框架
 

## snet.v1 适用于十分简陋的任务

### 数据包协议
### ---------------------------------------------------------------------------------
### |    1 byte     |    4 byte     |    4 byte     |    n byte     |    1 byte     |
### ---------------------------------------------------------------------------------
### |    起始符(!)  |   数据长度(n)  | 长度校验(CRC)  |     数据      |   结束符(\n)   |
### ---------------------------------------------------------------------------------
###


```Go

type server struct {
}

func main() {
	s := NewServer("localhost", "496", &server{})
	s.Serve()

}

func (this *server) OnConnect(conn *net.TCPConn) {
	gcmd.Println(gcmd.Info, "客户端连接:", conn.RemoteAddr())
}

func (this *server) OnDisConnect(conn *net.TCPConn, reason string) {
	gcmd.Println(gcmd.Info, "客户端断开连接:", conn.RemoteAddr(), "原因:", reason)
}

func (this *server) OnRecvMessage(conn *net.TCPConn, msg []byte) {
	// gcmd.Println(gcmd.Ok, "接收客户端消息:", msg)
	gcmd.Println(gcmd.Ok, "接收客户端消息:", string(msg))
}

func (this *server) OnSendMessage(conn *net.TCPConn) {

}



