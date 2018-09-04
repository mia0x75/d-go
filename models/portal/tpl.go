package portal

import "time"

/*
DESC tpl;
+-------------+------------------+------+-----+---------------------+----------------+
| Field       | Type             | Null | Key | Default             | Extra          |
+-------------+------------------+------+-----+---------------------+----------------+
| id          | int(10) unsigned | NO   | PRI | NULL                | auto_increment |
| tpl_name    | varchar(255)     | NO   | UNI |                     |                |
| parent_id   | int(10) unsigned | NO   |     | 0                   |                |
| action_id   | int(10) unsigned | NO   |     | 0                   |                |
| create_user | varchar(64)      | NO   | MUL |                     |                |
| create_at   | timestamp        | NO   |     | current_timestamp() |                |
+-------------+------------------+------+-----+---------------------+----------------+
*/
type Tpl struct {
	Id       uint      `xorm:"id notnull int pk autoincr"`
	Name     string    `xorm:"tpl_name notnull varchar(255)"`
	ParentId uint      `xorm:"parent_id notnull int default 0"`
	ActionId uint      `xorm:"action_id notnull int default 0"`
	Creator  string    `xorm:"create_user notnull varchar(64)"`
	Created  time.Time `xorm:"create_at notnull datetime created"`
}
