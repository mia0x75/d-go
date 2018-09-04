package portal

import "time"

/*
DESC tag;
+-----------+------------------+------+-----+---------------------+----------------+
| Field     | Type             | Null | Key | Default             | Extra          |
+-----------+------------------+------+-----+---------------------+----------------+
| id        | int(10) unsigned | NO   | PRI | NULL                | auto_increment |
| name      | varchar(255)     | NO   | UNI | NULL                |                |
| update_at | timestamp        | NO   |     | current_timestamp() |                |
+-----------+------------------+------+-----+---------------------+----------------+
*/
type Tag struct {
	Id      uint      `xorm:"id notnull int pk autoincr"`
	Name    string    `xorm:"name notnull varchar(255)"`
	Updated time.Time `xorm:"update_at notnull datetime updated"`
}
