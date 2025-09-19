package models

type VerificationRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type IsInDB struct {
	IsVer    bool `json:"is_ver"`
	IsInData bool `json:"is_in_data"`
}

type SignupJson struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type StatusR struct {
	Status bool `json:"status"`
}
