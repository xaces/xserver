package nats

import (
	"github.com/wlgd/xutils/mq"
)

type nats struct {
	Conn    *mq.NatsClient
	chanmap map[string]chan []byte
}

var Default = &nats{}

func (n *nats) MsgPub(msg []byte) {
	if msg == nil {
		return
	}
	for _, v := range n.chanmap {
		v <- msg
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
	n.chanmap = make(map[string]chan []byte)
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

func (n *nats) PushClient(key string, ch chan []byte) {
	if n.Conn != nil {
		n.chanmap[key] = ch
	}
}

func (n *nats) PullClient(key string) {
	delete(n.chanmap, key)
}

func (n *nats) Stop() {
	if n.Conn == nil {
		return
	}
	n.Conn.Release()
}
