package domain

type Recti_Data struct {
	Timestamp string
	Id        string
	PZEM      []Pzem
}
type Pzem struct {
	Id      uint8
	Current float32
	Power   float32
	Energy  float32
}
type PM_data struct {
	Id   string
	Data []Data
}
type Data struct {
	Id      uint8
	Current float32
	Power   float32
	Energy  float32
}
