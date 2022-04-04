package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ValidationRequest struct {
	Input *Message `json:"input"`
}

type ValidationResponse struct {
	Result string `json:"result"`
}

func main() {
	buf, err := os.ReadFile("message.bin")
	if err != nil {
		fmt.Printf("failed to read binary file: %v\n", err)
		os.Exit(1)
	}

	msg := Parse(buf)

	fmt.Println("--------------------")
	fmt.Printf("Message Type: %03d\n", msg.Type)
	fmt.Printf("PAN: %016d\n", msg.PrimaryAccountNumber)
	fmt.Printf("Processing Code: %06d\n", msg.ProcessingCode)
	fmt.Printf("Amount Transaction: %d\n", msg.AmountTransaction)
	fmt.Printf("Expiration Date: %04d\n", msg.ExpirationDate)
	fmt.Println("--------------------")

	vreq := ValidationRequest{
		Input: msg,
	}

	vreqBody, err := json.Marshal(vreq)
	if err != nil {
		fmt.Printf("failed to marshal message: %v\n", err)
		os.Exit(1)
	}

	res, err := http.Post("http://168.138.192.92/v1/data/iso8583/validation", "application/json", bytes.NewBuffer(vreqBody))
	if err != nil {
		fmt.Printf("failed to send validation request: %v\n", err)
		os.Exit(1)
	}
	defer res.Body.Close()

	vresBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("failed to read validation response: %v\n", err)
		os.Exit(1)
	}

	vres := ValidationResponse{}
	err = json.Unmarshal(vresBody, &vres)
	if err != nil {
		fmt.Printf("failed to unmarshal validation response: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Result: %s\n", vres.Result)
}
