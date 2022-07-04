package subject

import (
	"github.com/xaces/xutils/mq"
)

type client struct {
	msg    chan []byte
	active bool
}

type nats struct {
	Conn *mq.Client
	clis map[string]*client
}

var Default = &nats{}

func (o *nats) distrbuteMsg(b []byte) {
	if b == nil {
		return
	}
	for k, v := range o.clis {
		if v.active {
			v.msg <- b
			continue
		}
		close(v.msg)
		delete(o.clis, k)
	}
}

func (o *nats) Run(addr string) error {
	// 用四个goroutine连接服务，订阅
	c, err := mq.New(&mq.Options{Address: addr, Goc: 4}, mq.NewNats)
	if err != nil {
		c.Shutdown()
		return err
	}
	o.clis = make(map[string]*client)
	c.Subscribe("device.online", 1, o.distrbuteMsg)
	c.Subscribe("device.status", 1, o.distrbuteMsg)
	c.Subscribe("device.alarm", 1, o.distrbuteMsg)
	c.Subscribe("device.event", 1, o.distrbuteMsg)
	return nil
}

func (o *nats) Shutdown() {
	if o.Conn == nil {
		return
	}
	o.Conn.Shutdown()
}

func (o *nats) NewClient(key string) chan []byte {
	if o.Conn == nil {
		return nil
	}
	c := &client{
		active: true,
		msg:    make(chan []byte, 10),
	}
	o.clis[key] = c
	return c.msg
}

func (o *nats) DelClient(key string) {
	if o.Conn == nil {
		return
	}
	if v, ok := o.clis[key]; ok {
		v.active = false
	}
}
