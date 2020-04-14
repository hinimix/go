package demo

import "testing"

func TestGetArea(t *testing.T){
	area := GetArea(40, 50)
	if area != 2000 {
		t.Error("test failed")
	}
}

func BenchmarkGetArea(b *testing.B) {
	for i:=0; i<b.N; i++{
		GetArea(40,50)
	}
}