package graph

import "time"

/*
DESC endpoint;
+----------+------------------+------+-----+---------------------+-------------------------------+
| Field    | Type             | Null | Key | Default             | Extra                         |
+----------+------------------+------+-----+---------------------+-------------------------------+
| id       | int(10) unsigned | NO   | PRI | NULL                | auto_increment                |
| endpoint | varchar(255)     | NO   | UNI |                     |                               |
| ts       | int(11)          | YES  |     | NULL                |                               |
| t_create | datetime         | NO   |     | NULL                |                               |
| t_modify | timestamp        | NO   |     | current_timestamp() | on update current_timestamp() |
+----------+------------------+------+-----+---------------------+-------------------------------+
*/
type Endpoint struct {
	Id        uint      `xorm:"id notnull int pk autoincr"`
	Endpoint  string    `xorm:"endpoint notnull varchar(255) unique"`
	Timestamp int       `xorm:"ts notnull int"`
	Created   time.Time `xorm:"t_create notnull datetime created"`
	Updated   time.Time `xorm:"t_modify notnull datetime updated"`
}
