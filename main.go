package main

func main() {

	go brodudp()
	go listentcp()
	select {}
}
