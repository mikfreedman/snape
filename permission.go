package snape

type Permissions []Permission

type Permission struct {
	ID           string
	EmailAddress string
	Role         string
	FileID       string
}
