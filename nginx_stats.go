package exporter



import (
    "fmt"
    "bufio"
    "io"
)

func ScanNginxStats(r io.Reader) {
    fmt.Println("Activated")

    s := bufio.NewScanner(r)

    fmt.Println(s.Scan())
}