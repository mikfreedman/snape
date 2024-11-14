package snape

type Permissions []Permission

type Permission struct {
	ID           string
	Filename     string
	EmailAddress string
	Role         string
	FileID       string
}
