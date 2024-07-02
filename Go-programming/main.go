package main

import (
	"fmt"
	"go_programming/Methods"
	"image/color"
)

func main() {
	prim := Methods.Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	r := &Methods.Point{X: prim[0].X, Y: prim[0].Y}
	r.ScaleBy(2)
	fmt.Println(*r)
	fmt.Println(prim.Distance())

	var d1 Methods.Duration = 3600
	var d2 = d1
	fmt.Println(d1.Hours())
	fmt.Println(d2.Hours())

	var b1 Methods.Buffer
	_, err := b1.Write([]byte("hello"))
	if err != nil {
		fmt.Println("Error writing to buffer:", err)
		return
	}
	b2 := b1
	_, err = b2.Write([]byte(" world"))
	if err != nil {
		fmt.Println("Error writing to buffer:", err)
		return
	}
	fmt.Println(string(b1.Buf))
	fmt.Println(string(b2.Buf))

	list := &Methods.IntList{
		Value: 1,
		Tail: &Methods.IntList{
			Value: 2,
			Tail: &Methods.IntList{
				Value: 3,
				Tail:  nil,
			},
		},
	}

	fmt.Println(list.Sum())
	var emptyList *Methods.IntList
	fmt.Println(emptyList.Sum())

	m := Methods.Values{
		"lang": {"en"},
		"item": {"1", "2"},
	}
	m.Add("item", "3")
	m.Add("key", "value")

	fmt.Println(m.Get("lang"))
	fmt.Println(m.Get("q"))
	fmt.Println(m.Get("item"))
	fmt.Println(m["item"])

	var mNil Methods.Values
	fmt.Println(mNil.Get("item"))

	if mNil == nil {
		mNil = make(Methods.Values)
	}
	mNil.Add("item", "3")
	fmt.Println(mNil.Get("item"))

	var cp Methods.ColoredPoint
	cp.Point.X = 1
	fmt.Println(cp.X)
	red := color.RGBA{R: 255, A: 255}
	blue := color.RGBA{B: 255, A: 255}
	var p = Methods.ColoredPoint{Point: Methods.Point{X: 1, Y: 1}, Color: red}
	var q = Methods.ColoredPoint{Point: Methods.Point{X: 1, Y: 1}, Color: blue}
	fmt.Println(p.Distance(q.Point))
	p.ScaleBy(2)
	q.ScaleBy(3)
	fmt.Println(p.Distance(q.Point))

	var s Methods.IntSet
	s.Add(1)
	s.Add(144)
	fmt.Println(s.Has(1))
	fmt.Println(s.Has(144))
	fmt.Println(s.Has(100))

}
