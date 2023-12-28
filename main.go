package main

import "github.com/gookit/config/v2"

func main() {
	keys := []string{"config"}
	err := config.LoadFlags(keys)
	if err != nil {
		panic(err)
	}

	setupDB()
	router := setupRouter()

	_ = router.Run(":8080")
	closeDB()
}
