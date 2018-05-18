//Echo3 - Вывод аргументов командной строки
package main
import(
    "fmt"
    "os"
    "strings"
)
func main() {
    fmt.Println(strings.Join(os.Args[0:], "\n"))
}