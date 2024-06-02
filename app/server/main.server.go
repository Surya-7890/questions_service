package server

import (
	"net/http"

	"github.com/Surya-7890/questions_service/app/database"
	"github.com/Surya-7890/questions_service/app/routers"
	"github.com/Surya-7890/questions_service/app/utils"
	dotenv "github.com/joho/godotenv"
)

type Server struct {
	Address string
}

func NewServer(address string) *Server {
	return &Server{
		Address: address,
	}
}

/*
 * for loading env variables and connecting to database
 */
func (s *Server) Init() error {
	err := dotenv.Load(".env.example")
	if err != nil {
		panic(err)
	}

	repo, err := database.NewRepository()
	if err != nil {
		return err
	}
	utils.Repo = repo
	repo.CreateTables()
	return nil
}

func (s *Server) StartServer() {

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	io.WriteString(w, "hello ")
	// })

	http.HandleFunc("/company/*", routers.HandleCompanyRoutes)

	if err := http.ListenAndServe(s.Address, nil); err != nil {
		panic(err)
	}
}
