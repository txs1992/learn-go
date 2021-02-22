package main

import "fmt"

type WalkRun interface {
	Walk()
	Run()
}

func (p person) Walk() {
	fmt.Printf("%s能走\n", p.name)
}

func (p person) Run() {
	fmt.Printf("%s能跑\n", p.name)
}
