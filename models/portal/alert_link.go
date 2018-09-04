package portal

import "time"

/*
DESC alert_link;
+-----------+------------------+------+-----+---------------------+----------------+
| Field     | Type             | Null | Key | Default             | Extra          |
+-----------+------------------+------+-----+---------------------+----------------+
| id        | int(10) unsigned | NO   | PRI | NULL                | auto_increment |
| path      | varchar(16)      | NO   | UNI |                     |                |
| content   | text             | NO   |     | NULL                |                |
| create_at | timestamp        | NO   |     | current_timestamp() |                |
+-----------+------------------+------+-----+---------------------+----------------+
*/
type AlertLink struct {
	Id      uint      `xorm:"id notnull int pk autoincr"`
	Path    string    `xorm:"path notnull varchar(16)"`
	Content string    `xorm:"content notnull text"`
	Created time.Time `xorm:"create_at notnull datetime created"`
}
