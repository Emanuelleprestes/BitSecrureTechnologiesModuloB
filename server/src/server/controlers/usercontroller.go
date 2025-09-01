package controlers

import "database/sql"

type Usercontroller struct {
	conn *sql.Conn
}

func Newusercontroller(c *sql.Conn) *Usercontroller {
	return &Usercontroller{conn: c}
}
