package dealing

import (
	"container/ring"
	"fmt"
)

type CircularList struct {
	list *ring.Ring
}

func NewCircularList(items ...interface{}) *CircularList {
	cl := &CircularList{
		list: ring.New(len(items)),
	}
	for i := 0; i < cl.list.Len(); i++ {
		cl.list.Value = items[i]
		cl.list = cl.list.Next()
	}
	return cl
}

func (cl *CircularList) ShowAll() {
	cl.list.Do(func(x interface{}) {
		fmt.Printf("Item: %v\n", x)
	})
}

func (cl *CircularList) GetItem() interface{} {
	val := cl.list.Value
	cl.list = cl.list.Next()
	return val
}

func main() {

	cl := NewCircularList("win", "loss", "tie")

	for i := 0; i < 5; i++ {
		fmt.Printf("Iteration #%d is: %v\n", i, cl.GetItem())
	}
	fmt.Println("----------")

	cl2 := NewCircularList(0, 1, 2, 3, 5, 8)
	for i := 0; i < 10; i++ {
		fmt.Printf("Iteration #%d is: %v\n", i, cl2.GetItem())
	}
	fmt.Println("----------")

	cl2.ShowAll()

}
