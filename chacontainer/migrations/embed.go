// Package migrations embeds the CHACONTAINER SQL schema files so the server
// binary can apply them on startup without any external migration tool -
// required to make `docker compose up` self-sufficient on a local machine.
package migrations

import "embed"

//go:embed *.sql
var FS embed.FS
