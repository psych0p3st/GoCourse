package main

import "fmt"

func main() {
	var dist float64
	var time int
	fmt.Scan(&dist, &time)
	fmt.Println(dist / float64(time))

}
