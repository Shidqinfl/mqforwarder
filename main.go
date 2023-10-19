package main

import (
	"flag"
	convert "mqforwarder/converter"
	printd "mqforwarder/debug"
	pubsub "mqforwarder/pubsub"
	tme "mqforwarder/timestamp"
)

func main() {
	Topic_sub := flag.String("TS", "sewadingin/v3/espRemote/powermeter", "Insert topic for subscribe") //
	Topic_pub := flag.String("TP", "smartpoint/v1", "Insert topic for publish")
	MQ_Port := flag.String("P", "1883", "Insert MQTT port")
	flag.Parse()

	pubsub.ConnecttoMqtt("tcp://mqtt.miota.io"+":"+*MQ_Port, *Topic_pub+"/mqforwarder/")
	// pubsub.ConnecttoMqtt("tcp://mqtt.miota.io:1883", "smartpoint/mqforwarder/")
	for {
		printd.Debug(1, tme.GetTime())
		s, err := pubsub.Subs(*Topic_sub)
		if err != nil {
			printd.Debug(3, "failed to subs "+err.Error())
		} else {
			payload, err := convert.ConvertPayload(s)
			if err == nil {
				pubsub.Publish_Data(*Topic_pub, payload)
				printd.Debug(1, payload)
			} else {
				printd.Debug(3, "waiting for payload")
			}

		}
		tme.Delay(1)
	}

}
