# snet
# 这是一个简易的Tcp框架
 

## snet.v1 适用于十分简陋的任务

```Go

func serverGo() {
	s := New()
	s.OnConnect(func(conn *Connection) {
		//建立连接事件
	})

	s.OnDisConnect(func(conn *Connection) {
		//断开连接事件
	})

	s.OnSendMessage(func(conn *Connection) {
		//向客户端发送数据
		// conn.OnSendMsg(snet.Package{})
	})

	s.OnRecvMessage(func(conn *Connection, msg Package) {
		//收到客户端数据
		//fmt.Println(msg)
	})
}
