package convert

import (
	"encoding/json"
	printd "mqforwarder/debug"
	domain "mqforwarder/domain"
	tme "mqforwarder/timestamp"
)

func ConvertPayload(p string) (strJson string, err error) {
	pm := domain.PM_data{}
	// d := domain.Data{}
	recti := domain.Recti_Data{}
	err = json.Unmarshal([]byte(p), &pm)
	if err != nil {
		printd.Debug(3, "can't convert because " + err.Error())
		// fmt.Println(err)
	} else {
		printd.Debug(2, "Converted data")
		// fmt.Println("ID: ", pm.Id)
		// fmt.Println("Data: ", pm.Data)
		pmdata := pm.Data[0]
		// fmt.Println("id pzem: ", pmdata.Id)
		// fmt.Println("current: ", pmdata.Current)
		// fmt.Println("energy: ", pmdata.Energy)
		// fmt.Println("power: ", pmdata.Power)
		dpzem := domain.Pzem{}
		dpzem2 := domain.Pzem{}
		dpzem3 := domain.Pzem{}
		recti.Timestamp = tme.GetTime()
		recti.Id = pm.Id
		dpzem.Id = pmdata.Id
		dpzem.Current = pmdata.Current
		dpzem.Energy = pmdata.Energy
		dpzem.Power = pmdata.Power
		recti.PZEM = append(recti.PZEM, dpzem, dpzem2, dpzem3)
		// fmt.Println(recti)
		
		jpayload, err := json.Marshal(recti)
		if err != nil {
			printd.Debug(3,"can't Marshal data because "+ err.Error())
		}
		strJson = string(jpayload)
		printd.Debug(1, strJson)
	}

	return strJson, err
}
