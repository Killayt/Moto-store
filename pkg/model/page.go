package model

type Page struct {
	Title   string  `json:"title"`
	Text    string  `json:"text"`
	Product Product `json:"product"`
}

// func (p Page) CreateNewPage(title, text string, product Product) *Page {
// 	return &Page{
// 		Title:   title,
// 		Text:    text,
// 		Product: product,
// 	}
// }
