package main

func main() {

}

func waterProduct() {
	oxygen := oxygenProduct
}

func oxygenProduct() chan string {
	oxygen := make(chan string)
	go func() {
		oxygen <- "O"
	}()
}

func hydrogenProduct() {}
