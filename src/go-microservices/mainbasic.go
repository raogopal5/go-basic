package main

import (
	"fmt"

	"rsc.io/quote"
)

func calcRectange(legth, width float64) (float64, float64) {

	area := legth * width
	perameter := 2 * (legth * width)
	return area, perameter
}

func mainbasic() {
	//first way declartion
	var x int = 10
	var y, z = 2, 5
	var name string
	name = "goland development"
	//second way declartion
	isWorking := true
	n := 10
	var dayOftheMonth = map[string]int{"Jan": 31, "feb": 28}

	fmt.Println("Hello go")
	fmt.Println(dayOftheMonth)
	fmt.Println(quote.Go())
	fmt.Println(x, y, z, name, isWorking, n)
	a, p := calcRectange(3, 6)
	fmt.Println("Area is %f and parameter is %f", a, p)
}
