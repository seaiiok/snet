# snet
# 这是一个简易的Tcp框架
 

## snet.v1 适用于十分简陋的任务

```Go

func serverGo() {
	s := New("localhost","496")
	s.OnConnect(func(conn *snet.Connection) {
		//建立连接事件
	})

	s.OnDisConnect(func(conn *snet.Connection) {
		//断开连接事件
	})

	s.OnSendMessage(func(conn *snet.Connection) {
		//向客户端发送数据
		// conn.OnSendMsg([]byte{"msg..."})
	})

	s.OnRecvMessage(func(conn *snet.Connection, msg []byte) {
		//收到客户端数据
		//fmt.Println(string(msg))
	})
}
