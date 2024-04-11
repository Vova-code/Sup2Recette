package request

type RecipeRequest struct {
	Id    uint     `json:"id"`
	Title string   `json:"title"`
	Steps []string `json:"steps"`
}

type RecipeResponse struct {
	Id         uint     `json:"id"`
	Title      string   `json:"title"`
	Steps      []string `json:"steps"`
	ThumbsUp   int      `json:"thumbs_up"`
	ThumbsDown int      `json:"thumbs_down"`
}
