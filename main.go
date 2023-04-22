package main

import (
	"api-book-go_gin/routers"
)

func main(){
	var PORT = ":8080"

	routers.StartServer().Run(PORT)
}