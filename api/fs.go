package api

import "embed"

//go:embed account/v1/account.openapi.yaml
var FS embed.FS
