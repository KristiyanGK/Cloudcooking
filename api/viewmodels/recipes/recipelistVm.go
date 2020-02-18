package recipes

type RecipeListVm struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Picture string `json:"picture"`
	CookingTime int `json:"cookingtime"`
	Category string `json:"category"`
}