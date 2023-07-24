package main 

import "fmt"

type Stack struct{
	empty bool
	size int
	value []string
}

func NewStack() *Stack{
	return &Stack{
		empty: true,
		size: 0,
	}
}

func (s Stack) isEmpty() bool{
	return s.empty
}

func (s *Stack) Add(value string){
	s.empty = false
	s.size++
	s.value = append(s.value, value) 
}

func (s Stack) ZeroSize() int{
	return s.size
}

func (s *Stack) Pop() string{
	s.size --
	v := s.value[len(s.value) - 1]
	s.value = s.value[0: s.size]
	return v
}

func (s Stack) View(){
	for _, v := range s.value{
		fmt.Println(v)
	}
}

func main(){

}