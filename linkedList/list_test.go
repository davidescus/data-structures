package main

import (
	"reflect"
	"testing"
)

func TestList_PushBack(t *testing.T) {
	type fields struct {
		head *Node
		size int
	}
	type args struct {
		n *Node
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &List{
				head: tt.fields.head,
				size: tt.fields.size,
			}
			l.PushBack(tt.args.n)
		})
	}
}

func TestList_PushFront(t *testing.T) {
	type fields struct {
		head *Node
		size int
	}
	type args struct {
		n *Node
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &List{
				head: tt.fields.head,
				size: tt.fields.size,
			}
			l.PushFront(tt.args.n)
		})
	}
}

func TestList_Back(t *testing.T) {
	type fields struct {
		head *Node
		size int
	}
	tests := []struct {
		name   string
		fields fields
		want   *Node
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &List{
				head: tt.fields.head,
				size: tt.fields.size,
			}
			if got := l.Back(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("List.Back() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_Front(t *testing.T) {
	type fields struct {
		head *Node
		size int
	}
	tests := []struct {
		name   string
		fields fields
		want   *Node
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &List{
				head: tt.fields.head,
				size: tt.fields.size,
			}
			if got := l.Front(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("List.Front() = %v, want %v", got, tt.want)
			}
		})
	}
}

//func Test_main(t *testing.T) {
//	tests := []struct {
//		name string
//	}{
//	// TODO: Add test cases.
//	}
//	for range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			main()
//		})
//	}
//}
