package email_search

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)

type Repo interface {
	GetEmails(ctx context.Context, filter string) (EmailSearchResponse, error)
}

type EmailSearchHandler struct {
	Repo Repo
}

func (es *EmailSearchHandler) Search(w http.ResponseWriter, r *http.Request) {
	log.Println("Initializing search")

	var body struct {
		Filter string `json:"filter"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response, err := es.Repo.GetEmails(r.Context(), body.Filter)
	if err != nil {
		log.Printf("failed to get a response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(response)
	if err != nil {
		log.Println("failed to encode response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(res)
	w.WriteHeader(http.StatusOK)
}
