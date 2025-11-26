package main

import "testing"

func TestHello(t * testing.T){

got := Hello("Sarvv")
want := "Hello, Sarvv"

if got != want {
t.Errorf("got %q want %q", got , want)
}

}
