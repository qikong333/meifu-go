package apis

import "time"

type Banner struct {
	Id string `json:"id",`
	Name string `json:"name"`
	Image string `json:"image"`
	Time time.Time `json:"time"`

}

type New struct {
	Id string `json:"id"`
	Title string `json:"title"`
	Time time.Time `json:"time"`
	Content string `json:"content"`
	Author string `json:"author"`
	Source string `json:"source"`
}

type Product struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Img string `json:"img"`
	Info string `json:"info"`
	Content string `json:"content"`
	Time time.Time `json:"time"`
	Author string `json:"author"`

} 
