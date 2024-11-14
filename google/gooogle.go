package gooogle

import (
	"net/http"

	"google.golang.org/api/drive/v3"
)

type API struct {
	HttpClient *http.Client
	Client     *drive.Service
}
