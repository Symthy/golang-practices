// Package client provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.8.2 DO NOT EDIT.
package client

// Error response
type Error struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

// Items defines model for Items.
type Items []struct {
	Body          string  `json:"body"`
	Coediting     bool    `json:"coediting"`
	CommentsCount float32 `json:"comments_count"`
	CreatedAt     string  `json:"created_at"`
	Group         struct {
		CreatedAt string  `json:"created_at"`
		Id        float32 `json:"id"`
		Name      string  `json:"name"`
		Private   bool    `json:"private"`
		UpdatedAt string  `json:"updated_at"`
		UrlName   string  `json:"url_name"`
	} `json:"group"`
	Id             string  `json:"id"`
	LikesCount     float32 `json:"likes_count"`
	PageViewsCount float32 `json:"page_views_count"`
	Private        bool    `json:"private"`
	ReactionsCount float32 `json:"reactions_count"`
	RenderedBody   string  `json:"rendered_body"`
	Tags           []struct {
		Name     string         `json:"name"`
		Versions *[]interface{} `json:"versions,omitempty"`
	} `json:"tags"`
	TeamMembership struct {
		Name string `json:"name"`
	} `json:"team_membership"`
	Title     string `json:"title"`
	UpdatedAt string `json:"updated_at"`
	Url       string `json:"url"`
	User      struct {
		Description       string  `json:"description"`
		FacebookId        string  `json:"facebook_id"`
		FolloweesCount    float32 `json:"followees_count"`
		FollowersCount    float32 `json:"followers_count"`
		GithubLoginName   string  `json:"github_login_name"`
		Id                string  `json:"id"`
		ItemsCount        float32 `json:"items_count"`
		LinkedinId        string  `json:"linkedin_id"`
		Location          string  `json:"location"`
		Name              string  `json:"name"`
		Organization      string  `json:"organization"`
		PermanentId       float32 `json:"permanent_id"`
		ProfileImageUrl   string  `json:"profile_image_url"`
		TeamOnly          bool    `json:"team_only"`
		TwitterScreenName string  `json:"twitter_screen_name"`
		WebsiteUrl        string  `json:"website_url"`
	} `json:"user"`
}

// User defines model for User.
type User struct {
	Description                 string  `json:"description"`
	FacebookId                  string  `json:"facebook_id"`
	FolloweesCount              float32 `json:"followees_count"`
	FollowersCount              float32 `json:"followers_count"`
	GithubLoginName             string  `json:"github_login_name"`
	Id                          string  `json:"id"`
	ImageMonthlyUploadLimit     float32 `json:"image_monthly_upload_limit"`
	ImageMonthlyUploadRemaining float32 `json:"image_monthly_upload_remaining"`
	ItemsCount                  float32 `json:"items_count"`
	LinkedinId                  string  `json:"linkedin_id"`
	Location                    string  `json:"location"`
	Name                        string  `json:"name"`
	Organization                string  `json:"organization"`
	PermanentId                 float32 `json:"permanent_id"`
	ProfileImageUrl             string  `json:"profile_image_url"`
	TeamOnly                    bool    `json:"team_only"`
	TwitterScreenName           string  `json:"twitter_screen_name"`
	WebsiteUrl                  string  `json:"website_url"`
}

// Error response
type N400 Error

// Error response
type N401 Error

// Error response
type N403 Error

// Error response
type N404 Error

// Error response
type N500 Error

// GetAuthenticatedUserParams defines parameters for GetAuthenticatedUser.
type GetAuthenticatedUserParams struct {
	Authorization *string `json:"Authorization,omitempty"`
}

// GetUserItemsParams defines parameters for GetUserItems.
type GetUserItemsParams struct {
	Page          *string `json:"page,omitempty"`
	PerPage       *string `json:"per_page,omitempty"`
	Authorization *string `json:"Authorization,omitempty"`
}

