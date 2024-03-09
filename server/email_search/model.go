package email_search

type EmailSearchResponse struct {
	Time   int     `json:"time"`
	Emails []Email `json:"emails"`
}

type Email struct {
	Id        string   `json:"id"`
	From      string   `json:"from"`
	To        string   `json:"to"`
	Content   string   `json:"content"`
	Subject   string   `json:"subject"`
	Date      string   `json:"date"`
	Highlight []string `json:"highlight"`
}

type EmailSearchResult struct {
	Took     int  `json:"took"`
	TimedOut bool `json:"timed_out"`
	Hits     struct {
		Total struct {
			Value int `json:"value"`
		} `json:"total"`
		Hits []struct {
			Index     string  `json:"_index"`
			Type      string  `json:"_type"`
			ID        string  `json:"_id"`
			Score     float64 `json:"_score"`
			Timestamp string  `json:"@timestamp"`
			Source    Email   `json:"_source"`
			Highlight struct {
				Content []string `json:"content"`
			} `json:"highlight"`
		} `json:"hits"`
	} `json:"hits"`
}

type SearchDocumentsBody struct {
	SearchType string                   `json:"search_type"`
	SortFields []string                 `json:"sort_fields"`
	From       int                      `json:"from"`
	MaxResults int                      `json:"max_results"`
	Query      SearchDocumentsQuery `json:"query"`
	Source     []string                 `json:"_source"`
	Highlight  Highlight                `json:"highlight"`
}

type SearchDocumentsQuery struct {
	Term      string  `json:"term"`
	Field     string  `json:"field"`
	StartTime *string `json:"start_time,omitempty"`
	EndTime   *string `json:"end_time,omitempty"`
}

type Highlight struct {
	PreTags  []string                  `json:"pre_tags"`
	PostTags []string                  `json:"post_tags"`
	Fields   *map[string]interface{} `json:"fields,omitempty"`
}
