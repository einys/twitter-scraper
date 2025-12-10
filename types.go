package twitterscraper

import "time"

type (
	// Mention type.
	Mention struct {
		ID       string
		Username string
		Name     string
	}

	// Url represents a URL with display, expanded, and index data.
	Url struct {
		DisplayURL  string `json:"display_url"`
		ExpandedURL string `json:"expanded_url"`
		URL         string `json:"url"`
		Indices     []int  `json:"indices"`
	}

	// Photo type.
	Photo struct {
		ID  string
		URL string
	}

	// Video type.
	Video struct {
		ID      string
		Preview string
		URL     string
		HLSURL  string
	}

	// GIF type.
	GIF struct {
		ID      string
		Preview string
		URL     string
	}

	// Tweet type.
	Tweet struct {
		ConversationID    string
		GIFs              []GIF
		Hashtags          []string
		HTML              string
		ID                string
		InReplyToStatus   *Tweet
		InReplyToStatusID string
		IsQuoted          bool
		IsPin             bool
		IsReply           bool
		IsRetweet         bool
		IsSelfThread      bool
		Likes             int
		Name              string
		Mentions          []Mention
		PermanentURL      string
		Photos            []Photo
		Place             *Place
		QuotedStatus      *Tweet
		QuotedStatusID    string
		Replies           int
		Retweets          int
		RetweetedStatus   *Tweet
		RetweetedStatusID string
		Text              string
		Thread            []*Tweet
		TimeParsed        time.Time
		Timestamp         int64
		URLs              []string
		UserID            string
		Username          string
		Videos            []Video
		Views             int
		SensitiveContent  bool
	}

	// ProfileResult of scrapping.
	ProfileResult struct {
		Profile
		Error error
	}

	// TweetResult of scrapping.
	TweetResult struct {
		Tweet
		Error error
	}

	ScheduledTweet struct {
		ID        string
		State     string
		ExecuteAt time.Time
		Text      string
		Videos    []Video
		Photos    []Photo
		GIFs      []GIF
	}

	ExtendedMedia struct {
		IDStr                    string `json:"id_str"`
		MediaURLHttps            string `json:"media_url_https"`
		ExtSensitiveMediaWarning struct {
			AdultContent    bool `json:"adult_content"`
			GraphicViolence bool `json:"graphic_violence"`
			Other           bool `json:"other"`
		} `json:"ext_sensitive_media_warning"`
		Type      string `json:"type"`
		URL       string `json:"url"`
		VideoInfo struct {
			Variants []struct {
				Type    string `json:"content_type"`
				Bitrate int    `json:"bitrate"`
				URL     string `json:"url"`
			} `json:"variants"`
		} `json:"video_info"`
	}

	// UserV2 represents a Twitter user profile.
	UserV2 struct {
		ID                       string `json:"id"`
		RestID                   string `json:"rest_id"`
		AffiliatesHighlightLabel struct {
			Label struct {
				URL struct {
					URL     string `json:"url"`
					URLType string `json:"urlType"`
				} `json:"url"`
				Badge struct {
					URL string `json:"url"`
				} `json:"badge"`
				Description          string `json:"description"`
				UserLabelType        string `json:"userLabelType"`
				UserLabelDisplayType string `json:"userLabelDisplayType"`
			} `json:"label"`
		} `json:"affiliates_highlighted_label"`
		Avatar struct {
			ImageURL string `json:"image_url"`
		} `json:"avatar"`
		Core struct {
			CreatedAt  string `json:"created_at"`
			Name       string `json:"name"`
			ScreenName string `json:"screen_name"`
		} `json:"core"`
		DmPermissions struct {
			CanDM bool `json:"can_dm"`
		} `json:"dm_permissions"`
		FollowedBy         bool         `json:"followed_by"`
		Following          bool         `json:"following"`
		HasGraduatedAccess bool         `json:"has_graduated_access"`
		IsBlueVerified     bool         `json:"is_blue_verified"`
		Legacy             legacyUserV2 `json:"legacy"`
		Location           struct {
			Location string `json:"location"`
		} `json:"location"`
		MediaPermissions struct {
			CanMediaTag bool `json:"can_media_tag"`
		} `json:"media_permissions"`
		ParodyCommentaryFanLabel string `json:"parody_commentary_fan_label"`
		ProfileImageShape        string `json:"profile_image_shape"`
		Professional             struct {
			RestID           string   `json:"rest_id"`
			ProfessionalType string   `json:"professional_type"`
			Category         []string `json:"category"`
		} `json:"professional"`
		ProfileBio struct {
			Description string `json:"description"`
		} `json:"profile_bio"`
		Privacy struct {
			Protected bool `json:"protected"`
		} `json:"privacy"`
		RelationshipPerspectives struct {
			Following bool `json:"following"`
		} `json:"relationship_perspectives"`
		TipjarSettings struct {
			IsEnabled bool `json:"is_enabled"`
		} `json:"tipjar_settings"`
		SuperFollowEligible bool `json:"super_follow_eligible"`
		Verification        struct {
			Verified bool `json:"verified"`
		} `json:"verification"`
	}

	// legacyUserV2 represents the legacy user data structure.
	legacyUserV2 struct {
		CreatedAt   string `json:"created_at"`
		Description string `json:"description"`
		Entities    struct {
			Description struct {
				URLs []Url `json:"urls"`
			} `json:"description"`
			URL struct {
				URLs []Url `json:"urls"`
			} `json:"url"`
		} `json:"entities"`
		FavouritesCount         int      `json:"favourites_count"`
		FollowersCount          int      `json:"followers_count"`
		FriendsCount            int      `json:"friends_count"`
		IDStr                   string   `json:"id_str"`
		ListedCount             int      `json:"listed_count"`
		Name                    string   `json:"name"`
		Location                string   `json:"location"`
		PinnedTweetIdsStr       []string `json:"pinned_tweet_ids_str"`
		ProfileBannerURL        string   `json:"profile_banner_url"`
		ProfileImageURLHTTPS    string   `json:"profile_image_url_https"`
		Protected               bool     `json:"protected"`
		ScreenName              string   `json:"screen_name"`
		StatusesCount           int      `json:"statuses_count"`
		Verified                bool     `json:"verified"`
		FollowedBy              bool     `json:"followed_by"`
		Following               bool     `json:"following"`
		CanDm                   bool     `json:"can_dm"`
		CanMediaTag             bool     `json:"can_media_tag"`
		DefaultProfile          bool     `json:"default_profile"`
		DefaultProfileImage     bool     `json:"default_profile_image"`
		FastFollowersCount      int      `json:"fast_followers_count"`
		HasCustomTimelines      bool     `json:"has_custom_timelines"`
		IsTranslator            bool     `json:"is_translator"`
		MediaCount              int      `json:"media_count"`
		NeedsPhoneVerification  bool     `json:"needs_phone_verification"`
		NormalFollowersCount    int      `json:"normal_followers_count"`
		PossiblySensitive       bool     `json:"possibly_sensitive"`
		ProfileInterstitialType string   `json:"profile_interstitial_type"`
		TranslatorType          string   `json:"translator_type"`
		WantRetweets            bool     `json:"want_retweets"`
		WithheldInCountries     []string `json:"withheld_in_countries"`
	}

	legacyTweet struct {
		ConversationIDStr string `json:"conversation_id_str"`
		CreatedAt         string `json:"created_at"`
		FavoriteCount     int    `json:"favorite_count"`
		FullText          string `json:"full_text"`
		Entities          struct {
			Hashtags []struct {
				Text string `json:"text"`
			} `json:"hashtags"`
			Media []struct {
				MediaURLHttps string `json:"media_url_https"`
				Type          string `json:"type"`
				URL           string `json:"url"`
			} `json:"media"`
			URLs         []Url `json:"urls"`
			UserMentions []struct {
				IDStr      string `json:"id_str"`
				Name       string `json:"name"`
				ScreenName string `json:"screen_name"`
			} `json:"user_mentions"`
		} `json:"entities"`
		ExtendedEntities struct {
			Media []ExtendedMedia `json:"media"`
		} `json:"extended_entities"`
		IDStr                 string `json:"id_str"`
		InReplyToStatusIDStr  string `json:"in_reply_to_status_id_str"`
		Place                 Place  `json:"place"`
		ReplyCount            int    `json:"reply_count"`
		RetweetCount          int    `json:"retweet_count"`
		RetweetedStatusIDStr  string `json:"retweeted_status_id_str"`
		RetweetedStatusResult struct {
			Result *result `json:"result"`
		} `json:"retweeted_status_result"`
		QuotedStatusIDStr string `json:"quoted_status_id_str"`
		SelfThread        struct {
			IDStr string `json:"id_str"`
		} `json:"self_thread"`
		Time      time.Time `json:"time"`
		UserIDStr string    `json:"user_id_str"`
		Views     struct {
			State string `json:"state"`
			Count string `json:"count"`
		} `json:"ext_views"`
	}

	Place struct {
		ID          string `json:"id"`
		PlaceType   string `json:"place_type"`
		Name        string `json:"name"`
		FullName    string `json:"full_name"`
		CountryCode string `json:"country_code"`
		Country     string `json:"country"`
		BoundingBox struct {
			Type        string        `json:"type"`
			Coordinates [][][]float64 `json:"coordinates"`
		} `json:"bounding_box"`
	}

	fetchProfileFunc func(query string, maxProfilesNbr int, cursor string) ([]*Profile, string, error)
	fetchTweetFunc   func(query string, maxTweetsNbr int, cursor string) ([]*Tweet, string, error)

	legacyExtendedProfile struct {
		Birthdate struct {
			Day            int    `json:"day"`
			Month          int    `json:"month"`
			Year           int    `json:"year"`
			Visibility     string `json:"visibility"`
			YearVisibility string `json:"year_visibility"`
		} `json:"birthdate"`
	}

	verificationInfo struct {
		IsIdentityVerified bool `json:"is_identity_verified"`
		Reason             struct {
			Description struct {
				Text     string `json:"text"`
				Entities []struct {
					FromIndex int `json:"from_index"`
					ToIndex   int `json:"to_index"`
					Ref       struct {
						URL     string `json:"url"`
						URLType string `json:"url_type"`
					} `json:"ref"`
				} `json:"entities"`
			} `json:"description"`
			VerifiedSinceMsec string `json:"verified_since_msec"`
		} `json:"reason"`
	}

	highlightsInfo struct {
		CanHighlightTweets bool   `json:"can_highlight_tweets"`
		HighlightedTweets  string `json:"highlighted_tweets"`
	}
)
