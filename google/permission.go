package gooogle

import (
	"context"
	"fmt"
	"log"

	"github.com/mikfreedman/snape"
)

func (a *API) GetPermissions(ctx context.Context, folderID string) (snape.Permissions, error) {
	query := fmt.Sprintf("'%s' in parents", folderID)
	pageToken := ""
	var permissions snape.Permissions

	for {

		q := a.Client.Files.List().Q(query).PageSize(100)
		if pageToken != "" {
			q = q.PageToken(pageToken)
		}

		r, err := q.Do()
		if err != nil {
			log.Fatalf("Unable to retrieve files: %v", err)
		}

		if len(r.Files) == 0 {
			fmt.Println("No files found.")
		} else {
			for _, file := range r.Files {
				ppageToken := ""
				for {

					qq := a.Client.Permissions.List(file.Id).
						SupportsAllDrives(true).
						Fields("*").
						PageSize(100)

					if ppageToken != "" {
						q = q.PageToken(ppageToken)
					}

					pl, err := qq.Do()

					if err != nil {
						log.Fatalf("Unable to retrieve permission for file %s", file.Id)
					}

					for _, permission := range pl.Permissions {
						permissions = append(permissions, snape.Permission{ID: file.Id, EmailAddress: permission.EmailAddress, Role: permission.Role, FileID: permission.Id})
					}

					ppageToken = r.NextPageToken

					if ppageToken == "" {
						break
					}
				}
			}

			pageToken = r.NextPageToken
			if pageToken == "" {
				break
			}
		}

	}
	return permissions, nil
}
