
package main

import "fmt"

type Channel struct {
    name    string
}


func main() {
var channels []Channel

//and finally add this object
channels = append(channels, Channel{name:"juan carlos anez mejias"})


     	fmt.Println(channels)
}
