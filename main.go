package main

func main() {
	router := setupRouter()

	_ = router.Run(":8080")
}
