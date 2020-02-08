package controllers

import (
	"github.com/KristiyanGK/cloudcooking/persistence/stores"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/KristiyanGK/cloudcooking/persistence"
	"os"
	"log"
	"github.com/joho/godotenv"
)

// App is struct for application 
type App struct {
	Router *mux.Router
	RecipeStore *stores.RecipeStore
}

//Init initiates the app
func (a *App) Init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	initDb()
	
	a.Router = mux.NewRouter()

	a.RegisterRoutes()

	a.RecipeStore = stores.NewRecipeStore()
}

//Run starts the rest api server
func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func initDb() {
	persistence.InitializeDb(os.Getenv("DB_DRIVER"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
}
