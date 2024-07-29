package model

type RouteRequest struct {
	Url       string `json:"url" validate:"required,url"`
	Shortened string
	UserID    string
	Counter   int
}

func (r *RouteRequest) AddCounter() {
	r.Counter++
}
