package main

import "fmt"

func Add(a int, b int) int {
    return a + b
}

func f1() {
    fmt.Println("f1 function")
}
func f2(params int) {
    fmt.Println("f2 function", params)
}
func f3() (params1 int, params2 string) {
    fmt.Println("f3 function")
    return 10, "hello"
}
func f4(params int) (result int) {
    fmt.Println("f4 function", params)
    return params + 1
}

func main() {
    f1()
    f2(10)
    int_result, string_result := f3()
    fmt.Println(int_result, string_result)
    f4(10)
}
