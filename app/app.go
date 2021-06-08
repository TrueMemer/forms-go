package app

import (
	middleware "forms/app/middleware"
	"forms/app/routes"
	"github.com/go-chi/chi/v5"
	defaultMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
	"os"
	"time"
)

type App struct {
	Router chi.Router
	Database Database
	tokenAuth *jwtauth.JWTAuth
}

func New() *App {
	a := &App{
		Router: chi.NewRouter(),
		Database: Database{},
		tokenAuth: jwtauth.New("HS256", []byte(os.Getenv("ACCESS_SECRET")), nil),
	}

	a.initDatabase()
	a.initMiddleware()
	a.initRoutes()
	return a
}

func (app *App) initMiddleware() {
	app.Router.Use(render.SetContentType(render.ContentTypeJSON))
	app.Router.Use(defaultMiddleware.Logger)
	app.Router.Use(defaultMiddleware.RequestID)
	app.Router.Use(defaultMiddleware.RealIP)
	app.Router.Use(defaultMiddleware.Logger)
	app.Router.Use(defaultMiddleware.Recoverer)

	app.Router.Use(middleware.Database(app.Database.GetDatabase()))

	app.Router.Use(defaultMiddleware.Timeout(60 * time.Second))
}

func (app *App) initDatabase() {
	app.Database.InitDatabase()
}

func (app *App) initRoutes() {
	app.Router.Mount("/", routes.RootRouter())
	app.Router.Mount("/auth", routes.AuthRouter())

	// Protected routes
	app.Router.Group(func (r chi.Router) {
		r.Use(jwtauth.Verifier(app.tokenAuth))

		r.Use(jwtauth.Authenticator)

		r.Use(middleware.UserCtx)

		r.Mount("/profile", routes.ProfileRouter())
		r.Mount("/users", routes.UsersRouter())
		r.Mount("/groups", routes.GroupsRouter())
	})
}