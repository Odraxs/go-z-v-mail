package email_search_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/Odraxs/go-z-v-mail/server/config"
	"github.com/Odraxs/go-z-v-mail/server/email_search"
	"github.com/Odraxs/go-z-v-mail/server/test/fixtures"
)

func TestGetEmails(t *testing.T) {
	config.LoadZincsearchCredentials()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("creating server")
		if r.URL.Path != "/emails/_search" {
			t.Errorf("Expected to request '/emails/_search', got: %s", r.URL.Path)
		}

		w.WriteHeader(http.StatusOK)
		w.Write(fixtures.Get_Emails_Response())
	}))
	defer server.Close()

	test := []struct {
		name     string
		filter   email_search.SearchEmailRequest
		builder  email_search.ZincsearchRepo
		expected email_search.EmailSearchResponse
		err      error
		context  context.Context
	}{{
		name:    "get emails",
		context: context.Background(),
		filter: email_search.SearchEmailRequest{
			Term:       "California",
			Field:      "content",
			SortFields: []string{},
			MaxResults: 2,
		},
		builder: email_search.ZincsearchRepo{
			HttpClient:         server.Client(),
			ZincsearchEndpoint: server.URL,
		},
		expected: email_search.EmailSearchResponse{
			Time: 123,
			Emails: []email_search.Email{
				{
					Id:      "26yn6kHS7H9",
					From:    "email@enron.com",
					To:      "email2@enron.com, email3@enron.com",
					Content: "Guys, good job with the California issues - we didn't get everything but we managed the position in a very effective manner.",
					Subject: "california",
					Date:    "2001-11-26T01:23:00-08:00",
					Highlight: email_search.HighlightSearchResult{
						Subject: []string{"<strong>california</strong>"},
					},
				},
			},
		},
	}}

	for _, test := range test {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.builder.GetEmails(tt.context, tt.filter)
			if err != tt.err {
				t.Fatalf("expecting error %v but got %v", tt.err, err)
			}

			if !reflect.DeepEqual(result, tt.expected) {
				t.Fatalf("Expected %v but got %v", tt.expected, result)
			}
		})
	}
}
