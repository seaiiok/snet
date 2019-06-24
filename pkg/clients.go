package pkg

// import (
// 	"snet/iface"
// 	"sync"
// )

// import (
// 	"context"
// 	"errors"
// 	"fmt"
// 	"net"
// 	"snet/iface"
// 	"time"
// )

// type clients struct {
// 	counts    int
// 	clientMap map[string]iface.IClient
// 	sync.RWMutex
// }

// func NewClients(contx iface.IContext) iface.IClients {
// 	ctx, cancel := context.WithCancel(contx.GetContext())
// 	return &Clients{
// 		Counts:  0,
// 		Clients: make(map[string]iface.IClient, 0),
// 		ClientsContext: &Context{
// 			Contx:       ctx,
// 			ContxCancel: cancel,
// 		},
// 	}
// }

// func (clients *Clients) AddClient(conn *net.TCPConn) error {
// 	newClient := NewClient(clients.ClientsContext, conn)
// 	clientName := conn.RemoteAddr().String()
// 	if _, found := clients.Clients[clientName]; found != true {
// 		clients.Clients[clientName] = newClient
// 		return nil
// 	}
// 	return errors.New("this client is exist,add client failed")
// }

// func (clients *Clients) DeleteClient(clientName string) error {
// 	if _, found := clients.Clients[clientName]; found == true {
// 		delete(clients.Clients, clientName)
// 		return nil
// 	}
// 	return errors.New("this client is not exist,delete client failed")
// }

// func (c *Clients) Start() {
// 	for {
// 		for k, v := range c.Clients {
// 			fmt.Println("name:", k, "client:", v)

// 			select {
// 			case <-v.IsClose():
// 				c.DeleteClient(k)
// 				v.StopWork()
// 				return
// 			default:

// 			}
// 		}
// 		time.Sleep(2 * time.Second)
// 	}

// }

// // func (client *Clients) GetCounts() int {
// // 	return len(client.Clients)
// // }

// // func (c *Clients) ClientsHandle(ctx context.Context, connection *net.TCPConn) {
// // 	conn := NewConnection(ctx,connection)
// // 	err := c.addClient(conn)
// // 	if err != nil {
// // 		fmt.Println("Add Client failed:", err)
// // 	}

// // 	go conn.Recv()

// // 	go func() {
// // 		for {
// // 			select {
// // 			case closeConnChan := <-conn.CloseConn():
// // 				c.deleteClient(closeConnChan)
// // 				fmt.Println("Client Delete succ:", closeConnChan)
// // 				return
// // 			default:
// // 				time.Sleep(1 * time.Second)
// // 				fmt.Println("Client Delete Emtpy")
// // 			}
// // 		}
// // 	}()
// // }

