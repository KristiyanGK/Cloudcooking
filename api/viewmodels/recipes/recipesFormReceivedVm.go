package recipes

type RecipesFormReceivedVm struct {
	Title string `json:"title"`
	Description string `json:"description"`
	Picture string `json:"picture"`
	CookingTime int `json:"cookingTime"`
	CategoryID string `json:"categoryid"`
	UsedProducts string `json:"usedProducts"`
}