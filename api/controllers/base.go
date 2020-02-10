package controllers

import (
	"github.com/go-playground/locales/en"
	"github.com/KristiyanGK/cloudcooking/persistence/seeder"
	"github.com/KristiyanGK/cloudcooking/persistence/stores"
	"github.com/go-playground/validator/v10"
	ut "github.com/go-playground/universal-translator"
	en_translations "github.com/go-playground/validator/v10/translations/en"
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
	Validator  *validator.Validate
	Translator ut.Translator
}

//Init initiates the app
func (a *App) Init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	initDb()

	a.initValidator()
	
	a.Router = mux.NewRouter()

	a.RegisterRoutes()

	a.RecipeStore = stores.NewRecipeStore()
}

//Run starts the rest api server
func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) initValidator() {
	// Create and configure validator and translator
	a.Validator = validator.New()
	eng := en.New()
	var uni *ut.UniversalTranslator
	uni = ut.New(eng, eng)
	// this is usually know or extracted from http 'Accept-Language' header
	// also see uni.FindTranslator(...)
	var found bool
	a.Translator, found = uni.GetTranslator("en")
	if !found {
		log.Fatal("translator not found")
	}
	if err := en_translations.RegisterDefaultTranslations(a.Validator, a.Translator); err != nil {
		log.Fatal(err)
	}
}

func initDb() {
	persistence.InitializeDb(os.Getenv("DB_DRIVER"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	seeder.Seed()
}
