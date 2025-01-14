package linkedmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoublyLink_InsertAtHead(t *testing.T) {
	link := NewDoublyLinkedList[int]()
	link.InsertAtHead(1)
	link.InsertAtHead(2)
	link.InsertAtHead(3)
	link.InsertAtHead(4)

	assert.Equal(t, link.Values(), []int{4, 3, 2, 1})
	assert.Equal(t, link.Length(), 4)
	assert.Equal(t, link.Head.Data, 4)
	assert.Equal(t, link.Tail.Data, 1)
}

func TestSinglyLink_InsertAtTail(t *testing.T) {
	link := NewDoublyLinkedList[int]()
	link.InsertAtTail(1)
	link.InsertAtTail(2)
	link.InsertAtTail(3)
	link.InsertAtTail(4)

	assert.Equal(t, link.Values(), []int{1, 2, 3, 4})
	assert.Equal(t, link.Length(), 4)
	assert.Equal(t, link.Head.Data, 1)
	assert.Equal(t, link.Tail.Data, 4)
}

func TestDeleteAtHead(t *testing.T) {
	link := NewDoublyLinkedList[int]()
	link.InsertAtTail(1)
	link.InsertAtTail(2)
	link.InsertAtTail(3)

	link.DeleteAtHead()
	assert.Equal(t, link.Values(), []int{2, 3})
	assert.Equal(t, link.Length(), 2)

	link.DeleteAtHead()
	assert.Equal(t, link.Values(), []int{3})
	assert.Equal(t, link.Length(), 1)

	link.DeleteAtHead()
	assert.Equal(t, link.Values(), []int{})
	assert.Equal(t, link.Length(), 0)

	link.DeleteAtHead()
	assert.Equal(t, link.Values(), []int{})
	assert.Equal(t, link.Length(), 0)
}

func TestDeleteAtTail(t *testing.T) {
	link := NewDoublyLinkedList[int]()
	link.InsertAtTail(1)
	link.InsertAtTail(2)
	link.InsertAtTail(3)

	link.DeleteAtTail()
	assert.Equal(t, link.Values(), []int{1, 2})
	assert.Equal(t, link.Length(), 2)

	link.DeleteAtTail()
	assert.Equal(t, link.Values(), []int{1})
	assert.Equal(t, link.Length(), 1)

	link.DeleteAtTail()
	assert.Equal(t, link.Values(), []int{})
	assert.Equal(t, link.Length(), 0)

	link.DeleteAtTail()
	assert.Equal(t, link.Values(), []int{})
	assert.Equal(t, link.Length(), 0)
}

func TestDelete(t *testing.T) {
	link := NewDoublyLinkedList[int]()
	n1 := link.InsertAtTail(1)
	n2 := link.InsertAtTail(2)
	n3 := link.InsertAtTail(3)
	n4 := link.InsertAtTail(4)

	link.Delete(n1)
	assert.Equal(t, link.Values(), []int{2, 3, 4})
	assert.Equal(t, link.Length(), 3)

	link.Delete(n4)
	assert.Equal(t, link.Values(), []int{2, 3})
	assert.Equal(t, link.Length(), 2)

	link.Delete(n2)
	assert.Equal(t, link.Values(), []int{3})
	assert.Equal(t, link.Length(), 1)

	link.Delete(n3)
	assert.Equal(t, link.Values(), []int{})
	assert.Equal(t, link.Length(), 0)
}

func TestClear(t *testing.T) {
	link := NewDoublyLinkedList[int]()
	link.InsertAtTail(1)
	link.InsertAtTail(2)
	link.InsertAtTail(3)

	link.Clear()
	assert.Equal(t, link.Values(), []int{})
	assert.Equal(t, link.Length(), 0)
}
