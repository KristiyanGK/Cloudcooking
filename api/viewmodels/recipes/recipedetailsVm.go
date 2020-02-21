package recipes

type RecipeDetailsVm struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Picture string `json:"picture"`
	CookingTime int `json:"cookingTime"`
	Category string `json:"category"`
	UsedProducts string `json:"usedProducts"`
	User string `json:"user"`
}