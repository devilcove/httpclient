// copywrite 2022 Matthew R Kasun mkasun@nusak.ca
package main

func main() {
	SigningKey = []byte("secretkey")
	router := SetupRouter()
	router.Run("127.0.0.1:8010")
}
