package request

type RecipeRequest struct {
	Title string   `json:"title"`
	Steps []string `json:"steps"`
}

type RecipeResponse struct {
	Title       string   `json:"title"`
	Steps       []string `json:"steps"`
	Evaluations int      `json:"evaluations"`
}
