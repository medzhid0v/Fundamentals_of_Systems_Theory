package main

import (
	"practice/api"
	"fmt"
)

func main() {
	data := api.GetData()
	// fmt.Printf("%+v\n", data)
	fmt.Printf("Title - %+v\n",data.Title)
	fmt.Printf("Year - %+v\n",data.Year)
	fmt.Printf("Released - %+v\n",data.Released)
	fmt.Printf("Runtime - %+v\n",data.Runtime)
	fmt.Printf("Genre - %+v\n",data.Genre)
	fmt.Printf("Director - %+v\n",data.Director)
	fmt.Printf("Writer - %+v\n",data.Writer)
	fmt.Printf("Actors - %+v\n",data.Actors)
	fmt.Printf("Language - %+v\n",data.Language)
	fmt.Printf("Country - %+v\n",data.Country)
	fmt.Printf("Awards - %+v\n",data.Awards)

}

