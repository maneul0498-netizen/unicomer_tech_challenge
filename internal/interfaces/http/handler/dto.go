package handler

import "encoding/xml"

type HoliDay struct {
	Date        string `json:"date" xml:"date"`
	Title       string `json:"title" xml:"title"`
	Type        string `json:"type" xml:"type"`
	Inalienable bool   `json:"inalienable" xml:"inalienable"`
	Extra       string `json:"extra" xml:"extra"`
}

type HoliDays []HoliDay

type Dto struct {
	XMLName xml.Name `json:"-" xml:"response"`

	Status string   `json:"status" xml:"status"`
	Data   HoliDays `json:"data" xml:"data>holiday"`
}

type ErrResponse struct {
	Message string `json:"message" example:"wallet not found"`
	Code    int    `json:"code" example:"404"`
}
