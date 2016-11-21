package highlight

import (
	"container/heap"
	"reflect"
	"testing"
)

func TestPOIHeap(t *testing.T) {
	pois := []poi{
		{i: 100},
		{i: 15},
		{i: 15, start: true},
		{i: 10},
		{i: 5},
	}
	hp := &poiHeap{}
	for _, p := range pois {
		heap.Push(hp, p)
	}
	for i, p := range pois {
		h := hp.Peek()
		if !reflect.DeepEqual(p, h) {
			t.Errorf("%d. Peek() expected %+v = %+v", i, p, h)
		}

		h = heap.Pop(hp).(poi)
		if !reflect.DeepEqual(p, h) {
			t.Errorf("%d. Pop() expected %+v = %+v", i, p, h)
		}
	}
}
