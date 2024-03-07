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
