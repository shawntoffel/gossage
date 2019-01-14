package gossage

import (
	"fmt"
	"log"
)

var Logger func(format string, a ...interface{})

func msglog(format string, a ...interface{}) {
	if Logger != nil {
		Logger(format, a...)
	} else {
		msg := fmt.Sprintf(format, a...)
		log.Printf("gossage: %s\n", msg)
	}
}
