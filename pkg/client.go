package pkg

// import (
// 	"context"
// 	"net"
// )

// type client struct {
// 	clientName    string
// 	clientAddr    string
// 	conn          *net.TCPConn
// 	clientStatus  int
// 	clientIsClose chan bool
// 	contx         context.Context
// 	contxCancel   context.CancelFunc
// }

// func newClient(ct context.Context, conn *net.TCPConn) *client {
// 	ctx, cancel := context.WithCancel(ct)
// 	clientName := conn.RemoteAddr().String()
// 	return &client{
// 		clientName:    clientName,
// 		clientAddr:    clientName,
// 		conn:          conn,
// 		clientIsClose: make(chan bool, 0),
// 		contx:         ctx,
// 		contxCancel:   cancel,
// 	}
// }

// func (c *client) StartWork() {

// }

// func (c *client) StopWork() {
// 	c.ClientContext.ContxCancel()
// }

// func (c *client) IsClose() chan bool {
// 	return c.ClientIsClose
// }

// func (c *client) CloseClient() {

// }

// type Client struct {
// 	Conf map[string]string
// }

// func (s *Client) ClientStart() {

// }

// func (s *Client) ClientStop() {

// }
