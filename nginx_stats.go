package exporter



import (
    "fmt"
    "bufio"
    "io"
    "strings"
    "strconv"
)

type NginxStub struct {
    ActiveConnection float64

    Connections []Work
}

type Work struct{

    Total float64
    // is it reading / writing or waiting for request. 
    // what work the Nginx connection is doing.
    NginxWork string
}

func ScanNginxStats(r io.Reader) ([]NginxStub){
    // fmt.Println("Activated")

    s := bufio.NewScanner(r)

    // fmt.Println(s.Scan())

    var hist_stats []NginxStub
    var stats NginxStub
    var conns []Work

    for s.Scan() {
        field := strings.Fields(string(s.Bytes()))
        // fmt.Println(field)
        if len(field) == 3 && field[0] == "Active" {
            client_connected, _ := strconv.ParseFloat(field[2], 64)
            // fmt.Println("No. of active connections ", client_connected)
            stats.ActiveConnection = client_connected
        }

        if field[0] == "Reading:" {
            read_tt, _ := (strconv.ParseFloat(field[1], 64))
            write_tt, _ := (strconv.ParseFloat(field[1], 64))
            wait_tt, _ := (strconv.ParseFloat(field[1], 64))
            read_con := Work{NginxWork: "reading", Total: read_tt}
            write_con := Work{NginxWork: "writing",Total: write_tt}
            wait_con := Work{NginxWork: "waiting",Total: wait_tt}

            conns = append(conns, read_con, write_con, wait_con)
            stats.Connections = conns
        }
    }
    hist_stats = append(hist_stats, stats)
    fmt.Println(len(hist_stats))
    return hist_stats
}