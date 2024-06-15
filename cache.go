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

func (obj *registeredClassList) push(n string, rc registeredClass) {
	newNode := &rc
	newNode.name = n

	if obj.head == nil {
		obj.head = newNode
		return
	}

	curr := obj.head
	for curr.next != nil {
		curr = curr.next
	}

	curr.next = newNode
}

func (obj *registeredClassList) shift() {
	if obj.head == nil {
		return
	}

	if obj.head.next != nil {
		obj.head = obj.head.next
	} else {
		obj.head = nil
	}

}

func (obj *registeredClassList) classExists(n string) bool {
	// fmt.Println("searching for class " + n)
	if obj.head == nil {
		// fmt.Println("class not found")
		return false
	}

	current := obj.head
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

func (obj *registeredClassList) get(n string) *registeredClass {
	if obj.head == nil {
		return nil
	}

	current := obj.head
	for current != nil {
		if current.name == n {
			return current
		}
		current = current.next
	}
	return nil
}
