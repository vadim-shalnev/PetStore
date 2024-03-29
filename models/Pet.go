package models

type Pets []Pet

type Pet struct {
	ID       int64    `json:"id"`
	Category Category `json:"category"`
	Name     string   `json:"name"`
	//PhotoUrls []string   `json:"photoUrls"`
	Tags   []Category `json:"tags"`
	Status string     `json:"status"`
}

type Category struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
