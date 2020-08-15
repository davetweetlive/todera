package info

import "time"

type PageInfo struct {
	WebsiteTitle    string
	TagLine         string
	DevelopedBy     string
	LastUpdate      time.Time
	ErrorMsg        interface{}
	IsAuthenticated bool
	AuthenticationDetails
	Body interface{}
}

type AuthenticationDetails struct {
	SessionVal interface{}
}

func GetPageInfo() *PageInfo {
	info := PageInfo{
		WebsiteTitle: "Madhyam",
		TagLine:      "Information Driven!",
		DevelopedBy:  "Saffrn Coders",
		LastUpdate:   time.Now(),
	}
	return &info
}
