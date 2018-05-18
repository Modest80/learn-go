//Echo1 выводит аргументы командной строки
package main

import(
    "fmt"
    "os"
)
func main() {
    //var s, sep string
    for i := 0; i < len(os.Args); i++ {
        //s += sep + i + " " +  os.Args[i]
        //sep = "\n"
        fmt.Print(i, " ")
        fmt.Println(os.Args[i])
    }
    //fmt.Println(s)
}