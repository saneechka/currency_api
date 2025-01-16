package models

type Rate struct {
	Cur_ID           int     `json:"Cur_ID"`
	Date             string  `json:"Date"`
	Cur_Abbreviation string  `json:"Cur_Abbreviation"`
	Cur_Scale        int     `json:"Cur_Scale"`
	Cur_Name         string  `json:"Cur_Name"`
	Cur_OfficialRate float64 `json:"Cur_OfficialRate"`
}
