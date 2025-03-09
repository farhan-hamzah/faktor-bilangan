package main
import "fmt"

func faktor(x int, keluaran *int){
    *keluaran = 1
    for *keluaran <= x{
        if x%*keluaran == 0{
            fmt.Print(" ", *keluaran)
        }
        *keluaran++
    }
}

func main(){
    var a, b int
    fmt.Scan(&a)
    faktor(a, &b)
}