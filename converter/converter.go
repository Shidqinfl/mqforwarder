package convert

import (
	"encoding/json"
	printd "mqforwarder/debug"
	domain "mqforwarder/domain"
	// "strconv"
)
var C [87]domain.Cell
func GetCell(p string) (Cell interface{}, err error) {
	all_c := domain.AllCells{}
	if p != ""{
		err = json.Unmarshal([]byte(p), &all_c)
		if err != nil {
			printd.Debug(3, "can't unmarshall all_c because "+err.Error())
		}else {
			printd.Debug(2, "=== Data ===")
			printd.Debug(2, "ID ESP = "+all_c.Dev_ID)
			printd.Debug(2, "TS ESP = "+all_c.Timestamp)
			// printd.Debug(2, "CELL ESP = " + all_c.Cells)
			// fmt.Println("cn: ", all_c.Cells)
			
			for i := 0; i < 6; i++ {
				id_x := all_c.Cells[i].Id
				C[id_x].Id = all_c.Cells[i].Id
				C[id_x].Voltage = all_c.Cells[i].Voltage
				C[id_x].Temp = all_c.Cells[i].Temp
				C[id_x].Liquid = all_c.Cells[i].Liquid

				if id_x == 86{
					break;
				}
			}
		}
	}
	return C, err
}
