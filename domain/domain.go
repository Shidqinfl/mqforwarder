package domain

type AllCells struct {
	Timestamp string
	Dev_ID    string
	Cells     []Cell
}
type Cell struct {
	Id      uint8
	Voltage float64
	Temp    float64
	Liquid  uint8
}
type Summary struct {
	Dev_ID        string
	Timestamp     string
	Total_Voltage float32
	Total_Cell    uint16
	Not_Con_Cell uint16
	Not_Con_CELL_ID []uint
}
