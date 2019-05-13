package mira

import (
	"encoding/json"
	"testing"
)

func TestGetId(t *testing.T) {
	sub := Subreddit{}
	json.Unmarshal([]byte(orig), &sub)
	if v := sub.GetId(); v != `t5_m0je4` {
		t.Error(
			"For GetId()",
			"expected", `t5_m0je4`,
			"got", v,
		)
	}
}

func TestGetName(t *testing.T) {
	sub := Subreddit{}
	json.Unmarshal([]byte(orig), &sub)
	if v := sub.GetName(); v != `MemeInvestor_bot` {
		t.Error(
			"For GetName()",
			"expected", `MemeInvestor_bot`,
			"got", v,
		)
	}
}

func TestGetDisplayName(t *testing.T) {
	sub := Subreddit{}
	json.Unmarshal([]byte(orig), &sub)
	if v := sub.GetDisplayName(); v != `MemeInvestor_bot` {
		t.Error(
			"For GetDisplayName()",
			"expected", `MemeInvestor_bot`,
			"got", v,
		)
	}
}

func TestGetUrl(t *testing.T) {
	sub := Subreddit{}
	json.Unmarshal([]byte(orig), &sub)
	if v := sub.GetUrl(); v != `/r/MemeInvestor_bot/` {
		t.Error(
			"For GetUrl()",
			"expected", "/r/MemeInvestor_bot/",
			"got", v,
		)
	}
}

func TestIsOver18(t *testing.T) {
	sub := Subreddit{}
	json.Unmarshal([]byte(orig), &sub)
	if v := sub.IsOver18(); v != false {
		t.Error(
			"For IsOver18()",
			"expected", false,
			"got", v,
		)
	}
}

func TestGetPublicDescription(t *testing.T) {
	sub := Subreddit{}
	json.Unmarshal([]byte(orig), &sub)
	if v := sub.GetPublicDescription(); v != "This subreddit is for questions, reports, or suggestions regarding /u/MemeInvestor_Bot. \n\nFor quick information see https://memes.market" {
		t.Error(
			"For GetPublicDescription()",
			"expected", "This subreddit is for questions, reports, or suggestions regarding /u/MemeInvestor_Bot. \n\nFor quick information see https://memes.market",
			"got", v,
		)
	}
}

func TestGetDescription(t *testing.T) {
	sub := Subreddit{}
	json.Unmarshal([]byte(orig), &sub)
	if v := sub.GetDescription(); v != "######This is the official subreddit for the bot, /u/MemeInvestor_bot\n\n***\n\nHere you're encouraged to **report bugs, ask questions, and submit suggestions** regarding the bot. This subreddit is frequently viewed by the developers and, whether or not you receive a reply, it's very likely that your submission has been viewed and noted by someone on the team.\n\n***\n\n#####Rules:\n1. This is a no-meme subreddit. Only serious suggestions, reports, or questions allowed.\n\n2.  All content must be regarding the bot. Keep it on-topic please.\n\n3. Be respectful. We're all nice people here.\n\n***\n\n&amp;nbsp;\n\n####^(Please don't send a message before first submitting your post on the subreddit.)\n\n######**[Message us anyway.](https://www.reddit.com/message/compose?to=%2Fr%2FMemeInvestor_Bot)**" {
		t.Error(
			"For GetDescription()",
			"expected", "######This is the official subreddit for the bot, /u/MemeInvestor_bot\n\n***\n\nHere you're encouraged to **report bugs, ask questions, and submit suggestions** regarding the bot. This subreddit is frequently viewed by the developers and, whether or not you receive a reply, it's very likely that your submission has been viewed and noted by someone on the team.\n\n***\n\n#####Rules:\n1. This is a no-meme subreddit. Only serious suggestions, reports, or questions allowed.\n\n2.  All content must be regarding the bot. Keep it on-topic please.\n\n3. Be respectful. We're all nice people here.\n\n***\n\n&amp;nbsp;\n\n####^(Please don't send a message before first submitting your post on the subreddit.)\n\n######**[Message us anyway.](https://www.reddit.com/message/compose?to=%2Fr%2FMemeInvestor_Bot)**",
			"got", v,
		)
	}
}

func TestGetCreated(t *testing.T) {
	sub := Subreddit{}
	json.Unmarshal([]byte(orig), &sub)
	if v := sub.GetCreated(); v != 1532014741.0 {
		t.Error(
			"For GetCreated()",
			"expected", 1532014741.0,
			"got", v,
		)
	}
}

func TestGetSubscribers(t *testing.T) {
	sub := Subreddit{}
	json.Unmarshal([]byte(orig), &sub)
	if v := sub.GetSubscribers(); v != 1339 {
		t.Error(
			"For GetSubscribers()",
			"expected", 1339,
			"got", v,
		)
	}
}

const (
	orig = `{"kind": "t5","data": {"user_flair_background_color": null,"submit_text_html": null,"restrict_posting": true,"user_is_banned": false,"free_form_reports": true,"wiki_enabled": null,"user_is_muted": false,"user_can_flair_in_sr": null,"display_name": "MemeInvestor_bot","header_img": null,"title": "MemeInvestor_bot","icon_size": [256,256],"primary_color": "","active_user_count": 2,"icon_img": "https:\/\/b.thumbs.redditmedia.com\/bz7E9y0Ca7UZS7QBCesV2aRavkZ9ZydYBYdsqRzdY1A.png","display_name_prefixed": "r\/MemeInvestor_bot","accounts_active": 2,"public_traffic": false,"subscribers": 1339,"user_flair_richtext": [],"videostream_links_count": 0,"name": "t5_m0je4","quarantine": false,"hide_ads": false,"emojis_enabled": false,"advertiser_category": "","public_description": "This subreddit is for questions, reports, or suggestions regarding \/u\/MemeInvestor_Bot. \n\nFor quick information see https:\/\/memes.market","comment_score_hide_mins": 0,"user_has_favorited": false,"user_flair_template_id": null,"community_icon": "","banner_background_image": "https:\/\/styles.redditmedia.com\/t5_m0je4\/styles\/bannerBackgroundImage_nin8va4paya11.png","original_content_tag_enabled": false,"submit_text": "","description_html": "&lt;!-- SC_OFF --&gt;&lt;div class=\"md\"&gt;&lt;h6&gt;This is the official subreddit for the bot, &lt;a href=\"\/u\/MemeInvestor_bot\"&gt;\/u\/MemeInvestor_bot&lt;\/a&gt;&lt;\/h6&gt;\n\n&lt;hr\/&gt;\n\n&lt;p&gt;Here you&amp;#39;re encouraged to &lt;strong&gt;report bugs, ask questions, and submit suggestions&lt;\/strong&gt; regarding the bot. This subreddit is frequently viewed by the developers and, whether or not you receive a reply, it&amp;#39;s very likely that your submission has been viewed and noted by someone on the team.&lt;\/p&gt;\n\n&lt;hr\/&gt;\n\n&lt;h5&gt;Rules:&lt;\/h5&gt;\n\n&lt;ol&gt;\n&lt;li&gt;&lt;p&gt;This is a no-meme subreddit. Only serious suggestions, reports, or questions allowed.&lt;\/p&gt;&lt;\/li&gt;\n&lt;li&gt;&lt;p&gt;All content must be regarding the bot. Keep it on-topic please.&lt;\/p&gt;&lt;\/li&gt;\n&lt;li&gt;&lt;p&gt;Be respectful. We&amp;#39;re all nice people here.&lt;\/p&gt;&lt;\/li&gt;\n&lt;\/ol&gt;\n\n&lt;hr\/&gt;\n\n&lt;p&gt;&amp;nbsp;&lt;\/p&gt;\n\n&lt;h4&gt;&lt;sup&gt;Please don&amp;#39;t send a message before first submitting your post on the subreddit.&lt;\/sup&gt;&lt;\/h4&gt;\n\n&lt;h6&gt;&lt;strong&gt;&lt;a href=\"https:\/\/www.reddit.com\/message\/compose?to=%2Fr%2FMemeInvestor_Bot\"&gt;Message us anyway.&lt;\/a&gt;&lt;\/strong&gt;&lt;\/h6&gt;\n&lt;\/div&gt;&lt;!-- SC_ON --&gt;","spoilers_enabled": true,"header_title": null,"header_size": null,"user_flair_position": "right","all_original_content": false,"has_menu_widget": false,"is_enrolled_in_new_modmail": null,"key_color": "#545452","can_assign_user_flair": false,"created": 1532043541.0,"wls": null,"show_media_preview": true,"submission_type": "any","user_is_subscriber": false,"disable_contributor_requests": false,"allow_videogifs": true,"user_flair_type": "text","collapse_deleted_comments": false,"emojis_custom_size": null,"public_description_html": "&lt;!-- SC_OFF --&gt;&lt;div class=\"md\"&gt;&lt;p&gt;This subreddit is for questions, reports, or suggestions regarding &lt;a href=\"\/u\/MemeInvestor_Bot\"&gt;\/u\/MemeInvestor_Bot&lt;\/a&gt;. &lt;\/p&gt;\n\n&lt;p&gt;For quick information see &lt;a href=\"https:\/\/memes.market\"&gt;https:\/\/memes.market&lt;\/a&gt;&lt;\/p&gt;\n&lt;\/div&gt;&lt;!-- SC_ON --&gt;","allow_videos": true,"notification_level": null,"can_assign_link_flair": true,"accounts_active_is_fuzzed": false,"submit_text_label": null,"link_flair_position": "right","user_sr_flair_enabled": true,"user_flair_enabled_in_sr": true,"allow_discovery": true,"user_sr_theme_enabled": true,"link_flair_enabled": true,"subreddit_type": "public","suggested_comment_sort": null,"banner_img": "https:\/\/b.thumbs.redditmedia.com\/Wqk7-zU-_lHm1-Oqd-Pg6fnShz2tKwNiFkV9Y23mwCM.png","user_flair_text": null,"banner_background_color": "","show_media": true,"id": "m0je4","user_is_moderator": false,"over18": false,"description": "######This is the official subreddit for the bot, \/u\/MemeInvestor_bot\n\n***\n\nHere you're encouraged to **report bugs, ask questions, and submit suggestions** regarding the bot. This subreddit is frequently viewed by the developers and, whether or not you receive a reply, it's very likely that your submission has been viewed and noted by someone on the team.\n\n***\n\n#####Rules:\n1. This is a no-meme subreddit. Only serious suggestions, reports, or questions allowed.\n\n2.  All content must be regarding the bot. Keep it on-topic please.\n\n3. Be respectful. We're all nice people here.\n\n***\n\n&amp;nbsp;\n\n####^(Please don't send a message before first submitting your post on the subreddit.)\n\n######**[Message us anyway.](https:\/\/www.reddit.com\/message\/compose?to=%2Fr%2FMemeInvestor_Bot)**","submit_link_label": null,"user_flair_text_color": null,"restrict_commenting": false,"user_flair_css_class": null,"allow_images": true,"lang": "en","whitelist_status": null,"url": "\/r\/MemeInvestor_bot\/","created_utc": 1532014741.0,"banner_size": [654,196],"mobile_banner_image": "","user_is_contributor": false}}`
)
