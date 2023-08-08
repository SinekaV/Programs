//Implement a queue data structure with the help of structures and pointers. 
package main

import "fmt"

type Queue struct {
    front,rear,size int
    capacity int
    array []int
}

func NewQueue(capacity int) *Queue {
    return &Queue{
        front:0,
        rear:-1,
        size:0,
        capacity:capacity,
        array:make([]int, capacity),
    }
}

func (q *Queue) Enqueue(item int) bool {
    if q.size == q.capacity {//queue full
        return false 
    }
    q.rear = (q.rear + 1) % q.capacity
    q.array[q.rear] = item//arr[i]=val;
    q.size++
    return true
}

func (q *Queue) Dequeue() (int, bool) {
    if q.size == 0 {
        return 0, false
    }
    item := q.array[q.front]
    q.front = (q.front + 1) % q.capacity
    q.size--
    return item, true
}

func main() {
    q := NewQueue(5)

    q.Enqueue(1)
    q.Enqueue(2)
    q.Enqueue(3)
    q.Enqueue(4)
    
    fmt.Println(q.Enqueue(5)) 
    fmt.Println(q.Enqueue(5)) 
    fmt.Println(q.Dequeue()) 
    fmt.Println(q.Dequeue()) 
}