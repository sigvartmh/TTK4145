package main

import (
    "fmt"
)

func subProcess(count *int){
 *count = *count + 1
}

func main(){
    count := 0
    subProcess(&count)
    count = count + 23
    fmt.Println("count is:", count)
}
