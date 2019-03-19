package main
import (
      //  "fmt"
	"log"
	"net/http"
	"encoding/json"
     //	"math/rand"
     //	"strconv"
	"github.com/gorilla/mux"
)

type ConfigurationManagement struct {
   ID              string  `json:"id"`
   Tool            string  `json:"tool"`
   LanguageBased   string  `json:"languagebased"`
   Company *Company `json:"company"`
}
// Author Struct
type Company struct {
  Name           string `json:"name"`
  Opensource     string `json:"opensource"`
  Certification  string `json:"certification"`
}

var configurationmanagement []ConfigurationManagement

// Get All CM Tools

func getcms(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(configurationmanagement)
}
// Get  Single CM Tool
func getcm(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r) // Get Params
  for _, item := range  configurationmanagement {
	  if item.ID == params["id"] {
		  json.NewEncoder(w).Encode(item)
		  return
	  }
  }
}






func main() {
        // Init Router
	r := mux.NewRouter()

        configurationmanagement = append(configurationmanagement, ConfigurationManagement{ID: "1", Tool: "Ansible", LanguageBased: "Python", Company: &Company{Name: "RedHat",Opensource: "AWX", Certification: "RedHat Specialist Ansible Automation"}})
	configurationmanagement = append(configurationmanagement, ConfigurationManagement{ID: "2", Tool: "Puppet", LanguageBased: "Ruby", Company: &Company{Name: "None",Opensource: "YES", Certification: "RedHat Specialist Configuration Management"}})

        // Route Handlers / Endpoints
	r.HandleFunc("/api/cms", getcms).Methods("GET")
	r.HandleFunc("/api/cm/{id}", getcm).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))

}
