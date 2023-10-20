package main

import (
	"encoding/json"
	"flag"
	"fmt"
	data "mqforwarder/converter"
	printd "mqforwarder/debug"
	dom "mqforwarder/domain"
	pubsub "mqforwarder/pubsub"
	tme "mqforwarder/timestamp"
	"strconv"
)

var (
	summcell     uint16
	summvolt     float64
	cfinal       [86]dom.Cell
	cell_null    uint
	listCellNull []uint
	listCellOn   []uint
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
	// c:= cfinal
	for {
		summcell = 0
		summvolt = 0
		cell_null = 0
		printd.Debug(1, tme.GetTime())
		s, err := pubsub.Subs(*Topic_sub)
		if err == nil {
			// printd.Debug(2, "subs data: "+s)
			datas, err := data.GetCell(s)
			payload, _ := json.Marshal(datas)
			cell := data.C
			if err == nil && s != "" {
				for i := 1; i < 87; i++ {
					// IDstr := strconv.FormatUint(uint64(cell[i].Id), 10)
					// Vstr := strconv.FormatFloat(cell[i].Voltage, 'f', -1, 32)
					// Tstr := strconv.FormatFloat(cell[i].Temp, 'f', -1, 32)
					// Lstr := strconv.FormatUint(uint64(cell[i].Liquid),10)
					// fmt.Println("============[",i,"]============")
					// fmt.Println("[ CELL ID          : " + IDstr+ " ]")
					// fmt.Println("[ CELL VOLTAGE     : " + Vstr + " ]")
					// fmt.Println("[ CELL TEMPERATURE : " + Tstr + " ]")
					// fmt.Println("[ CELL LIQUID	   : " + Lstr + "  ]")
					// fmt.Println("============[end]============")

					if (cell[i].Voltage != -1 || cell[i].Temp != -1) && cell[i].Id != 0 {
						summcell++
						summvolt += cell[i].Voltage
						// c_on := append(listCellOn, uint(cell[i].Id))
						// // fmt.Println("===============[ CELL ON ]=============")
						// fmt.Print(c_on)

					} else {
						cell_null++
						listCellNull := append(listCellNull, uint(cell[i].Id))
						// if c0[i] !=0{
						fmt.Print(listCellNull)
						// fmt.Println("===============[ CELL OFF ]============")
						// }

					}

				}
				fmt.Println("")
				fmt.Println("==================================")
				printd.Debug(1, "Total Cell : "+strconv.FormatUint(uint64(summcell), 10))
				printd.Debug(1, "Total Voltage : "+strconv.FormatFloat(summvolt, 'f', -1, 32))
				printd.Debug(1, "Not Connected Cell : "+strconv.FormatUint(uint64(cell_null), 10))
				fmt.Println("")
				sec := tme.GetSec()
				if sec%10 == 0 {
					pubsub.Publish_Data(*Topic_pub, payload)
					pubsub.Publish_Data(*Topic_pub+"/cell_off", listCellNull)

				}

			} //else {
			// 	printd.Debug(3, "waiting for payload")
			// }
		} // else {
		// 	printd.Debug(3, "failed to subs "+err.Error())
		// }
		tme.Delay(500)
	}

}
