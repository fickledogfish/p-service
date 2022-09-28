//go:generate cp -r ../SQL/prepared ./prepared

package prepared_sql

import _ "embed"

type PreparedQuery string

var (
	//go:embed prepared/create_user.sql
	CreateUser PreparedQuery
)
