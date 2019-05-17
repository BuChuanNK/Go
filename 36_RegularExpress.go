package main

import "fmt"
import "regexp"

func main() {

	match, _ := regexp.MatchString("p[a-z]+)ch", "peach")
	fmt.println(match)

	r, _ := regexp.Compile("p[a-z]+)ch")

	fmt.Println(r.MatchString("peach"))

	fmt.Println(r.FindString("peach punch"))

	fmt.Println(r.FindStringIndex("peach punch"))
}
