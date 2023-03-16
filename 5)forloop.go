package main

import "fmt"

func main() {

    i := 1
    for i <= 3 {
        fmt.Println(i)
        i = i + 1
    }

    for j := 5; j <= 9; j++ {
        fmt.Println(j)
   for a := 5 ; a <=7; a++ {
        fmt.Println(a)
        
    }
  }
	for j := 10; j >= 5; j-- {
        fmt.Println(j)
    }

   
    for n := 0; n <= 5; n++ {
        if n%2 == 0 {
            continue
        }
        fmt.Println(n)
    }
}