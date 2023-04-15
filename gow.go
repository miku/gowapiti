package main

/*
#include "src/wapiti.h"
*/
import "C"
import "fmt"

func main() {
	fmt.Println(C.VERSION)
}
