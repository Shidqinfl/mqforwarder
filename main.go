package main

import (
	"encoding/json"
	"flag"
	"fmt"
	data "mqforwarder/converter"
	printd "mqforwarder/debug"
	pubsub "mqforwarder/pubsub"
	tme "mqforwarder/timestamp"
	"strconv"
)
var (
	summcell uint16
	summvolt float64
)

func main() {
	// var topic string
	mtopic := "pln_bms/110/v2/gejayan/"

	Topic_sub := flag.String("TS", mtopic+"285865A4AE30/cells", "Insert topic for subscribe") //
	Topic_pub := flag.String("TP", mtopic+"PC_NFL/cells", "Insert topic for publish")
	MQ_Port := flag.String("P", "2883", "Insert MQTT port")
	flag.Parse()

	pubsub.ConnecttoMqtt("tcp://mqtt.miota.io"+":"+*MQ_Port, *Topic_pub+"/mqforwarder/")
	// pubsub.ConnecttoMqtt("tcp://mqtt.miota.io:1883", "smartpoint/mqforwarder/")
	summcell = 0
	summvolt = 0
	for {
		printd.Debug(1, tme.GetTime())
		s, err := pubsub.Subs(*Topic_sub)
		if err == nil {
			// printd.Debug(2, "subs data: "+s)
			datas, err := data.GetCell(s)
			payload, _ := json.Marshal(datas)
			cell := data.C
			if err == nil && s != "" {
				for i := 1; i < 87; i++ {
					IDstr := strconv.FormatUint(uint64(cell[i].Id),10)
					Vstr := strconv.FormatFloat(cell[i].Voltage, 'f', -1, 32)
					Tstr := strconv.FormatFloat(cell[i].Temp, 'f', -1, 32)
					Lstr := strconv.FormatUint(uint64(cell[i].Liquid),10)
					fmt.Println("============[",i,"]============")
					fmt.Println("[ CELL ID          : " + IDstr+ " ]")
					fmt.Println("[ CELL VOLTAGE     : " + Vstr + " ]")
					fmt.Println("[ CELL TEMPERATURE : " + Tstr + " ]")
					fmt.Println("[ CELL LIQUID	   : " + Lstr + "  ]")
					fmt.Println("============[end]============")
					// if cell[i].Voltage != -1 || cell[i].Temp !=-1 {
					// 	summcell++
					// 	summvolt += cell[i].Voltage
					// }
				}
				// printd.Debug(1, "Total Cell : " + strconv.FormatUint(uint64(summcell), 10))
				// printd.Debug(1, "Total Voltage: " + strconv.FormatFloat(summvolt, 'f', -1, 32))
				pubsub.Publish_Data(*Topic_pub, payload)
			} //else {
			// 	printd.Debug(3, "waiting for payload")
			// }
		} // else {
		// 	printd.Debug(3, "failed to subs "+err.Error())
		// }
		tme.Delay(500)
	}

}
