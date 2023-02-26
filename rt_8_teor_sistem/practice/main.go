package main

import (
	"practice/api"
	"fmt"
)

func main() {
	data := api.GetData()
	
	fmt.Printf("%+v\n", data)
}

