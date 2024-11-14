package gooogle

import (
	"context"
	"fmt"
	"log"

	"github.com/mikfreedman/snape"
	"google.golang.org/api/drive/v3"
)

const FOLDER_MIME_TYPE = "application/vnd.google-apps.folder"

func (a *API) eachFile(folderID string, recursive bool, action func(*drive.File)) {
	query := fmt.Sprintf("'%s' in parents", folderID)
	pageToken := ""
	for {

		q := a.Client.Files.List().Q(query).PageSize(100)
		if pageToken != "" {
			q = q.PageToken(pageToken)
		}

		r, err := q.Do()
		if err != nil {
			log.Fatalf("Unable to retrieve files: %v", err)
		}

		for _, file := range r.Files {
			if recursive && file.MimeType == FOLDER_MIME_TYPE {
				a.eachFile(file.Id, recursive, action)
			}

			action(file)
		}

		pageToken = r.NextPageToken
		if pageToken == "" {
			break

		}

	}
}

func (a *API) GetPermissions(ctx context.Context, folderID string, recursive bool, debug bool) (snape.Permissions, error) {
	var permissions snape.Permissions

	a.eachFile(folderID, recursive, func(file *drive.File) {
		ppageToken := ""

		for {

			qq := a.Client.Permissions.List(file.Id).
				SupportsAllDrives(true).
				Fields("*").
				PageSize(100)

			if ppageToken != "" {
				qq = qq.PageToken(ppageToken)
			}

			pl, err := qq.Do()

			if err != nil {
				log.Fatalf("Unable to retrieve permission for file %s", file.Id)
			}

			for _, permission := range pl.Permissions {
				if debug {
					fmt.Printf("retrieved permission %s for file %s for user %s\n", permission.Id, file.Id, permission.EmailAddress)
				}

				permissions = append(permissions, snape.Permission{ID: file.Id, Filename: file.Name, EmailAddress: permission.EmailAddress, Role: permission.Role, FileID: permission.Id})
			}

			ppageToken = pl.NextPageToken

			if ppageToken == "" {
				break
			}
		}
	})

	return permissions, nil
}
