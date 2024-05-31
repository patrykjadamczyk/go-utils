package base

import (
	"sync"
	"testing"

	. "github.com/patrykjadamczyk/go-utils/base"
)

func TestBroadcastChannel(t *testing.T) {
	bc := MakeBroadcastChannel[int]()

	var wgReturn sync.WaitGroup
	var wgGoStart sync.WaitGroup
	wgReturn.Add(2)
	wgGoStart.Add(1)
	go func(b *BroadcastChannel[int]) {
		t.Log("GR1: Start")
		t.Log("GR1: Subscribed to test channel")
		bcg := bc.Subscribe("test")
		wgGoStart.Done()
		m1 := <-bcg
		if m1 != 12 {
			t.Error("First message on broadcast channel should be 12")
		}
		t.Log("GR1: Processed first message")
		m2 := <-bcg
		if m2 != 21 {
			t.Error("Second message on broadcast channel should be 21")
		}
		t.Log("GR1: Processed second message")
		b.Unsubscribe("test")
		t.Log("GR1: Unsubscribed from test channel")
		b.Broadcast(23)
		t.Log("GR1: Broadcasted message")
		t.Log("GR1: Returning")
		wgReturn.Done()
	}(&bc)
	go func(b *BroadcastChannel[int]) {
		t.Log("GR2: Start")
		wgGoStart.Wait()
        b.Broadcast(12)
		t.Log("GR2: Broadcasted message 12")
        b.Broadcast(21)
		t.Log("GR2: Broadcasted message 21")
		t.Log("GR2: Returning")
		wgReturn.Done()
	}(&bc)

	wgReturn.Wait()
}
