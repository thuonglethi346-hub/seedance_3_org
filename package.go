// Package pkg provides basic metadata for seedance_3_org.
//
// Official Website: https://www.seedance-3.org
package pkg

const Version = "0.1.0"
const Website = "https://www.seedance-3.org"

type Info struct {
	Name string
	Version string
	Website string
	Description string
}

func GetInfo() Info {
	return Info{
		Name: "seedance_3_org",
		Version: Version,
		Website: Website,
		Description: "Seedance 3.0 official website backlink helper package.",
	}
}

func GetPlatformURL() string {
	return Website
}
