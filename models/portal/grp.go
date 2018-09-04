package portal

import "time"

/*
DESC grp;
+-------------+------------------+------+-----+---------------------+----------------+
| Field       | Type             | Null | Key | Default             | Extra          |
+-------------+------------------+------+-----+---------------------+----------------+
| id          | int(10) unsigned | NO   | PRI | NULL                | auto_increment |
| grp_name    | varchar(255)     | NO   | UNI |                     |                |
| create_user | varchar(64)      | NO   |     |                     |                |
| create_at   | timestamp        | NO   |     | current_timestamp() |                |
| come_from   | tinyint(4)       | NO   |     | 0                   |                |
+-------------+------------------+------+-----+---------------------+----------------+
*/
type Grp struct {
	Id      uint      `xorm:"'id' notnull int pk autoincr"`
	Name    string    `xorm:"'grp_name' notnull varchar(255) unique"`
	Creator string    `xorm:"'create_user' notnull varchar(64)"`
	Created time.Time `xorm:"'create_at' notnull datetime created"`
	From    int8      `xorm:"'come_from' notnull int default 0"`
}
