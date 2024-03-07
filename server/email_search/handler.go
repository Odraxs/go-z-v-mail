package email_search

import (
	"context"
	"fmt"
	"net/http"
)

type Repo interface {
	GetEmails(ctx context.Context, filter string)
}

type EmailSearchHandler struct {
	Repo Repo
}

func (es *EmailSearchHandler) Search(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Initializing search")
}
