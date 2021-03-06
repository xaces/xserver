package snowflake

import (
	"errors"
	"sync"
	"time"
)

const (
	workerBits  uint8 = 10
	numberBits  uint8 = 12
	workerMax   int64 = -1 ^ (-1 << workerBits)
	numberMax   int64 = -1 ^ (-1 << numberBits)
	timeShift   uint8 = workerBits + numberBits
	workerShift uint8 = numberBits
	startTime   int64 = 1525705533000 // 如果在程序跑了一段时间修改了epoch这个值 可能会导致生成相同的ID
)

// Snowflake struct
type Snowflake struct {
	mu        sync.Mutex
	timestamp int64
	workerID  int64
	number    int64
}

// New 创建实例
func New(workerID int64) (*Snowflake, error) {
	if workerID < 0 || workerID > workerMax {
		return nil, errors.New("worker id excess of quantity")
	}
	// 生成一个新节点
	return &Snowflake{
		timestamp: 0,
		workerID:  workerID,
		number:    0,
	}, nil
}

// NextId 下一个ID
func (w *Snowflake) NextId() uint64 {
	w.mu.Lock()
	defer w.mu.Unlock()
	now := time.Now().UnixNano() / 1e6
	if w.timestamp == now {
		w.number++
		if w.number > numberMax {
			for now <= w.timestamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		w.number = 0
		w.timestamp = now
	}
	Id := uint64((now-startTime)<<timeShift | (w.workerID << workerShift) | (w.number))
	return Id
}

var (
	_snowflake *Snowflake = nil
)

func init() {
	_snowflake, _ = New(0xf)
}

func PrimaryKey() uint64 {
	return _snowflake.NextId()
}
