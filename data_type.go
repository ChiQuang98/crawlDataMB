package main

type Category struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}
type Categories struct {
	Total int        `json:"total"`
	List  []Category `json:"categories"`
}
type Method struct {
	//Package string `json:"package"`
	//ClassTitle string `json:"classtitle"`
	TypeReturn  string `json:"typereturn"`
	MethodName  string `json:"methodname"`
	Description string `json:"description"`
}
type Methods struct {
	Package    string   `json:"package"`
	ClassTitle string   `json:"classtitle"`
	List       []Method `json:"methods"`
}

func newCategoires() *Categories {
	return &Categories{}
}
func newMethods() *Methods {
	return &Methods{}
}
