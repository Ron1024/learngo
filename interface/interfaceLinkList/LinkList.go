package main

type LinkNode struct {
	data interface{}
	next *LinkNode
}

type Link struct {
	head *LinkNode
	tail *LinkNode
}

func (p *Link) InsertHead(data interface{}) {
	node := &LinkNode{
		data: data,
		next: nil,
	}
	if p.head == nil && p.tail == nil {
		p.tail = p.head
		p.head = node
		return
	}
	node.next = p.head
	p.head = node
}

func (p *Link) InsertTail(data interface{}) {
	node := &LinkNode{
		data: data,
		next: nil,
	}

	if p.head == nil && p.tail == nil {
		p.tail = p.head
		p.head = node
		return
	}

	p.tail.next = node
	p.tail = node
}

func ()  {
	
}

func main() {
	
}