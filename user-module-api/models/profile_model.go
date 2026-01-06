package models

type Profile struct {
	ID         string `json:"id,omitempty"`
	UserID     string `json:"user_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Gender     string `json:"gender"`
	Profession string `json:"profession"`
	ProfilePic string `json:"profile_pic"`
	DOB        string `json:"dob"`
}
