package fixtures

func Get_Emails_Response() []byte {
	return []byte(`{
		"took": 23,
		"timed_out": false,
		"hits": {
			"total": {
				"value": 123
			},
			"hits": [
				{
					"_index": "emails",
					"_type": "_doc",
					"_id": "26yn6kHS7H9",
					"_score": 6.845917312925424,
					"@timestamp": "2024-03-11T23:06:22.588701952Z",
					"_source": {
						"content": "Guys, good job with the California issues - we didn't get everything but we managed the position in a very effective manner.",
						"date": "2001-11-26T01:23:00-08:00",
						"from": "email@enron.com",
						"subject": "california",
						"to": "email2@enron.com, email3@enron.com"
					},
					"highlight": {
						"subject": [
							"<strong>california</strong>"
						]
					}
				}
			]
		}
	}`)
}
