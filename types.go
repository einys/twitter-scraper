package twitterscraper

import "time"

type (
	// Mention type.
	Mention struct {
		ID       string `bson:"id,omitempty" json:"id,omitempty"`
		Username string `bson:"username,omitempty" json:"username,omitempty"`
		Name     string `bson:"name,omitempty" json:"name,omitempty"`
	}

	// Url represents a URL with display, expanded, and index data.
	Url struct {
		DisplayURL  string `bson:"display_url,omitempty" json:"display_url,omitempty"`
		ExpandedURL string `bson:"expanded_url,omitempty" json:"expanded_url,omitempty"`
		URL         string `bson:"url,omitempty" json:"url,omitempty"`
		Indices     []int  `bson:"indices,omitempty" json:"indices,omitempty"`
	}

	// Photo type.
	Photo struct {
		ID  string `bson:"id,omitempty" json:"id,omitempty"`
		URL string `bson:"url,omitempty" json:"url,omitempty"`
	}

	// Video type.
	Video struct {
		ID      string `bson:"id,omitempty" json:"id,omitempty"`
		Preview string `bson:"preview,omitempty" json:"preview,omitempty"`
		URL     string `bson:"url,omitempty" json:"url,omitempty"`
		HLSURL  string `bson:"hls_url,omitempty" json:"hls_url,omitempty"`
	}

	// GIF type.
	GIF struct {
		ID      string `bson:"id,omitempty" json:"id,omitempty"`
		Preview string `bson:"preview,omitempty" json:"preview,omitempty"`
		URL     string `bson:"url,omitempty" json:"url,omitempty"`
	}

	// Tweet type.
	Tweet struct {
		Version        int      `bson:"version,omitempty" json:"version,omitempty"`
		ConversationID string   `bson:"conversation_id,omitempty" json:"conversation_id,omitempty"`
		GIFs           []GIF    `bson:"gifs,omitempty" json:"gifs,omitempty"`
		Hashtags       []string `bson:"hashtags,omitempty" json:"hashtags,omitempty"`
		HTML           string   `bson:"html,omitempty" json:"html,omitempty"`

		// ID는 핵심 식별자이므로 omitempty 제거
		ID string `bson:"id" json:"id"`

		InReplyToStatus   *Tweet `bson:"in_reply_to_status,omitempty" json:"in_reply_to_status,omitempty"`
		InReplyToStatusID string `bson:"in_reply_to_status_id,omitempty" json:"in_reply_to_status_id,omitempty"`

		IsQuoted     bool `bson:"is_quoted,omitempty" json:"is_quoted,omitempty"`
		IsPin        bool `bson:"is_pin,omitempty" json:"is_pin,omitempty"`
		IsReply      bool `bson:"is_reply,omitempty" json:"is_reply,omitempty"`
		IsRetweet    bool `bson:"is_retweet,omitempty" json:"is_retweet,omitempty"`
		IsSelfThread bool `bson:"is_self_thread,omitempty" json:"is_self_thread,omitempty"`

		Likes int    `bson:"likes,omitempty" json:"likes,omitempty"`
		Name  string `bson:"name,omitempty" json:"name,omitempty"`

		Mentions     []Mention `bson:"mentions,omitempty" json:"mentions,omitempty"`
		PermanentURL string    `bson:"permanent_url,omitempty" json:"permanent_url,omitempty"`
		Photos       []Photo   `bson:"photos,omitempty" json:"photos,omitempty"`

		Place *Place `bson:"place,omitempty" json:"place,omitempty"`

		QuotedStatus   *Tweet `bson:"quoted_status,omitempty" json:"quoted_status,omitempty"`
		QuotedStatusID string `bson:"quoted_status_id,omitempty" json:"quoted_status_id,omitempty"`

		Replies  int `bson:"replies,omitempty" json:"replies,omitempty"`
		Retweets int `bson:"retweets,omitempty" json:"retweets,omitempty"`

		RetweetedStatus   *Tweet `bson:"retweeted_status,omitempty" json:"retweeted_status,omitempty"`
		RetweetedStatusID string `bson:"retweeted_status_id,omitempty" json:"retweeted_status_id,omitempty"`

		Text       string    `bson:"text,omitempty" json:"text,omitempty"`
		Thread     []*Tweet  `bson:"thread,omitempty" json:"thread,omitempty"`
		TimeParsed time.Time `bson:"time_parsed,omitempty" json:"time_parsed,omitempty"`
		Timestamp  int64     `bson:"timestamp,omitempty" json:"timestamp,omitempty"`
		URLs       []string  `bson:"urls,omitempty" json:"urls,omitempty"`

		UserID   string `bson:"user_id,omitempty" json:"user_id,omitempty"`
		Username string `bson:"username,omitempty" json:"username,omitempty"`
		Avatar   string `bson:"avatar,omitempty" json:"avatar,omitempty"`

		Videos []Video `bson:"videos,omitempty" json:"videos,omitempty"`
		Views  int     `bson:"views,omitempty" json:"views,omitempty"`

		SensitiveContent bool `bson:"sensitive_content,omitempty" json:"sensitive_content,omitempty"`
	}

	// ProfileResult of scrapping.
	ProfileResult struct {
		Profile `bson:",inline,omitempty" json:",inline,omitempty"`
		Error   error `bson:"error,omitempty" json:"error,omitempty"`
	}

	// ScrappedTweetResult of scrapping.
	ScrappedTweetResult struct {
		Tweet `bson:",inline,omitempty" json:",inline,omitempty"`
		Error error `bson:"error,omitempty" json:"error,omitempty"`
	}

	ScheduledTweet struct {
		ID        string    `bson:"id,omitempty" json:"id,omitempty"`
		State     string    `bson:"state,omitempty" json:"state,omitempty"`
		ExecuteAt time.Time `bson:"execute_at,omitempty" json:"execute_at,omitempty"`
		Text      string    `bson:"text,omitempty" json:"text,omitempty"`
		Videos    []Video   `bson:"videos,omitempty" json:"videos,omitempty"`
		Photos    []Photo   `bson:"photos,omitempty" json:"photos,omitempty"`
		GIFs      []GIF     `bson:"gifs,omitempty" json:"gifs,omitempty"`
	}

	ExtendedMedia struct {
		IDStr                    string `bson:"id_str,omitempty" json:"id_str,omitempty"`
		MediaURLHttps            string `bson:"media_url_https,omitempty" json:"media_url_https,omitempty"`
		ExtSensitiveMediaWarning struct {
			AdultContent    bool `bson:"adult_content,omitempty" json:"adult_content,omitempty"`
			GraphicViolence bool `bson:"graphic_violence,omitempty" json:"graphic_violence,omitempty"`
			Other           bool `bson:"other,omitempty" json:"other,omitempty"`
		} `bson:"ext_sensitive_media_warning,omitempty" json:"ext_sensitive_media_warning,omitempty"`
		Type      string `bson:"type,omitempty" json:"type,omitempty"`
		URL       string `bson:"url,omitempty" json:"url,omitempty"`
		VideoInfo struct {
			Variants []struct {
				Type    string `bson:"content_type,omitempty" json:"content_type,omitempty"`
				Bitrate int    `bson:"bitrate,omitempty" json:"bitrate,omitempty"`
				URL     string `bson:"url,omitempty" json:"url,omitempty"`
			} `bson:"variants,omitempty" json:"variants,omitempty"`
		} `bson:"video_info,omitempty" json:"video_info,omitempty"`
	}
	Professional struct {
		RestID           string `bson:"rest_id,omitempty" json:"rest_id,omitempty"`
		ProfessionalType string `bson:"professional_type,omitempty" json:"professional_type,omitempty"`
		// interface{} 대신 구체적인 구조체 슬라이스로 변경
		Category []struct {
			ID       int    `bson:"id,omitempty" json:"id,omitempty"`
			Name     string `bson:"name,omitempty" json:"name,omitempty"`
			IconName string `bson:"icon_name,omitempty" json:"icon_name,omitempty"`
		} `bson:"category,omitempty" json:"category,omitempty"`
	}
	// UserV2 represents a Twitter user profile.
	UserV2 struct {
		ID                       string `bson:"id,omitempty" json:"id,omitempty"`
		RestID                   string `bson:"rest_id,omitempty" json:"rest_id,omitempty"`
		AffiliatesHighlightLabel struct {
			Label struct {
				URL struct {
					URL     string `bson:"url,omitempty" json:"url,omitempty"`
					URLType string `bson:"urlType,omitempty" json:"urlType,omitempty"`
				} `bson:"url,omitempty" json:"url,omitempty"`
				Badge struct {
					URL string `bson:"url,omitempty" json:"url,omitempty"`
				} `bson:"badge,omitempty" json:"badge,omitempty"`
				Description          string `bson:"description,omitempty" json:"description,omitempty"`
				UserLabelType        string `bson:"userLabelType,omitempty" json:"userLabelType,omitempty"`
				UserLabelDisplayType string `bson:"userLabelDisplayType,omitempty" json:"userLabelDisplayType,omitempty"`
			} `bson:"label,omitempty" json:"label,omitempty"`
		} `bson:"affiliates_highlighted_label,omitempty" json:"affiliates_highlighted_label,omitempty"`
		Avatar struct {
			ImageURL string `bson:"image_url,omitempty" json:"image_url,omitempty"`
		} `bson:"avatar,omitempty" json:"avatar,omitempty"`
		Core struct {
			CreatedAt  string `bson:"created_at,omitempty" json:"created_at,omitempty"`
			Name       string `bson:"name,omitempty" json:"name,omitempty"`
			ScreenName string `bson:"screen_name,omitempty" json:"screen_name,omitempty"`
		} `bson:"core,omitempty" json:"core,omitempty"`
		DmPermissions struct {
			CanDM bool `bson:"can_dm,omitempty" json:"can_dm,omitempty"`
		} `bson:"dm_permissions,omitempty" json:"dm_permissions,omitempty"`
		FollowedBy         bool         `bson:"followed_by,omitempty" json:"followed_by,omitempty"`
		Following          bool         `bson:"following,omitempty" json:"following,omitempty"`
		HasGraduatedAccess bool         `bson:"has_graduated_access,omitempty" json:"has_graduated_access,omitempty"`
		IsBlueVerified     bool         `bson:"is_blue_verified,omitempty" json:"is_blue_verified,omitempty"`
		Legacy             legacyUserV2 `bson:"legacy,omitempty" json:"legacy,omitempty"`
		Location           struct {
			Location string `bson:"location,omitempty" json:"location,omitempty"`
		} `bson:"location,omitempty" json:"location,omitempty"`
		MediaPermissions struct {
			CanMediaTag bool `bson:"can_media_tag,omitempty" json:"can_media_tag,omitempty"`
		} `bson:"media_permissions,omitempty" json:"media_permissions,omitempty"`
		ParodyCommentaryFanLabel string       `bson:"parody_commentary_fan_label,omitempty" json:"parody_commentary_fan_label,omitempty"`
		ProfileImageShape        string       `bson:"profile_image_shape,omitempty" json:"profile_image_shape,omitempty"`
		Professional             Professional `bson:"professional,omitempty" json:"professional,omitempty"`
		ProfileBio               struct {
			Description string `bson:"description,omitempty" json:"description,omitempty"`
		} `bson:"profile_bio,omitempty" json:"profile_bio,omitempty"`
		Privacy struct {
			Protected bool `bson:"protected,omitempty" json:"protected,omitempty"`
		} `bson:"privacy,omitempty" json:"privacy,omitempty"`
		RelationshipPerspectives struct {
			Following bool `bson:"following,omitempty" json:"following,omitempty"`
		} `bson:"relationship_perspectives,omitempty" json:"relationship_perspectives,omitempty"`
		TipjarSettings struct {
			IsEnabled bool `bson:"is_enabled,omitempty" json:"is_enabled,omitempty"`
		} `bson:"tipjar_settings,omitempty" json:"tipjar_settings,omitempty"`
		SuperFollowEligible bool `bson:"super_follow_eligible,omitempty" json:"super_follow_eligible,omitempty"`
		Verification        struct {
			Verified bool `bson:"verified,omitempty" json:"verified,omitempty"`
		} `bson:"verification,omitempty" json:"verification,omitempty"`
	}

	// legacyUserV2 represents the legacy user data structure.
	legacyUserV2 struct {
		CreatedAt   string `bson:"created_at,omitempty" json:"created_at,omitempty"`
		Description string `bson:"description,omitempty" json:"description,omitempty"`
		Entities    struct {
			Description struct {
				URLs []Url `bson:"urls,omitempty" json:"urls,omitempty"`
			} `bson:"description,omitempty" json:"description,omitempty"`
			URL struct {
				URLs []Url `bson:"urls,omitempty" json:"urls,omitempty"`
			} `bson:"url,omitempty" json:"url,omitempty"`
		} `bson:"entities,omitempty" json:"entities,omitempty"`
		FavouritesCount         int      `bson:"favourites_count,omitempty" json:"favourites_count,omitempty"`
		FollowersCount          int      `bson:"followers_count,omitempty" json:"followers_count,omitempty"`
		FriendsCount            int      `bson:"friends_count,omitempty" json:"friends_count,omitempty"`
		IDStr                   string   `bson:"id_str,omitempty" json:"id_str,omitempty"`
		ListedCount             int      `bson:"listed_count,omitempty" json:"listed_count,omitempty"`
		Name                    string   `bson:"name,omitempty" json:"name,omitempty"`
		Location                string   `bson:"location,omitempty" json:"location,omitempty"`
		PinnedTweetIdsStr       []string `bson:"pinned_tweet_ids_str,omitempty" json:"pinned_tweet_ids_str,omitempty"`
		ProfileBannerURL        string   `bson:"profile_banner_url,omitempty" json:"profile_banner_url,omitempty"`
		ProfileImageURLHTTPS    string   `bson:"profile_image_url_https,omitempty" json:"profile_image_url_https,omitempty"`
		Protected               bool     `bson:"protected,omitempty" json:"protected,omitempty"`
		ScreenName              string   `bson:"screen_name,omitempty" json:"screen_name,omitempty"`
		StatusesCount           int      `bson:"statuses_count,omitempty" json:"statuses_count,omitempty"`
		Verified                bool     `bson:"verified,omitempty" json:"verified,omitempty"`
		FollowedBy              bool     `bson:"followed_by,omitempty" json:"followed_by,omitempty"`
		Following               bool     `bson:"following,omitempty" json:"following,omitempty"`
		CanDm                   bool     `bson:"can_dm,omitempty" json:"can_dm,omitempty"`
		CanMediaTag             bool     `bson:"can_media_tag,omitempty" json:"can_media_tag,omitempty"`
		DefaultProfile          bool     `bson:"default_profile,omitempty" json:"default_profile,omitempty"`
		DefaultProfileImage     bool     `bson:"default_profile_image,omitempty" json:"default_profile_image,omitempty"`
		FastFollowersCount      int      `bson:"fast_followers_count,omitempty" json:"fast_followers_count,omitempty"`
		HasCustomTimelines      bool     `bson:"has_custom_timelines,omitempty" json:"has_custom_timelines,omitempty"`
		IsTranslator            bool     `bson:"is_translator,omitempty" json:"is_translator,omitempty"`
		MediaCount              int      `bson:"media_count,omitempty" json:"media_count,omitempty"`
		NeedsPhoneVerification  bool     `bson:"needs_phone_verification,omitempty" json:"needs_phone_verification,omitempty"`
		NormalFollowersCount    int      `bson:"normal_followers_count,omitempty" json:"normal_followers_count,omitempty"`
		PossiblySensitive       bool     `bson:"possibly_sensitive,omitempty" json:"possibly_sensitive,omitempty"`
		ProfileInterstitialType string   `bson:"profile_interstitial_type,omitempty" json:"profile_interstitial_type,omitempty"`
		TranslatorType          string   `bson:"translator_type,omitempty" json:"translator_type,omitempty"`
		WantRetweets            bool     `bson:"want_retweets,omitempty" json:"want_retweets,omitempty"`
		WithheldInCountries     []string `bson:"withheld_in_countries,omitempty" json:"withheld_in_countries,omitempty"`
	}

	legacyTweet struct {
		ConversationIDStr string `bson:"conversation_id_str,omitempty" json:"conversation_id_str,omitempty"`
		CreatedAt         string `bson:"created_at,omitempty" json:"created_at,omitempty"`
		FavoriteCount     int    `bson:"favorite_count,omitempty" json:"favorite_count,omitempty"`
		FullText          string `bson:"full_text,omitempty" json:"full_text,omitempty"`
		Entities          struct {
			Hashtags []struct {
				Text string `bson:"text,omitempty" json:"text,omitempty"`
			} `bson:"hashtags,omitempty" json:"hashtags,omitempty"`
			Media []struct {
				MediaURLHttps string `bson:"media_url_https,omitempty" json:"media_url_https,omitempty"`
				Type          string `bson:"type,omitempty" json:"type,omitempty"`
				URL           string `bson:"url,omitempty" json:"url,omitempty"`
			} `bson:"media,omitempty" json:"media,omitempty"`
			URLs         []Url `bson:"urls,omitempty" json:"urls,omitempty"`
			UserMentions []struct {
				IDStr      string `bson:"id_str,omitempty" json:"id_str,omitempty"`
				Name       string `bson:"name,omitempty" json:"name,omitempty"`
				ScreenName string `bson:"screen_name,omitempty" json:"screen_name,omitempty"`
			} `bson:"user_mentions,omitempty" json:"user_mentions,omitempty"`
		} `bson:"entities,omitempty" json:"entities,omitempty"`
		ExtendedEntities struct {
			Media []ExtendedMedia `bson:"media,omitempty" json:"media,omitempty"`
		} `bson:"extended_entities,omitempty" json:"extended_entities,omitempty"`
		IDStr                 string `bson:"id_str,omitempty" json:"id_str,omitempty"`
		InReplyToStatusIDStr  string `bson:"in_reply_to_status_id_str,omitempty" json:"in_reply_to_status_id_str,omitempty"`
		Place                 Place  `bson:"place,omitempty" json:"place,omitempty"`
		ReplyCount            int    `bson:"reply_count,omitempty" json:"reply_count,omitempty"`
		RetweetCount          int    `bson:"retweet_count,omitempty" json:"retweet_count,omitempty"`
		RetweetedStatusIDStr  string `bson:"retweeted_status_id_str,omitempty" json:"retweeted_status_id_str,omitempty"`
		RetweetedStatusResult struct {
			Result *result `bson:"result,omitempty" json:"result,omitempty"`
		} `bson:"retweeted_status_result,omitempty" json:"retweeted_status_result,omitempty"`
		QuotedStatusIDStr string `bson:"quoted_status_id_str,omitempty" json:"quoted_status_id_str,omitempty"`
		SelfThread        struct {
			IDStr string `bson:"id_str,omitempty" json:"id_str,omitempty"`
		} `bson:"self_thread,omitempty" json:"self_thread,omitempty"`
		Time      time.Time `bson:"time,omitempty" json:"time,omitempty"`
		UserIDStr string    `bson:"user_id_str,omitempty" json:"user_id_str,omitempty"`
		Views     struct {
			State string `bson:"state,omitempty" json:"state,omitempty"`
			Count string `bson:"count,omitempty" json:"count,omitempty"`
		} `bson:"ext_views,omitempty" json:"ext_views,omitempty"`
	}

	Place struct {
		ID          string `bson:"id,omitempty" json:"id,omitempty"`
		PlaceType   string `bson:"place_type,omitempty" json:"place_type,omitempty"`
		Name        string `bson:"name,omitempty" json:"name,omitempty"`
		FullName    string `bson:"full_name,omitempty" json:"full_name,omitempty"`
		CountryCode string `bson:"country_code,omitempty" json:"country_code,omitempty"`
		Country     string `bson:"country,omitempty" json:"country,omitempty"`
		BoundingBox struct {
			Type        string        `bson:"type,omitempty" json:"type,omitempty"`
			Coordinates [][][]float64 `bson:"coordinates,omitempty" json:"coordinates,omitempty"`
		} `bson:"bounding_box,omitempty" json:"bounding_box,omitempty"`
	}

	fetchProfileFunc func(query string, maxProfilesNbr int, cursor string) ([]*Profile, string, error)
	fetchTweetFunc   func(query string, maxTweetsNbr int, cursor string) ([]*Tweet, string, error)

	legacyExtendedProfile struct {
		Birthdate struct {
			Day            int    `bson:"day,omitempty" json:"day,omitempty"`
			Month          int    `bson:"month,omitempty" json:"month,omitempty"`
			Year           int    `bson:"year,omitempty" json:"year,omitempty"`
			Visibility     string `bson:"visibility,omitempty" json:"visibility,omitempty"`
			YearVisibility string `bson:"year_visibility,omitempty" json:"year_visibility,omitempty"`
		} `bson:"birthdate,omitempty" json:"birthdate,omitempty"`
	}

	verificationInfo struct {
		IsIdentityVerified bool `bson:"is_identity_verified,omitempty" json:"is_identity_verified,omitempty"`
		Reason             struct {
			Description struct {
				Text     string `bson:"text,omitempty" json:"text,omitempty"`
				Entities []struct {
					FromIndex int `bson:"from_index,omitempty" json:"from_index,omitempty"`
					ToIndex   int `bson:"to_index,omitempty" json:"to_index,omitempty"`
					Ref       struct {
						URL     string `bson:"url,omitempty" json:"url,omitempty"`
						URLType string `bson:"url_type,omitempty" json:"url_type,omitempty"`
					} `bson:"ref,omitempty" json:"ref,omitempty"`
				} `bson:"entities,omitempty" json:"entities,omitempty"`
			} `bson:"description,omitempty" json:"description,omitempty"`
			VerifiedSinceMsec string `bson:"verified_since_msec,omitempty" json:"verified_since_msec,omitempty"`
		} `bson:"reason,omitempty" json:"reason,omitempty"`
	}

	highlightsInfo struct {
		CanHighlightTweets bool   `bson:"can_highlight_tweets,omitempty" json:"can_highlight_tweets,omitempty"`
		HighlightedTweets  string `bson:"highlighted_tweets,omitempty" json:"highlighted_tweets,omitempty"`
	}
)
