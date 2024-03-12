package email_search_test

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/Odraxs/go-z-v-mail/server/config"
	"github.com/Odraxs/go-z-v-mail/server/email_search"
	"github.com/Odraxs/go-z-v-mail/server/test/fixtures"
)

func TestNewZincsearchRepository(t *testing.T) {
	httpClient := &http.Client{}
	test := struct {
		name     string
		expected *email_search.ZincsearchRepo
	}{
		name: "create new zincsearch repository",
		expected: &email_search.ZincsearchRepo{
			HttpClient:         httpClient,
			ZincsearchEndpoint: "http://zincsearch:4080/api",
		},
	}

	t.Run(test.name, func(t *testing.T) {
		repository := email_search.NewZincsearchRepository(httpClient)

		if !reflect.DeepEqual(repository, test.expected) {
			t.Fatalf("Expected %v but got %v", test.expected, repository)
		}
	})
}

func TestGetEmails(t *testing.T) {
	config.LoadZincsearchCredentials()
	basicSearchServer := fixtures.CreateEmailSearchMockServer(t, fixtures.GetBasicEmailsResponse())
	defer basicSearchServer.Close()
	filedFromSearchServer := fixtures.CreateEmailSearchMockServer(t, fixtures.GetFieldFromEmailsResponse())
	defer basicSearchServer.Close()

	test := []struct {
		name     string
		filter   email_search.SearchEmailRequest
		builder  email_search.ZincsearchRepo
		expected email_search.EmailSearchResponse
		err      error
		context  context.Context
	}{
		{
			name:    "get emails",
			context: context.Background(),
			filter: email_search.SearchEmailRequest{
				Term:       "California",
				Field:      "content",
				SortFields: []string{},
				MaxResults: 2,
			},
			builder: email_search.ZincsearchRepo{
				HttpClient:         basicSearchServer.Client(),
				ZincsearchEndpoint: basicSearchServer.URL,
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
		},
		{
			name:    "get emails with field from",
			context: context.Background(),
			filter: email_search.SearchEmailRequest{
				Term:       "email",
				Field:      "from",
				SortFields: []string{},
				MaxResults: 2,
			},
			builder: email_search.ZincsearchRepo{
				HttpClient:         filedFromSearchServer.Client(),
				ZincsearchEndpoint: filedFromSearchServer.URL,
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
							From: []string{"<strong>email@enron.com</strong>"},
						},
					},
				},
			},
		},
	}

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
