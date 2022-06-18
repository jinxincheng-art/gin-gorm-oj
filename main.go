package main

import (
	"gin-gorm-oj/router"
)

func main() {

	router.Router().Run() // listen and serve on 0.0.0.0:8080
}
