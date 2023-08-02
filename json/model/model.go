package model

// parse social link object
type SocialLink struct {
	Site string `json:"site"`
	URL  string `json:"url"`
}

// parse user object
type User struct {
	Firstname       string       `json:"firstname"`
	Lastname        string       `json:"lastname"`
	ID              string       `json:"id"`
	Age             string       `json:"age"`
	Gender          string       `json:"gender"`
	PreferredTopics []string     `json:"preferred_topics"`
	SocialLinks     []SocialLink `json:"social_links"`
}
