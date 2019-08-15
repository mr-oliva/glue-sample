package di

import (
	"net/http"

	"github.com/bookun/glue-sample/services/serviceA/controllers"
	"github.com/bookun/glue-sample/services/serviceA/handlers"
	"github.com/bookun/glue-sample/services/serviceA/repositories"
	"go.uber.org/fx"
)

//func StartServer(mux *http.ServeMux) error {
//	if err := http.ListenAndServe(":8080", mux); err != nil {
//		return err
//	}
//	return nil
//}

func NewRouter(userHandler *handlers.User) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/users", userHandler.HandleTaskGetUsers)
	mux.HandleFunc("/user/", userHandler.HandleTaskGetUserNameById)
	return mux
}

func NewString(path string) string {
	return path
}

func NewApp() *fx.App {
	app := fx.New(
		fx.Provide(
			func() string {
				return "testdata/user.csv"
			},
			repositories.NewFile,
			controllers.NewUser,
			handlers.NewUser,
			NewRouter,
		),
	)
	return app
}

var Module = fx.Options(
    fx.Option
	Config,
	Logger,
	Server,
)
