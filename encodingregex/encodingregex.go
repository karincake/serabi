package encodingregex

import s "github.com/karincake/serabi"

func init() {
	s.AddTagForRegex("base64", "^(?:[A-Za-z0-9+\\/]{4})*(?:[A-Za-z0-9+\\/]{2}==|[A-Za-z0-9+\\/]{3}=|[A-Za-z0-9+\\/]{4})$", "must be a valid base64 encoded string")
	s.AddTagForRegex("base64URL", "^(?:[A-Za-z0-9-_]{4})*(?:[A-Za-z0-9-_]{2}==|[A-Za-z0-9-_]{3}=|[A-Za-z0-9-_]{4})$", "must be a valid base64 URL encoded string")
	s.AddTagForRegex("base64RawURL", "^(?:[A-Za-z0-9-_]{4})*(?:[A-Za-z0-9-_]{2,4})$", "must be a valid base64 Raw URL encoded string")
	s.AddTagForRegex("url", "`^(?:[^%]|%[0-9A-Fa-f]{2})*$`", "must be a valid URL encoded string")
	s.AddTagForRegex("html", "&#[x]?([0-9a-fA-F]{2})|(&gt)|(&lt)|(&quot)|(&amp)+[;]?", "must be a valid HTML encoded")
}
