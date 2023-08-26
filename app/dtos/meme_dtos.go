package dtos

type MemeDTO struct {
	PostLink  string   `json:"postLink"`
	Subreddit string   `json:"subreddit"`
	Title     string   `json:"title"`
	Url       string   `json:"url"`
	Nsfw      bool     `json:"nsfw"`
	Spoiler   bool     `json:"spoiler"`
	Author    string   `json:"author"`
	Ups       int      `json:"ups"`
	Preview   []string `json:"preview"`
}

type MemeResponseDTO struct {
	Count int       `json:"count"`
	Memes []MemeDTO `json:"memes"`
}
