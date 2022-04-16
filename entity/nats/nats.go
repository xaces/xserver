package nats

import (
	"github.com/wlgd/xutils/mq"
)

type client struct {
	msg    chan []byte
	active bool
}

type nats struct {
	Conn *mq.NatsClient
	clis map[string]*client
}

var Default = &nats{}

func (n *nats) MsgPub(msg []byte) {
	if msg == nil {
		return
	}
	for k, v := range n.clis {
		if v.active {
			v.msg <- msg
			continue
		}
		close(v.msg)
		delete(n.clis, k)
	}
}

func (n *nats) Run(addr string) error {
	if addr == "" {
		addr = mq.DefaultURL
	}
	conn, err := mq.NewNatsClient(addr, true)
	if err != nil {
		return err
	}
	n.Conn = conn
	n.clis = make(map[string]*client)
	n.Conn.Subscribe("device.online", func(b []byte) {
		n.MsgPub(b)
	})
	n.Conn.Subscribe("device.status", func(b []byte) {
		n.MsgPub(b)
	})
	n.Conn.Subscribe("device.alarm", func(b []byte) {
		n.MsgPub(b)
	})
	n.Conn.Subscribe("device.event", func(b []byte) {
		n.MsgPub(b)
	})
	return nil
}

func (n *nats) Push(key string) chan []byte {
	if n.Conn == nil {
		return nil
	}
	c := &client{
		active: true,
		msg:    make(chan []byte, 10),
	}
	n.clis[key] = c
	return c.msg
}

func (n *nats) Pop(key string) {
	if n.Conn == nil {
		return
	}
	if v, ok := n.clis[key]; ok {
		v.active = false
	}
}

func (n *nats) Stop() {
	if n.Conn == nil {
		return
	}
	n.Conn.Release()
}
