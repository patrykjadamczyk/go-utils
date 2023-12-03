package base

// Broadcast Channel Structure
type BroadcastChannel[T any] struct {
	subchannels map[string]chan T
}

// Subscribe to Broadcast Channel by getting your own subchannel
func (bc BroadcastChannel[T]) Subscribe(id string) chan T {
	if bc.subchannels == nil {
		bc.subchannels = make(map[string]chan T)
	}
	bc.subchannels[id] = make(chan T)
	return bc.subchannels[id]
}

// Unsubscribe from Broadcast Channel by deleting the subchannel
func (bc BroadcastChannel[T]) Unsubscribe(id string) {
	delete(bc.subchannels, id)
}

// Broadcast Message to all subchannels
func (bc BroadcastChannel[T]) Broadcast(message T) {
	for _, subchannel := range bc.subchannels {
		subchannel <- message
	}
}

// Make Broadcast Channel
func MakeBroadcastChannel[T any]() BroadcastChannel[T] {
    return BroadcastChannel[T]{}
}
