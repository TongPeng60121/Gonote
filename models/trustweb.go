package models

import "time"

type Url struct {
	Url string `json:"Url"`
}

type Session struct {
	SessionID int   `json:"SessionID"`
	TrustWeb  []Url `json:"TrustWeb"`
	ClientID  int   `json:"ClientID"`
}

type Trustweb struct {
	SessionID int       `json:"sessionID"  gorm:"column:SessionID"`
	ClientID  int       `json:"clientID"  gorm:"column:ClientID"`
	Url       string    `json:"url"`
	Cdate     time.Time `json:"cdate"`
}

type Searchtrustweb struct {
	Url         string    `json:"url"`
	Cdate       time.Time `json:"cdate"`
	CdateString string    `json:"cdateString"`
}

type UrlCount struct {
	Url   string `json:"url"`
	Count int    `json:"count"`
}
