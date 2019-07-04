package models

/*
* @Author : Arputha Tony King P @Created at : Apr 19
* Definition : Schema to create the sessionDetails and insert in the MongoDB collection
 */

type SessionDetails_limited struct {
	Token         string `json:"token" bson:"token"`
	Status        string `json:"status" bson:"status"`
	LoggedTime    int64  `json:"logged_Time" bson:"logged_Time"`
	LoggedTimeUTC string `json:"logged_Time_utc" bson:"logged_Time_utc"`
	ExpiryTime    int64  `json:"expiry_Time" bson:"expiry_Time"`
	ExpiryTimeUTC string `json:"expiry_Time_utc" bson:"expiry_Time_utc"`
	//Ip_Details IpDetails `bson:"ip_details"`

	Mobile         bool   `json:"isMobile" bson:"isMobile"`
	BrowserName    string `json:"browser_Name" bson:"browser_Name"`
	BrowserVersion string `json:"browser_Version" bson:"browser_Version"`
	OsName         string `json:"os_name" bson:"os_name"`
	Platform       string `json:"platform" bson:"platform"`
}
