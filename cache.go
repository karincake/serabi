package serabi

import (
	"reflect"
)

// Status of using cache or not
var CacheEnabled = false

// Count of classes allowed to be cached
var CacheMaxCount = 5

type registeredClass struct {
	name       string
	inputVNFC  int
	fieldT     []reflect.StructField
	tag        []string
	key        []string
	typeString []string
	parsedTag  [][]keyVal
	tagNames   []string
	next       *registeredClass
}

type registeredClassList struct {
	head *registeredClass
}

var cache registeredClassList = registeredClassList{}

func (r *registeredClassList) push(n string, rc registeredClass) {
	newNode := &rc
	newNode.name = n

	if r.head == nil {
		r.head = newNode
		return
	}

	curr := r.head
	for curr.next != nil {
		curr = curr.next
	}

	curr.next = newNode
}

func (r *registeredClassList) shift() {
	if r.head == nil {
		return
	}

	if r.head.next != nil {
		r.head = r.head.next
	} else {
		r.head = nil
	}

}

func (r *registeredClassList) classExists(n string) bool {
	// fmt.Println("searching for class " + n)
	if r.head == nil {
		// fmt.Println("class not found")
		return false
	}

	current := r.head
	for current != nil {
		// fmt.Println(current.name)
		if current.name == n {
			return true
		}
		current = current.next
	}
	// fmt.Println("class not found")
	return false
}

func (r *registeredClassList) get(n string) *registeredClass {
	if r.head == nil {
		return nil
	}

	current := r.head
	for current != nil {
		if current.name == n {
			return current
		}
		current = current.next
	}
	return nil
}
