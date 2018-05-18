//Echo2 - выводит аргументы командной строки
package main

import(
    "fmt"
    "os"
)
func main() {
    s, sep := "", ""
    for _, arg := range os.Args[0:] {
        s += sep + _ + arg
        sep = "\n"
    }
    fmt.Println(s)
}