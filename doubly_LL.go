//Create a doubly linked list with the help of Pointers and structure
package main

import "fmt"

type Node struct {
    data int
    prev *Node
    next *Node
}

type Doubly_LL struct {
    head *Node
    tail *Node
}

func (list *Doubly_LL) AddToFront(data int) {
    newNode := &Node{data: data, next: list.head}
    if list.head != nil {
        list.head.prev = newNode
    }
    list.head = newNode
    if list.tail == nil {
        list.tail = newNode
    }
}

func (list *Doubly_LL) AddToBack(data int) {
    newNode := &Node{data: data, prev: list.tail}
    if list.tail != nil {
        list.tail.next = newNode
    }
    list.tail = newNode
    if list.head == nil {
        list.head = newNode
    }
}

func (list *Doubly_LL) Display() {
    current := list.head
    for current != nil {
        fmt.Printf("%d <---> ", current.data)
        current = current.next
    }
    fmt.Println("nil")
}

func main() {
    dll := Doubly_LL{}

    dll.AddToFront(3)
    dll.AddToFront(2)
    dll.AddToFront(1)

    dll.AddToBack(4)
    dll.AddToBack(5)

    fmt.Println("Doubly Linked List:")
    dll.Display()
}
