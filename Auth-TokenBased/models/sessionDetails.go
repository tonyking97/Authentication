package models

/*
* @Author : Arputha Tony King P @Created at : Apr 19
* Definition : Schema to create the sessionDetails and insert in the MongoDB collection
 */

type SessionDetails struct {
	FF_Id      string    `bson:"ff_id"`
	FS_Id      string    `bson:"fs_id"`
	Token      string    `bson:"token"`
	Status     string    `bson:"status"`
	LoggedTime int64     `bson:"logged_Time"`
	ExpiryTime int64     `bson:"expiry_Time"`
	Ip_Details IpDetails `bson:"ip_details"`
	Referer    string    `bson:"referer"`

	Bot            bool   `bson:"isBot"`
	Mobile         bool   `bson:"isMobile"`
	BrowserName    string `bson:"browser_Name"`
	BrowserVersion string `bson:"browser_Version"`
	MozillaVersion string `bson:"mozillaVersion"`
	EngineName     string `bson:"engine_Name"`
	EngineVersion  string `bson:"engine_Version"`
	Localization   string `bson:"localization"`
	OsFullName     string `bson:"os_full_name"`
	OsName         string `bson:"os_name"`
	OsVersion      string `bson:"os_version"`
	Platform       string `bson:"platform"`

	TimeZone string `bson:"time_zone"`
	AppName  string `bson:"app_name"`

	UserAgent string `bson:"user_agent"`
}

type IpDetails struct {
	Ip            string   `json:"ip" bson:"ip"`
	Type          string   `json:"type" bson:"type"`
	ContinentCode string   `json:"continent_code" bson:"continent_code"`
	ContinentName string   `json:"continent_name" bson:"continent_name"`
	CountryCode   string   `json:"country_code" bson:"country_code"`
	CountryName   string   `json:"country_name" bson:"country_name"`
	RegionCode    string   `json:"region_code" bson:"region_code"`
	RegionName    string   `json:"region_name" bson:"region_name"`
	City          string   `json:"city" bson:"city"`
	Zip           string   `json:"zip" bson:"zip"`
	Latitude      float64  `json:"latitude" bson:"latitude"`
	Longitude     float64  `json:"longitude" bson:"longitude"`
	Location      location `json:"location" bson:"location"`
}

type location struct {
	GeonameId               int64       `json:"geoname_id" bson:"geoname_id"`
	Capital                 string      `json:"capital" bson:"capital"`
	Languages               []languages `json:"languages" bson:"languages"`
	CountryFlag             string      `json:"country_flag" bson:"country_flag"`
	CountryFlagEmoji        string      `json:"country_flag_emoji" bson:"country_flag_emoji"`
	CountryFlagEmojiUnicode string      `json:"country_flag_emoji_unicode" bson:"country_flag_emoji_unicode"`
	CallingCode             string      `json:"calling_code" bson:"calling_code"`
	IsEu                    bool        `json:"is_eu" bson:"is_eu"`
}

type languages struct {
	Code   string `json:"code" bson:"code"`
	Name   string `json:"name" bson:"name"`
	Native string `json:"native" bson:"native"`
}
