package userdata

type User struct {
	FirstName   string
	LastName    string
	Designation string
	Github      string
	Email       string
	Twitter     string
	Phone       string
	Work        []WorkExperience
	Education   []Education
}
