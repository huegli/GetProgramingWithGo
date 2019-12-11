// How long does it take to get to Mars?
package main

import "fmt"

func main() {
    const lightspeed = 299792 // km/s
    var distance = 56000000   // km

    fmt.Println(distance/lightspeed, "seconds")

    distance = 401000000
    fmt.Println(distance/lightspeed, "seconds")
    
    const spacexspeed = 100800/3600 // km/h
    distance = 56000000   // km

    fmt.Println(distance/spacexspeed/(3600*24), "days")

    distance = 401000000
    fmt.Println(distance/spacexspeed/(3600*24), "days")
}
