package routers

import (
	"net/http"

	"github.com/Surya-7890/questions_service/app/controllers"
)

func HandleCompanyRoutes(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/company/create":
		controllers.CreateCompany(w, r)
	case "/company/remove":
		controllers.RemoveCompany(w, r)
	}
}
