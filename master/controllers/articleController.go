package controllers

import (
	"blog-article-go/master/models"
	"blog-article-go/master/usecases"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type articleHandler struct {
	articleUsecase usecases.ArticleUsecase
}

var (
	article         models.Article
	articleResponse models.Response
)

func getMuxArticle(value string, r *http.Request) string {
	parameter := mux.Vars(r)
	return parameter[value]
}

func ArticleController(r *mux.Router, service usecases.ArticleUsecase) {
	articleHandler := articleHandler{service}
	r.HandleFunc("/articles/{creator}", articleHandler.ListArticle).Methods(http.MethodGet)
	r.HandleFunc("/article/{creator}/{articleid}", articleHandler.ListArticleById).Methods(http.MethodGet)
	r.HandleFunc("/article/{creator}", articleHandler.AddArticlePage).Methods(http.MethodPost)
	r.HandleFunc("/article/{creator}/{articleid}", articleHandler.UpdateArticlePage).Methods(http.MethodPut)
	r.HandleFunc("/article/{creator}/{articleid}", articleHandler.DeleteArticlePage).Methods(http.MethodDelete)
}

func (a *articleHandler) ListArticle(w http.ResponseWriter, r *http.Request) {
	creator := getMuxArticle("creator", r)
	articles, err := a.articleUsecase.GetAllArticle(creator)

	if err != nil {
		w.Write([]byte("Data Not Found"))
	} else {
		articleResponse.Messages = "List of Article Data"
		articleResponse.Data = articles
		articleResponse.Status = http.StatusOK
		byteOfArticle, err := json.Marshal(articleResponse)
		if err != nil {
			w.Write([]byte("Something went wrong!"))
		}
		w.Header().Set("content-type", "application/json")
		w.Write(byteOfArticle)
	}
}

func (a *articleHandler) ListArticleById(w http.ResponseWriter, r *http.Request) {
	creator := getMuxArticle("creator", r)
	idArticle := getMuxArticle("articleid", r)
	article, err := a.articleUsecase.GetArticleByID(creator, idArticle)
	if err != nil {
		w.Write([]byte("Data Not Found"))
	} else {
		articleResponse.Messages = "List of Article Data"
		articleResponse.Data = article
		articleResponse.Status = http.StatusOK
		byteOfArticle, err := json.Marshal(articleResponse)
		if err != nil {
			w.Write([]byte("Something went wrong!"))
		}
		w.Header().Set("content-type", "application/json")
		w.Write(byteOfArticle)
	}
}

//AddArticlePage is a controller of AddAArticle service (Insert / POST). It got the json data from Request Body RAW
func (a *articleHandler) AddArticlePage(w http.ResponseWriter, r *http.Request) {
	_ = json.NewDecoder(r.Body).Decode(&article)
	creator := getMuxArticle("creator", r)
	err := a.articleUsecase.AddArticle(article.ArticleName, article.ArticleContent, creator)
	if err != nil {
		w.Write([]byte("Data can not be empty!"))
	} else {
		articleResponse.Messages = "Insert Article Data Success!"
		articleResponse.Status = http.StatusOK
		articleResponse.Data = ""
		byteOfArticle, err := json.Marshal(articleResponse)
		if err != nil {
			w.Write([]byte("Something went wrong!"))
		}
		w.Header().Set("content-type", "application/json")
		w.Write(byteOfArticle)
	}
}

//UpdateArticlePage is a controller of UpdateAArticle service (Update/PUT). It got the json data from Request Body RAW
func (a *articleHandler) UpdateArticlePage(w http.ResponseWriter, r *http.Request) {
	_ = json.NewDecoder(r.Body).Decode(&article)
	creator := getMuxArticle("creator", r)
	idArticle := getMuxArticle("articleid", r)
	err := a.articleUsecase.UpdateArticle(idArticle, article.ArticleName, article.ArticleContent, creator)

	if err != nil {
		w.Write([]byte("Something went wrong !"))
	} else {
		articleResponse.Messages = "Update Article Success"
		articleResponse.Status = http.StatusOK
		articleResponse.Data = ""
		byteOfArticle, err := json.Marshal(articleResponse)
		if err != nil {
			w.Write([]byte("Something went wrong!"))
		}
		w.Header().Set("content-type", "application/json")
		w.Write(byteOfArticle)
	}
}

// DeleteArticlePage is a controller of  DeleteAArticle service  (DeleTE). It got the json data from Params
func (a *articleHandler) DeleteArticlePage(w http.ResponseWriter, r *http.Request) {
	creator := getMuxArticle("creator", r)
	idArticle := getMuxArticle("articleid", r)
	err := a.articleUsecase.DeleteArticle(creator, idArticle)
	if err != nil {
		w.Write([]byte("Data Not Found"))
	} else {
		articleResponse.Messages = "Delete Article Success"
		articleResponse.Status = http.StatusOK
		articleResponse.Data = ""
		byteOfArticle, err := json.Marshal(articleResponse)
		if err != nil {
			w.Write([]byte("Something went wrong!"))
		}
		w.Header().Set("content-type", "application/json")
		w.Write(byteOfArticle)
	}
}
