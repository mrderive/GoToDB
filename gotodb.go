package main

import "fmt"
import "os"
import "net"

func main() {
    
    if len(os.Args) < 4 {
        fmt.Println("not enough arguments")
        fmt.Println("usage: go run gotodb.go <table> <file> <delimeter>")
        return
    }
    
    table_name := os.Args[1]
    file_name := os.Args[2]
    delim := os.Args[3][0]
    
    conn, e_conn := net.Dial("tcp", "localhost:27000")
    if e_conn != nil {
        fmt.Println("error connecting to database: ", e_conn.Error())
        return
    }

    rec := []byte{'i','n','s','e','r','t',']'}
    for i := 0; i < len(table_name); i++ {
        rec = append(rec, table_name[i])
    }
    rec = append(rec, ']')
    
    file_handle, _ := os.Open(file_name)
    buffer := make([]byte, 1)
    _, e_file := file_handle.Read(buffer)
    for e_file == nil {
        if buffer[0] == delim {
            rec = append(rec, ']')
        } else if buffer[0] == '\n' {
            rec = append(rec, '\n')
            send(rec, conn)
            rec = rec[0:8+len(table_name)]
        } else {
            rec = append(rec, buffer[0])
        }
        
        _, e_file = file_handle.Read(buffer)
    }
    
    file_handle.Close()
    conn.Close()
}

func send(rec []byte, conn net.Conn) {
    
    _, e_conn := conn.Write(rec)
    if e_conn != nil {
        fmt.Println(e_conn.Error())
        return
    }
    
    response := make([]byte, 99)
    _, e_conn = conn.Read(response)
    if e_conn != nil {
        fmt.Println(e_conn.Error())
        return
    }
    
    fmt.Print(string(response))
}
