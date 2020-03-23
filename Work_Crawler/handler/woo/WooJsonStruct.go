package woo

//AutoGenerated ...
type AutoGenerated []struct {
	ID            int           `json:"id"`
	Date          string        `json:"date"`
	DateGmt       string        `json:"date_gmt"`
	GUID          GUID          `json:"guid"`
	Modified      string        `json:"modified"`
	ModifiedGmt   string        `json:"modified_gmt"`
	Slug          string        `json:"slug"`
	Status        string        `json:"status"`
	Type          string        `json:"type"`
	Link          string        `json:"link"`
	Title         Title         `json:"title"`
	Content       Content       `json:"content"`
	Excerpt       Excerpt       `json:"excerpt"`
	FeaturedMedia int           `json:"featured_media"`
	CommentStatus string        `json:"comment_status"`
	PingStatus    string        `json:"ping_status"`
	Template      string        `json:"template"`
	Meta          []interface{} `json:"meta"`
	Links         Links         `json:"_links"`
}

//GUID ...
type GUID struct {
	Rendered string `json:"rendered"`
}

//Title ...
type Title struct {
	Rendered string `json:"rendered"`
}

//Content ...
type Content struct {
	Rendered  string `json:"rendered"`
	Protected bool   `json:"protected"`
}

//Excerpt ...
type Excerpt struct {
	Rendered  string `json:"rendered"`
	Protected bool   `json:"protected"`
}

//Self ...
type Self struct {
	Href string `json:"href"`
}

//Collection ...
type Collection struct {
	Href string `json:"href"`
}

//About ...
type About struct {
	Href string `json:"href"`
}

//Replies ...
type Replies struct {
	Embeddable bool   `json:"embeddable"`
	Href       string `json:"href"`
}

//WpFeaturedmedia ...
type WpFeaturedmedia struct {
	Embeddable bool   `json:"embeddable"`
	Href       string `json:"href"`
}

//WpAttachment ...
type WpAttachment struct {
	Href string `json:"href"`
}

//Curies ...
type Curies struct {
	Name      string `json:"name"`
	Href      string `json:"href"`
	Templated bool   `json:"templated"`
}

//Links ...
type Links struct {
	Self            []Self            `json:"self"`
	Collection      []Collection      `json:"collection"`
	About           []About           `json:"about"`
	Replies         []Replies         `json:"replies"`
	WpFeaturedmedia []WpFeaturedmedia `json:"wp:featuredmedia"`
	WpAttachment    []WpAttachment    `json:"wp:attachment"`
	Curies          []Curies          `json:"curies"`
}