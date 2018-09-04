package dashboard

import "time"

/*
DESC dashboard_screen;
+-------+------------------+------+-----+---------------------+-------------------------------+
| Field | Type             | Null | Key | Default             | Extra                         |
+-------+------------------+------+-----+---------------------+-------------------------------+
| id    | int(11) unsigned | NO   | PRI | NULL                | auto_increment                |
| pid   | int(11) unsigned | NO   | MUL | 0                   |                               |
| name  | char(128)        | NO   |     | NULL                |                               |
| time  | timestamp        | NO   |     | current_timestamp() | on update current_timestamp() |
+-------+------------------+------+-----+---------------------+-------------------------------+
*/
type Screen struct {
	Id   uint      `xorm:"'id' notnull int pk autoincr"`
	PId  uint      `xorm:"'pid' notnull int default 0"`
	Name string    `xorm:"'name' notnull char(128)"`
	Time time.Time `xorm:"'time' notnull datetime updated"`
}
