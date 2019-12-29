package main
 
import (
    "io/ioutil"
    "fmt"
    "log"
    "os"
    "strings"
)
//WriteFile writes data to file based on fileName
func WriteFile(fileName string, data string) {

    file, err := os.OpenFile(fileName, os.O_RDWR, 0644)
    if err != nil {
        log.Fatalf("failed opening file: %s", err)
    }
    defer file.Close()
     
    len, err := file.WriteString(data) // Write at 0 beginning
    if err != nil {
        log.Fatalf("failed writing to file: %s", err)
    }
    fmt.Println(len, "bytes written successfully")
    err = file.Close()
    if err != nil {
        fmt.Println(err)
    }
    
}
//ReadFile reads data from a file based on the fileName
func ReadFile(fileName string) string { 
    data, err := ioutil.ReadFile(fileName)
    if err != nil {
        log.Panicf("failed reading data from file: %s", err)
    }
    s := string(data)
    return s
}
 
func main() {

    path := os.Args[1]
    find := os.Args[2]
    replace := os.Args[3]

    data := ReadFile(path)
    
    fmt.Println(data)
    
    data = strings.Replace(data, find, replace, -1)
    
    fmt.Println(data)
    
    WriteFile(path, data)
}