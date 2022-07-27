package anotherpackage

import (
	"log"

	"github.com/erhcc/testfor-init/somepackage"
)


func init()  {
	log.Println("init of anotherpackage called")
	
}

func CallOther() {
	log.Println("callanother of anotherpackage called")
	somepackage.Api()
}