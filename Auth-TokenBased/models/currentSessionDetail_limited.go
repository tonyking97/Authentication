package models

type CurrentSessionDetail_limited struct {
	Active_Sessions int64   `json:"active_sessions"`
	Active_Tokens   int64   `json:"active_token"`
	Ip              string  `json:"ip"`
	City            string  `json:"city"`
	Latitude        float32 `json:"latitude"`
	Longitude       float32 `json:"longitude"`
}
