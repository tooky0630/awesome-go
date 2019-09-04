package main

import "C"

/*
#include <stdio.h>

static void SayHello(const char* s) {
    puts(s);
}
*/
import "C"

func main() {
	println("hello cgo")

	C.puts(C.CString("Hello World")) // c string not free, maybe memory leak

	C.SayHello(C.CString("Hello My World")) // call customize c method
}
