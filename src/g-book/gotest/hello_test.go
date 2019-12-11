package main

import "testing"

func TestHello(t *testing.T) {
	t.Log("hello,world")
}

func TestA(t *testing.T) {
	t.Log("hello,world -A")
}

func TestAK(t *testing.T) {
	t.Log("hello,world -AK")
}
