package model

type Page struct {
	Page_ID int32   `json:"page_id"`
	Title   string  `json:"page_title"`
	Text    string  `json:"page_text"`
	Product Product `json:"page_product"`
}
