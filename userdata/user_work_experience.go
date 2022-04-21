package userdata

type WorkExperience struct {
	CompanyName string
	JobTitle    string
	StartDate   string
	EndDate     string
	Projects    []Project
}

type Project struct {
	Name        string
	ProjectDesc []string
}
