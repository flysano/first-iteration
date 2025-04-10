package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(data string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, data := dataset {
		err := dp.Parse(data)
		if err != nil {
			log.Printf("data parsing error: %v", err)
			continue
		}
		info, err := dp.ActionInfo()
		if err != nil {
			log.Printf("error getting activity information: %v", err)
			continue
		}
		fmt.Println(info)
	}
}
