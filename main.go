package main

import (
	"log"

	"github.com/erhcc/testfor-init/anotherpackage"
	"github.com/erhcc/testfor-init/somepackage"
)

func init(){

	log.Println("init of main called")

}

func main() {
	log.Println("call anotherpackage")
	anotherpackage.CallOther()

	log.Println("call somepackage")
	somepackage.Api()

	log.Println("done")
}