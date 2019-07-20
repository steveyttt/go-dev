package main

import (
	"fmt"
	"time"
)

func playing() {

	t := time.Now()
	fmt.Println(t)
	//time.kitchen is a constant you can reference from type time
	//from a value of type time (t), run method Format & request constant kitche from package time (time.kitchen)
	tf := t.Format(time.Kitchen)
	fmt.Println(tf)

}
