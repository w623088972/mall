package handler

type Page struct {
	PageNum int `json:"page_num"`
	Total   int `json:"total"`
}

type RequestPage struct {
	PageSize int `json:"page_size"`
	PageNum  int `json:"page_num"`
}

func GetPage(requestPage *RequestPage) {
	if requestPage.PageNum > 0 {
		requestPage.PageNum = (requestPage.PageNum - 1) * requestPage.PageSize
	} else {
		requestPage.PageNum = 0
	}
}
