package main

import (
	"encoding/json"
	"flag"
	"fmt"
	data "mqforwarder/converter"
	printd "mqforwarder/debug"
	"mqforwarder/domain"
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
	listCellNull [87]uint
	cell_summary domain.Summary
	c_sum        []byte
)

func main() {
	// var topic string
	mtopic := "pln_bms/110/v2/gejayan/"

	Topic_sub := flag.String("TS", mtopic+"285865A4AE30/cells", "Insert topic for subscribe") //
	Topic_pub := flag.String("TP", mtopic+"PC_NFL", "Insert topic for publish")
	MQ_Port := flag.String("P", "2883", "Insert MQTT port")
	flag.Parse()

	pubsub.ConnecttoMqtt("tcp://mqtt.miota.io"+":"+*MQ_Port, *Topic_pub+"/mqforwarder/")
	// pubsub.ConnecttoMqtt("tcp://mqtt.miota.io:1883", "smartpoint/mqforwarder/")
	summcell = 0
	summvolt = 0
	// c:= cfinal
	printd.Debug(1, tme.GetTime())
	printd.Debug(2, "Processing")
	for {
		summcell = 0
		summvolt = 0
		cell_null = 0
		for i := 1; i < 87; i++ {
			listCellNull[i]= 0
		}
		
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
					} else {
						cell_null++
						listCellNull[i] = uint(cell[i].Id)

					}
				}
				fmt.Print(listCellNull)
				fmt.Println("")
				fmt.Println("==================================")
				printd.Debug(1, "Total Cell : "+strconv.FormatUint(uint64(summcell), 10))
				printd.Debug(1, "Total Voltage : "+strconv.FormatFloat(summvolt, 'f', -1, 32))
				printd.Debug(1, "Not Connected Cell : "+strconv.FormatUint(uint64(cell_null), 10))
				cell_summary.Dev_ID = "PC_QC"
				cell_summary.Timestamp = tme.GetTime()
				cell_summary.Total_Cell = summcell
				cell_summary.Not_Con_Cell = uint16(cell_null)
				cell_summary.Total_Voltage = float32(summvolt)
				cell_summary.Not_Con_CELL_ID = listCellNull[:]
				c_sum, _ = json.Marshal(cell_summary)
				fmt.Println("")
				sec := tme.GetSec()
				if sec%30 == 0 {
					pubsub.Publish_Data(*Topic_pub+"/cells", payload)
					pubsub.Publish_Data(*Topic_pub+"/summ", c_sum)

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
