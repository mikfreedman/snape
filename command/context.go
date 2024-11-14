package command

type Context struct {
	ClientID string `default:"" hidden:"" env:"SNAPE_CLIENT_ID"`
}
