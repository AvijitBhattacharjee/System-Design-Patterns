package messagequeue

import (
	"fmt"
	"hash/crc32"
	"sync"
)

//////////////////////////////////////////////////////
// Message
//////////////////////////////////////////////////////

type Message struct {
	ID      string
	Key     string
	Payload string
	Offset  int64
}

//////////////////////////////////////////////////////
// Partition
//////////////////////////////////////////////////////

type Partition struct {
	ID       int
	Messages []Message
	Offset   int64
	mu       sync.RWMutex
}

func NewPartition(id int) *Partition {
	return &Partition{
		ID:       id,
		Messages: []Message{},
	}
}

func (p *Partition) Append(msg Message) int64 {
	p.mu.Lock()
	defer p.mu.Unlock()

	msg.Offset = p.Offset
	p.Offset++

	p.Messages = append(p.Messages, msg)

	return msg.Offset
}

func (p *Partition) Read(offset int64) []Message {
	p.mu.RLock()
	defer p.mu.RUnlock()

	if offset >= int64(len(p.Messages)) {
		return nil
	}

	result := make([]Message, len(p.Messages[offset:]))
	copy(result, p.Messages[offset:])

	return result
}

//////////////////////////////////////////////////////
// Partitioner
//////////////////////////////////////////////////////

type Partitioner interface {
	GetPartition(key string, totalPartitions int) int
}

type HashPartitioner struct{}

func (h HashPartitioner) GetPartition(
	key string,
	totalPartitions int,
) int {

	hash := crc32.ChecksumIEEE([]byte(key))

	return int(hash) % totalPartitions
}

//////////////////////////////////////////////////////
// Topic
//////////////////////////////////////////////////////

type Topic struct {
	Name        string
	Partitions  []*Partition
	Partitioner Partitioner
}

func NewTopic(name string, partitionCount int) *Topic {

	t := &Topic{
		Name:        name,
		Partitioner: HashPartitioner{},
	}

	for i := 0; i < partitionCount; i++ {
		t.Partitions = append(
			t.Partitions,
			NewPartition(i),
		)
	}

	return t
}

func (t *Topic) Publish(
	key string,
	payload string,
) {

	pid := t.Partitioner.GetPartition(
		key,
		len(t.Partitions),
	)

	msg := Message{
		ID:      fmt.Sprintf("%s-%s", key, payload),
		Key:     key,
		Payload: payload,
	}

	offset := t.Partitions[pid].Append(msg)

	fmt.Printf(
		"[Producer] Key=%s stored in Partition=%d Offset=%d\n",
		key,
		pid,
		offset,
	)
}

//////////////////////////////////////////////////////
// Offset Manager
//////////////////////////////////////////////////////

type OffsetManager struct {
	offsets map[string]int64
	mu      sync.RWMutex
}

func NewOffsetManager() *OffsetManager {
	return &OffsetManager{
		offsets: make(map[string]int64),
	}
}

func (o *OffsetManager) Get(consumer string) int64 {

	o.mu.RLock()
	defer o.mu.RUnlock()

	return o.offsets[consumer]
}

func (o *OffsetManager) Commit(
	consumer string,
	offset int64,
) {

	o.mu.Lock()
	defer o.mu.Unlock()

	o.offsets[consumer] = offset
}

//////////////////////////////////////////////////////
// Consumer
//////////////////////////////////////////////////////

type Consumer struct {
	ID      string
	Topic   *Topic
	Offsets *OffsetManager
}

func NewConsumer(
	id string,
	topic *Topic,
	offsets *OffsetManager,
) *Consumer {

	return &Consumer{
		ID:      id,
		Topic:   topic,
		Offsets: offsets,
	}
}

func (c *Consumer) Poll(partitionID int) {

	offset := c.Offsets.Get(
		fmt.Sprintf(
			"%s-%d",
			c.ID,
			partitionID,
		),
	)

	msgs := c.Topic.
		Partitions[partitionID].
		Read(offset)

	if len(msgs) == 0 {
		fmt.Printf(
			"[Consumer %s] No Messages\n",
			c.ID,
		)
		return
	}

	for _, msg := range msgs {

		fmt.Printf(
			"[Consumer %s] Partition=%d Offset=%d Payload=%s\n",
			c.ID,
			partitionID,
			msg.Offset,
			msg.Payload,
		)

		c.Offsets.Commit(
			fmt.Sprintf(
				"%s-%d",
				c.ID,
				partitionID,
			),
			msg.Offset+1,
		)
	}
}

//////////////////////////////////////////////////////
// Broker
//////////////////////////////////////////////////////

type Broker struct {
	Topics map[string]*Topic
	mu     sync.RWMutex
}

func NewBroker() *Broker {
	return &Broker{
		Topics: make(map[string]*Topic),
	}
}

func (b *Broker) CreateTopic(
	name string,
	partitions int,
) {

	b.mu.Lock()
	defer b.mu.Unlock()

	b.Topics[name] = NewTopic(
		name,
		partitions,
	)

	fmt.Printf(
		"[Broker] Topic %s created with %d partitions\n",
		name,
		partitions,
	)
}

func (b *Broker) Publish(
	topic string,
	key string,
	payload string,
) {

	b.mu.RLock()
	defer b.mu.RUnlock()

	t, ok := b.Topics[topic]
	if !ok {
		fmt.Println("topic not found")
		return
	}

	t.Publish(
		key,
		payload,
	)
}

//////////////////////////////////////////////////////
// Main
//////////////////////////////////////////////////////

func main() {

	broker := NewBroker()

	broker.CreateTopic(
		"orders",
		4,
	)

	//////////////////////////////////////////////////
	// Producer
	//////////////////////////////////////////////////

	broker.Publish(
		"orders",
		"order-123",
		"created",
	)

	broker.Publish(
		"orders",
		"order-123",
		"payment-complete",
	)

	broker.Publish(
		"orders",
		"order-123",
		"shipped",
	)

	broker.Publish(
		"orders",
		"order-456",
		"created",
	)

	broker.Publish(
		"orders",
		"order-456",
		"payment-complete",
	)

	//////////////////////////////////////////////////
	// Consumer
	//////////////////////////////////////////////////

	offsetManager := NewOffsetManager()

	consumer1 := NewConsumer(
		"C1",
		broker.Topics["orders"],
		offsetManager,
	)

	fmt.Println("\n--- Reading Partition 0 ---")
	consumer1.Poll(0)

	fmt.Println("\n--- Reading Partition 1 ---")
	consumer1.Poll(1)

	fmt.Println("\n--- Reading Partition 2 ---")
	consumer1.Poll(2)

	fmt.Println("\n--- Reading Partition 3 ---")
	consumer1.Poll(3)
}