package portal

import "time"

/*
DESC variable;
+-------------+------------------+------+-----+---------------------+----------------+
| Field       | Type             | Null | Key | Default             | Extra          |
+-------------+------------------+------+-----+---------------------+----------------+
| id          | int(10) unsigned | NO   | PRI | NULL                | auto_increment |
| grp_id      | int(10) unsigned | NO   | MUL | NULL                |                |
| name        | varchar(255)     | NO   |     | NULL                |                |
| content     | varchar(1024)    | NO   |     | NULL                |                |
| note        | varchar(1024)    | NO   |     | NULL                |                |
| create_user | varchar(64)      | NO   |     |                     |                |
| create_at   | timestamp        | NO   |     | current_timestamp() |                |
+-------------+------------------+------+-----+---------------------+----------------+
*/
type Variable struct {
	Id      uint      `xorm:"'id' notnull int pk autoincr"`
	GrpId   uint      `xorm:"'grp_id' notnull int"`
	Name    string    `xorm:"'name' notnull varchar(255)"`
	Content string    `xorm:"'content' notnull varchar(1024)"`
	Note    string    `xorm:"'note' notnull varchar(1024)"`
	Creator string    `xorm:"'create_user' notnull varchar(64)"`
	Created time.Time `xorm:"'create_at' notnull datetime created"`
}

func (s *Variable) TableName() string {
	return "variable"
}
