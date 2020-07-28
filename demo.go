package main

import "fmt"

type machine interface {
	aircraft
	car
}

type aircraft interface {
	flight()
}
type car interface {
	run()
}

type ferrari struct {
	brand string
}

type tesla struct {
	brand string
}

func (f ferrari) run() {
	fmt.Printf("%s 速度100迈\n", f.brand)
}

func (t tesla) run() {
	fmt.Printf("%s 速度200迈\n", t.brand)
}

func drive(c car) {
	c.run()
}

func main() {
	f := ferrari{"法拉利"}
	t := tesla{"特斯拉电动车"}
	drive(f)
	drive(t)
}
