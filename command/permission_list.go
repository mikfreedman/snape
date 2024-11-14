package command

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/mikfreedman/snape"
)

type PermissionList struct {
	FolderID string `arg:"" name:"folder_id" help:"Folder ID to retrieve permissions for"`
}

func (p *PermissionList) Run(ctx Context, api snape.API, w io.Writer) error {

	permissionList, err := api.GetPermissions(context.Background(), p.FolderID)

	if err != nil {
		log.Printf("Unable to process permission list for folder ID '%s'", p.FolderID)
		return err
	}

	for _, permission := range permissionList {
		fmt.Printf("%s, %s, %s\n", permission.FileID, permission.EmailAddress, permission.ID)
	}

	return nil
}
