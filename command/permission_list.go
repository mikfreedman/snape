package command

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/mikfreedman/snape"
)

type PermissionList struct {
	FolderID  string `arg:"" name:"folder_id" help:"Folder ID to retrieve permissions for"`
	Recursive bool   `name:"recursive" optional:"" default:"true" help:"Recursively evaluate folders"`
}

func (p *PermissionList) Run(ctx Context, api snape.API, w io.Writer) error {

	permissionList, err := api.GetPermissions(context.Background(), p.FolderID, p.Recursive, ctx.Debug)

	if err != nil {
		log.Printf("Unable to process permission list for folder ID '%s'", p.FolderID)
		return err
	}

	for _, permission := range permissionList {
		w.Write([]byte(fmt.Sprintf("%s, %s, %s, %s\n", permission.FileID, permission.Filename, permission.EmailAddress, permission.ID)))
	}

	return nil
}
