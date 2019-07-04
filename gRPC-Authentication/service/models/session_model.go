package models

type SessionDetails struct {
	Id             string         `json:"id" bson:"_id,omitempty"`
	FfId           string         `bson:"ff_id"`
	UserDetails    UserDetails    `json:"user_details" bson:"user_details"`
	Status         string         `bson:"status"`
	LoggedTime     int64          `bson:"logged_Time"`
	ExpiryTime     int64          `bson:"expiry_Time"`
	IpDetails      IpDetails      `bson:"ip_details"`
	Referer        string         `bson:"referer"`
	BrowserDetails BrowserDetails `json:"browser_details" bson:"browser_details"`
	UserAgent      string         `bson:"user_agent"`
}

type UserDetails struct {
	Username  string `json:"username" bson:"username"`
	Email     string `json:"email" bson:"email"`
	FirstName string `json:"first_name" bson:"first_name"`
	LastName  string `json:"last_name" bson:"last_name"`
}

type BrowserDetails struct {
	Bot            bool   `bson:"isBot"`
	Mobile         bool   `bson:"isMobile"`
	BrowserName    string `bson:"browser_Name"`
	BrowserVersion string `bson:"browser_version"`
	MozillaVersion string `bson:"mozilla_version"`
	EngineName     string `bson:"engine_Name"`
	EngineVersion  string `bson:"engine_Version"`
	Localization   string `bson:"localization"`
	OsFullName     string `bson:"os_full_name"`
	OsName         string `bson:"os_name"`
	OsVersion      string `bson:"os_version"`
	Platform       string `bson:"platform"`
	TimeZone       string `bson:"time_zone"`
	AppName        string `bson:"app_name"`
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
	Location      Location `json:"location" bson:"location"`
}

type Location struct {
	GeoNameId               int64       `json:"geo_name_id" bson:"geo_name_id"`
	Capital                 string      `json:"capital" bson:"capital"`
	Languages               []Languages `json:"languages" bson:"languages"`
	CountryFlag             string      `json:"country_flag" bson:"country_flag"`
	CountryFlagEmoji        string      `json:"country_flag_emoji" bson:"country_flag_emoji"`
	CountryFlagEmojiUnicode string      `json:"country_flag_emoji_unicode" bson:"country_flag_emoji_unicode"`
	CallingCode             string      `json:"calling_code" bson:"calling_code"`
	IsEu                    bool        `json:"is_eu" bson:"is_eu"`
}

type Languages struct {
	Code   string `json:"code" bson:"code"`
	Name   string `json:"name" bson:"name"`
	Native string `json:"native" bson:"native"`
}
