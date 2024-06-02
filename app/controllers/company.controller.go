package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Surya-7890/questions_service/app/models"
	"github.com/Surya-7890/questions_service/app/utils"
)

type RequestBody struct {
	Name string `json:"company_name"`
}

func CreateCompany(w http.ResponseWriter, r *http.Request) {
	var body RequestBody
	var company models.Company

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if err := utils.Repo.DB.Where("name = ?", body.Name).First(&company).Error; err == nil {
		http.Error(w, "company already exists", http.StatusBadRequest)
	}

	new_company := models.Company{
		Name: body.Name,
	}
	if err := utils.Repo.DB.Create(&new_company).Error; err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&new_company)
}

func RemoveCompany(w http.ResponseWriter, r *http.Request) {
	var body RequestBody
	var company models.Company
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if err := utils.Repo.DB.Where("name = ?", body.Name).Delete(&company).Error; err != nil {
		fmt.Println(err)
		http.Error(w, "company with the given name not found", http.StatusNotFound)
	}

	var response struct {
		Message string `json:"message"`
	}
	response.Message = "successfully deleted"

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&response)
}
