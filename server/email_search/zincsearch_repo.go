package email_search

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Odraxs/go-z-v-mail/server/config"
)

const (
	zincsearchEndpoint = "http://zincsearch:4080/api"
	defaultSearchType  = "matchphrase"
)

type ZincsearchRepo struct {
	HttpClient         *http.Client
	ZincsearchEndpoint string
}

func NewZincsearchRepository(httpClient *http.Client) Repo {
	return &ZincsearchRepo{
		HttpClient:         httpClient,
		ZincsearchEndpoint: zincsearchEndpoint,
	}
}

// GetEmails implements Repo.
func (z *ZincsearchRepo) GetEmails(ctx context.Context, filter SearchEmailRequest) (EmailSearchResponse, error) {
	var searchType string

	if filter.Field == "from" {
		searchType = "prefix"
	} else {
		searchType = defaultSearchType
	}

	requestBody := SearchDocumentsBody{
		SearchType: searchType,
		Query: SearchDocumentsQuery{
			Term:  filter.Term,
			Field: filter.Field,
		},
		SortFields: filter.SortFields,
		From:       0,
		MaxResults: filter.MaxResults,
		Highlight: Highlight{
			PreTags:  []string{"<strong>"},
			PostTags: []string{"</strong>"},
			Fields: map[string]interface{}{
				filter.Field: map[string]interface{}{
					"pre_tags":  []string{},
					"post_tags": []string{},
				},
			},
		},
	}

	var response EmailSearchResponse

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return response, fmt.Errorf("failed to encode zincsearch request body: %w", err)
	}

	zRequest, err := http.NewRequest(http.MethodPost, z.ZincsearchEndpoint+"/emails/_search", bytes.NewBuffer(jsonBody))
	if err != nil {
		return response, fmt.Errorf("failed to create zincsearch request: %w", err)
	}

	zRequest = prepareRequest(zRequest)

	zResponse, err := z.HttpClient.Do(zRequest)
	if err != nil || zResponse.StatusCode != http.StatusOK {
		return response, fmt.Errorf("request to zincsearch failed: %w", err)
	}
	defer zResponse.Body.Close()

	zBody, err := io.ReadAll(zResponse.Body)
	if err != nil {
		return response, fmt.Errorf("failed to read the zincsearch response body: %w", err)
	}

	emailsSearchResult := EmailSearchResult{}

	if err := json.Unmarshal(zBody, &emailsSearchResult); err != nil {
		return EmailSearchResponse{}, fmt.Errorf("failed to decode zincsearch response body")
	}

	emails := convertToEmails(emailsSearchResult)
	response = EmailSearchResponse{
		Time:   emailsSearchResult.Hits.Total.Value,
		Emails: emails,
	}

	return response, nil
}

func prepareRequest(zRequest *http.Request) *http.Request {
	credentials := config.GetZincsearchCredentials()

	zRequest.Header.Set("Content-Type", "application/json")
	zRequest.SetBasicAuth(credentials.User, credentials.Password)

	return zRequest
}

func convertToEmails(response EmailSearchResult) []Email {
	emails := []Email{}

	for _, hit := range response.Hits.Hits {
		email := Email{
			Id:        hit.ID,
			From:      hit.Source.From,
			To:        hit.Source.To,
			Subject:   hit.Source.Subject,
			Content:   hit.Source.Content,
			Date:      hit.Source.Date,
			Highlight: hit.Highlight,
		}
		emails = append(emails, email)
	}

	return emails
}
