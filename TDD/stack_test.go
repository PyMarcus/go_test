package main 

import (
	"testing"
)

func TestEmpty(t *testing.T){
	s := NewStack()

	if !s.isEmpty(){
		t.Error("Stack was't empty.")
	}
}

func TestAdd(t *testing.T){
	s := NewStack()
	s.Add("Marcus")

	if s.isEmpty(){
		t.Error("Stack was empty")
	}
}

func TestZeroSize(t *testing.T){
	s := NewStack()
	if s.ZeroSize() != 0{
		t.Errorf("Expected 0 elements, but %d found!", s.size)
	}
}

func TestPopOne(t *testing.T){
	s := NewStack()
	s.Add("MV")

	if value := s.Pop(); value != "MV"{
		t.Errorf("Expected 'MV', but %v found!", value)
	}

	if s.ZeroSize() != 0{
		t.Errorf("Expected size = 0, but %d found!", s.size)
	}
}

func TestPopTwo(t *testing.T){
	s := NewStack()
	s.Add("A")
	s.Add("B")

	if value := s.Pop(); value != "B"{
		t.Errorf("Expected 'B', but %v found!", value)
	}

	if value := s.Pop(); value != "A"{
		t.Errorf("Expected 'A', but %v found!", value)
	}
  
	if s.ZeroSize() != 0{
		t.Errorf("Expected size = 0, but %d found!", s.size)
	}
}

func TestView(t *testing.T){
	s := NewStack()
	s.Add("A")
	s.Add("B")
	s.Add("C")

	if s.size != 3{
		t.Errorf("Expected size = 3, but %d found!", s.size)
	}

	if s.value[0] != "A" || s.value[1] != "B" || s.value[2] != "C"{
		t.Errorf(`Expected 0 => A 1 => B 2 => C,
				 but 0 => %s 1 => %s 2 => %s found!`, s.value[0], s.value[1], s.value[2])
	}
}