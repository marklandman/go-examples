// some first attempts at playing around with golang
package main

import (
        "fmt"
        "os"
        "strings"
        "io/ioutil"
        "path/filepath"
)

func readdir(dirname string, searchstring string){
        files, err := ioutil.ReadDir(dirname)
        if err != nil {
                //fmt.Println(err)
        }

        for _, f := range files {
                pname := strings.ToLower(filepath.Join(dirname,f.Name()))
                if (strings.Contains(pname,searchstring)) {
                fmt.Println(pname)
        }
        }
}

func main() {
        if len(os.Args) != 2 {
                fmt.Println("Usage:", os.Args[0], "<command> # searches for the command in path")
                return
        }
        lowersearch := strings.ToLower(os.Args[1])
        thepath := os.Getenv("PATH")
        paths := strings.Split(thepath, ":")
        for indexpath := range(paths) {
                readdir(paths[indexpath],lowersearch)

        }
}
