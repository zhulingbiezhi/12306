package main

import (
	"log"

	"github.com/denisbrodbeck/machineid"
)

func getpackStr(cookieCode string) {
	getMachineCode()

}
func getMachineCode() string {
	id, err := machineid.ProtectedID("myAppName")
	if err != nil {
		log.Fatal(err)
	}
	return id
}
