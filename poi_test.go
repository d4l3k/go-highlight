package highlight

import (
	"container/heap"
	"reflect"
	"testing"
)

func TestPOIHeap(t *testing.T) {
	pois := []poi{
		{i: 100},
		{i: 15, start: true},
		{i: 15},
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

func TestPOIHeapRemove(t *testing.T) {
	h := &highlight{}
	pois := []poi{
		{i: 100},
		{i: 10, highlight: h},
		{i: 5},
	}
	hp := &poiHeap{}
	for _, p := range pois {
		heap.Push(hp, p)
	}
	if eq := reflect.DeepEqual([]poi(*hp), pois); !eq {
		t.Fatalf("expected heap %+v to equal %+v", *hp, pois)
	}

	hp.Remove(poi{highlight: h})

	poisWant := []poi{
		{i: 100},
		{i: 5},
	}
	if eq := reflect.DeepEqual([]poi(*hp), poisWant); !eq {
		t.Fatalf("expected heap %+v to equal %+v", *hp, poisWant)
	}
}
