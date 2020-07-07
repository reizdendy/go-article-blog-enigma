package master

import (
	"blog-article-go/master/controllers"
	"blog-article-go/master/repositories"
	"blog-article-go/master/usecases"
	"database/sql"

	"github.com/gorilla/mux"
)

func Init(db *sql.DB, r *mux.Router) {
	//user routers
	userRepo := repositories.InitUserRepositoryImpl(db)
	userUsecase := usecases.InitUserUsecase(userRepo)
	controllers.UserController(r, userUsecase)

	//article routers
	articleRepo := repositories.InitArticleRepositoryImpl(db)
	articleUsecase := usecases.InitArticleUsecase(articleRepo)
	controllers.ArticleController(r, articleUsecase)

}
