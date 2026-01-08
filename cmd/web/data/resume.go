package data

type Resume struct {
	Basics   Basics         `json:"basics"`
	Sections ResumeSections `json:"sections"`
}

type Basics struct {
	Name     string `json:"name"`
	Headline string `json:"headline"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Location string `json:"location"`
	URL      struct {
		Label string `json:"label"`
		Href string `json:"href"`
	} `json:"url"`
	Fields []*CustomFields `json:"customFields"`
}

type CustomFields struct {
	ID string `json:"id"`
	ICON string `json:"icon"`
	NAME string `json:"name"`
	VALUE string `json:"value"`
}

type ResumeSections struct {
	Summary    SectionContent `json:"summary"`
	Experience SectionItems   `json:"experience"`
	Projects   SectionItems   `json:"projects"`
	Skills     SectionItems   `json:"skills"`
	Education  SectionItems   `json:"education"`
}

type SectionContent struct {
	Content string `json:"content"`
}

type SectionItems struct {
	Items []Item `json:"items"`
}

type Item struct {
	Name     string `json:"name"`
	Company  string `json:"company"`
	Position string `json:"position"`
	Date     string `json:"date"`
	Summary  string `json:"summary"`
	URL      *URL   `json:"url,omitempty"`
}

type URL struct {
	Label string `json:"label"`
	Href  string `json:"href"`
}

// type Theme struct {
//     Name string
//     Colors *ColorPalette
// }

// var Light, Dark *Theme