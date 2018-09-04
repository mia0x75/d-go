package uic

import "time"

/*
DESC team;
+---------+------------------+------+-----+---------------------+----------------+
| Field   | Type             | Null | Key | Default             | Extra          |
+---------+------------------+------+-----+---------------------+----------------+
| id      | int(10) unsigned | NO   | PRI | NULL                | auto_increment |
| name    | varchar(64)      | NO   | UNI | NULL                |                |
| resume  | varchar(255)     | NO   |     |                     |                |
| creator | int(10) unsigned | NO   |     | 0                   |                |
| created | timestamp        | NO   |     | current_timestamp() |                |
+---------+------------------+------+-----+---------------------+----------------+
*/
type Team struct {
	Id        uint      `xorm:"id notnull int pk autoincr"`
	Name      string    `xorm:"name notnull varchar(64) unique"`
	Resume    string    `xorm:"passwd notnull varchar(255)"`
	CreatedBy uint      `xorm:"creator notnull int"`
	Created   time.Time `xorm:"'created' notnull datetime created"`
}
