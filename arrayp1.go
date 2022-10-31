package main

import (
	"fmt"
)

func main() {
	stu_array := [4]int{1, 2, 3, 5}
	fmt.Println("Array", stu_array)

	var course_array [4]string
	course_array[0] = "Web Development"
	course_array[1] = "Introduction to Mathematics"
	course_array[2] = "Fundamentals of Programming"
	course_array[3] = "Programming Fundamentals"

	fmt.Println("Interesting Array : ", course_array)
}
