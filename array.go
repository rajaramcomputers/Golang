package main 

import "fmt"


func main (){
 // define arrays
//  var arr [4]string // ' '
// //  fmt.Println("arr",arr)

//  // var a [3]int 
// //  fmt.Println("int arr", a)


//  arr[0] = "Ahmed"
//  arr[2] = "sara"
//  arr[3] = "adam"

// //  fmt.Println("arr",arr)

//  arr = [4]string{"jan", "jack", "hady", "sara"}
//  fmt.Println("re assigen arr", arr)

//  var num = [...]int{1, 2, 3, 4, 5}
//  fmt.Println("num", num)

//  names := [2]string{"string"}
//  fmt.Println(names)

//  // len()
//  arrLen := len(arr)
//  fmt.Println(arrLen)

// -- 2D Arrays


XoBoard := [3][3] string{}
//fmt.Println(XoBoard)


XoBoard = [3][3] string {
	[3]string{ "-", "-", "x" },
	[3]string{ "-", "o", "-" },
	[3]string{ "x", "o", "-" },
} 

//fmt.Println(XoBoard)
fmt.Println(XoBoard[0])
fmt.Println(XoBoard[1])
fmt.Println(XoBoard[2])

//

XoBoard[0][1] = "x"
fmt.Println(XoBoard[0])
fmt.Println(XoBoard[1])
fmt.Println(XoBoard[2])

}