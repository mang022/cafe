package main

import "strconv"

func main() {
	setupConfig()
	setupDB()
	router := setupRouter()

	_ = router.Run(":" + strconv.Itoa(conf.Host.Port))
	closeDB()
}
