package main

import (
    "fmt"
	"os"
    "os/exec"
    "sync"

)

func main() {
 setpirglives()
    wg := new(sync.WaitGroup)
    wg.Add(1)

	fmt.Println(" Run program server.go")
	
	out, err := exec.Command("..\\srir-go-soap_server.exe").Output()
	  if err != nil {
        fmt.Println("error occured")
        fmt.Printf("%s", err)
    }
	    fmt.Printf("%s", out)
		    wg.Done()

    wg.Wait()
}
// nadaj podstawowe prawa serverowi
func setpirglives() {
	
	err := os.Chmod("..\\srir-go-soap_server.exe", 0444)
	if err != nil {
     fmt.Println(err)
	 os.Exit(0)
	}
	
  fmt.Println("Premmision Succes")

}