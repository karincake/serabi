package identifierregex

// Just-side-effect package for it's registering its tag regex list

import s "github.com/karincake/serabi"

func init() {
	s.AddTagForRegex("uuid", "^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$", "must be a valid base64 format")
	s.AddTagForRegex("uuid3", "^[0-9a-f]{8}-[0-9a-f]{4}-3[0-9a-f]{3}-[0-9a-f]{4}-[0-9a-f]{12}$", "must be a valid base64 format")
	s.AddTagForRegex("uuid4", "^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$", "must be a valid base64 format")
	s.AddTagForRegex("uuid5", "^[0-9a-f]{8}-[0-9a-f]{4}-5[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$", "must be a valid base64 format")
	s.AddTagForRegex("uuidRfc4122", "^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$", "must be a valid uuidRfc4122 format")
	s.AddTagForRegex("uuid3Rfc4122", "^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-3[0-9a-fA-F]{3}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$", "must be a valid uuid3Rfc4122 format")
	s.AddTagForRegex("uuid4Rfc4122", "^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$", "must be a valid uuid4Rfc4122 format")
	s.AddTagForRegex("uuid5Rfc4122", "^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-5[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$", "must be a valid uuid5Rfc4122 format")
}
