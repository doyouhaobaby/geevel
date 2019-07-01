package postgres

import (
	"database/sql"
)

func init() {
	sql.Register("postgress", new(PostgressDrivier))
}