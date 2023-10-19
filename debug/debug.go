package debug

import "log"

func Debug(d uint8, data string) { //1 debug, 2 info, 3
	switch d {
	case 1:
		log.Println("[Data ]-> ", "[ ",data," ]")
	case 2:
		log.Println("[INFO ]-> ", "[ ",data," ]")
	case 3:
		log.Println("[ERROR]-> ", "[ ",data," ]")
	}
}
