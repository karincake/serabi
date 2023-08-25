package cryptographyregex

// Just-side-effect package for it's registering its tag regex list

import s "github.com/karincake/serabi"

func init() {
	s.AddTagForRegex("md4", "^[0-9a-f]{32}$", "must be a valid md4 format")
	s.AddTagForRegex("md5", "^[0-9a-f]{32}$", "must be a valid md5 format")
	s.AddTagForRegex("sha256", "^[0-9a-f]{64}$", "must be a valid sha256 format")
	s.AddTagForRegex("sha384", "^[0-9a-f]{96}$", "must be a valid sha384 format")
	s.AddTagForRegex("sha512", "^[0-9a-f]{128}$", "must be a valid sha512 format")
	s.AddTagForRegex("ripemd128", "^[0-9a-f]{32}$", "must be a valid ripemd128 format")
	s.AddTagForRegex("ripemd160", "^[0-9a-f]{40}$", "must be a valid ripemd160 format")
	s.AddTagForRegex("tiger128", "^[0-9a-f]{32}$", "must be a valid tiger128 format")
	s.AddTagForRegex("tiger160", "^[0-9a-f]{40}$", "must be a valid tiger160 format")
	s.AddTagForRegex("tiger192", "^[0-9a-f]{48}$", "must be a valid tiger192 format")
}
