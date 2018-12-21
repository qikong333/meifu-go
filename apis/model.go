package apis

type Banner struct {
	Id int `json:"id",`
	Name string `json:"name"`
	Image string `json:"image"`
}

type New struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Time string `json:"time"`
	Content string `json:"content"`
	Author string `json:"author"`
	Source string `json:"source"`
}

type Product struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Img string `json:"img"`
	Info string `json:"info"`
	Content string `json:"content"`
} 
