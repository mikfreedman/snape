package snape

import "context"

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate . API
type API interface {
	GetPermissions(ctx context.Context, folderID string, recursive bool, debug bool) (Permissions, error)
}
