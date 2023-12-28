package models

type Feature struct {
	Name string
	Link string
}

var (
	Features = []Feature{
		{
			Name: "Dashboard",
			Link: "/",
		},
		{
			Name: "User",
			Link: "/user",
		},
		{
			Name: "Other",
			Link: "/other",
		},
	}
)
