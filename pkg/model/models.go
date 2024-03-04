package model

type LoginUserInput struct {
	AccentColor      float64 `json:"accent_color"`
	Avatar           string  `json:"avatar"`
	AvatarDecoration string  `json:"avatar_decoratio_data"`
	Banner           string  `json:"banner"`
	Discriminator    string  `json:"discriminator"`
	GlobalName       string  `json:"global_name"`
	Locale           string  `json:"locale"`
	Username         string  `json:"username"`
}

type NewLinkInput struct {
	Title  string `json:"title"`
	Url    string `json:"url"`
	Image  string `json:"image"`
	Author string `json:"author"`
}
