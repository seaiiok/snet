package pkg

// import (
// 	"context"
// 	"fmt"
// 	"net"
// 	"snet/iface"
// 	"time"
// )

// type Connection struct {
// 	ClientName string
// 	Online     int64
// 	Conn       *net.TCPConn
// 	*Context
// }

// func NewConnection(cont iface.IContext, conn *net.TCPConn) iface.IConnection {
// 	subctx, subcancel := context.WithCancel(cont.GetContext())
// 	return &Connection{
// 		ClientName: conn.RemoteAddr().String(),
// 		Online:     time.Now().Unix(),
// 		Conn:       conn,
// 		Contx: &Contx{
// 			Ctx:           subctx,
// 			ContextCancel: subcancel,
// 		},
// 	}
// }

// func (c *Connection) Send(buf []byte) {
// 	// c.WriteBuf(buf)
// 	// for {
// 	// 	// fmt.Println("Server Send succ")
// 	// 	select {
// 	// 	case <-c.Ctx.Done():
// 	// 		return
// 	// 	default:
// 	// 	}
// 	// }
// }

// func (c *Connection) Recv() {
// 	for {
// 		buf := make([]byte, 512)
// 		cnt, err := c.Conn.Read(buf)
// 		if err != nil {
// 			c.Stop()
// 			fmt.Println("Server Read Err:", err)
// 			return
// 		}
// 		fmt.Println("Server Read Data:", string(buf[:cnt]))
// 		c.WriteBuf(buf[:cnt])

// 		select {
// 		case <-c.Ctx.Done():
// 			return
// 		default:
// 		}
// 	}
// }

// func (c *Connection) GetClientName() string {
// 	return c.ClientName
// }

// func (c *Connection) Stop() {
// 	c.ContextCancel()
// }

// func (c *Connection) GetClientOnline() int64 {
// 	return c.Online
// }

// func (c *Connection) Start() {

// 	go func() {
// 		for {
// 			time.Sleep(5 * time.Second)
// 			c.Stop()
// 		}
// 	}()

// 	go c.Recv()

// 	for {
// 		select {
// 		case <-c.Ctx.Done():
// 			return
// 		default:
// 		}
// 	}

// }

// func (c *Connection) WriteBuf(buf []byte) {

// 	_, err := c.Conn.Write(buf)
// 	if err != nil {
// 		c.Stop()
// 		fmt.Println("Server Write Err:", err)
// 		return
// 	}

// }


