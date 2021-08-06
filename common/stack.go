package common

import "fmt"

type Element interface {
} //可存入任何类型

type Stack struct {
	list []Element
}

//初始化栈
func NewStack() *Stack {
	return &Stack{
		list: make([]Element, 0),
	}
}

func (s *Stack) Len() int {
	return len(s.list)
}

//判断栈是否空
func (s *Stack) IsEmpty() bool {
	if len(s.list) == 0 {
		return true
	} else {
		return false
	}
}

//入栈
func (s *Stack) Push(x interface{}) {
	s.list = append(s.list, x)
}

//连续传入
func (s *Stack) PushList(x []Element) {
	s.list = append(s.list, x...)
}

//出栈
func (s *Stack) Pop() Element {
	if len(s.list) <= 0 {
		fmt.Println("Stack is Empty")
		return nil
	} else {
		ret := s.list[len(s.list)-1]
		s.list = s.list[:len(s.list)-1]
		return ret
	}
}

//返回栈顶元素，空栈返nil
func (s *Stack) Top() Element {
	if s.IsEmpty() == true {
		fmt.Println("Stack is Empty")
		return nil
	} else {
		return s.list[len(s.list)-1]
	}
}

//清空栈
func (s *Stack) Clear() bool {
	if len(s.list) == 0 {
		return true
	}
	for i := 0; i < s.Len(); i++ {
		s.list[i] = nil
	}
	s.list = make([]Element, 0)
	return true
}

//打印测试
func (s *Stack) Show() {
	len := len(s.list)
	for i := 0; i != len; i++ {
		fmt.Println(s.Pop())
	}
}
