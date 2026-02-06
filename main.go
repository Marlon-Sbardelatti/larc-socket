package main

import (
	"bufio"
	"fmt"
	"github.com/joho/godotenv"
	"net"
	"os"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Não foi possível carregar as variáveis de ambiente.")
		os.Exit(1)
	}
	serverAddr := os.Getenv("ADDR")
	fmt.Println(serverAddr)

	message := NewMessage("SEND", "MESSAGE", "TESTE!")
	user := NewUser("6201", "gitxo")

	payload := BuildPayload(*message, *user)
	fmt.Println(payload)

	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// msg := "SEND MESSAGE 6201:gitxo:4109:TESTE TESTE!\n"
	msg := "GET MESSAGE 4109:ntnjb\n"
	conn.Write([]byte(msg))

	reader := bufio.NewReader(conn)
	resp, _ := reader.ReadString('\n')

	fmt.Printf("%s: %s", conn.RemoteAddr(), resp)
}
