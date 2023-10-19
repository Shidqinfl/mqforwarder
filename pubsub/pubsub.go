package pubsub

import (
	// "os"
	// "os/signal"
	// "syscall"
	printd "mqforwarder/debug"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var (
	public_opts   = MQTT.NewClientOptions()     //for mqtt client options
	public_client = MQTT.NewClient(public_opts) //for mqtt client
)

var Payloads string

var Public struct {
	Dns      string
	Clientid string
}

var Connection_MQTT struct {
	Public bool
}

var public_connectHandler MQTT.OnConnectHandler = func(client MQTT.Client) {
	//Log_Level(3, "func public_connectHandler            ", "MQTT Connected")
	Connection_MQTT.Public = true
	// Connection.Mqtt = "OK"
}
var public_connectLostHandler MQTT.ConnectionLostHandler = func(client MQTT.Client, err error) {
	// Log_Level(2, "func public_connectLostHandler", "Connection MQTT lost . . .")
	Connection_MQTT.Public = false
	// Connection.Mqtt = "E1"
}
var f MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	// fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
	Payloads = string(msg.Payload())
	// text:= fmt.Sprintf("this is result msg #%d!", knt)
	// knt++
}

func ConnecttoMqtt(host string, topic string) {
	Public.Dns = host
	Public.Clientid = "MIOTA"

	public_opts = MQTT.NewClientOptions()
	public_opts.AddBroker(Public.Dns).SetClientID(Public.Clientid).SetUsername("iot").SetPassword("iot")
	public_opts.SetKeepAlive(15 * time.Second)
	public_opts.SetCleanSession(true)
	public_opts.SetConnectTimeout(15 * time.Second)
	public_opts.SetConnectRetry(true)
	public_opts.SetConnectRetryInterval(15 * time.Second)
	public_opts.SetAutoReconnect(true).SetMaxReconnectInterval(15 * time.Second)
	public_opts.SetDefaultPublishHandler(f)
	public_opts.OnConnect = public_connectHandler
	public_opts.OnConnectionLost = public_connectLostHandler

	public_opts.SetWill(topic, "mqforwarder", 0, false)

	public_client = MQTT.NewClient(public_opts)
	tokenConn := public_client.Connect()
	if tokenConn.WaitTimeout(15*time.Second) && tokenConn.Error() != nil { //WaitTimeout(10*time.Second)
		// Log_Level(2, "func Public_connect", "MQTT Time Out")
	}
	// Log_Level(3, "func Public_connect-client.IsConnected", strconv.FormatBool(public_client.IsConnected()))

	time.Sleep(1 * time.Second)
	if !Connection_MQTT.Public {
		// Log_Level(2, "func Public_connect                   ", "MQTT NOT Connected")
	}
}

func Publish_Data(tops string, message interface{}) (err error) {
	Sending := public_client.Publish(tops, 0, false, message)
	if Sending.WaitTimeout(3*time.Second) != true {
		// log.Println("Failed to Publish", tops)
		printd.Debug(3, "Failed to Publish")
	} else {
		// log.Println("[published to topic ", tops, " ]")
		printd.Debug(2, "Published to topic " + tops)

	}

	return err
}
func Subs(topic string) (p string, err error) {
	token := public_client.Subscribe(topic, 0, f)
	// fmt.Println(Payloads)
	// printd.Debug(2, Payloads)
	err = token.Error()
	p = Payloads
	return p, err
}
