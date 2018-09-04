package graph

import "time"

/*
DESC tag_endpoint;
+-------------+------------------+------+-----+---------------------+-------------------------------+
| Field       | Type             | Null | Key | Default             | Extra                         |
+-------------+------------------+------+-----+---------------------+-------------------------------+
| id          | int(10) unsigned | NO   | PRI | NULL                | auto_increment                |
| tag         | varchar(255)     | NO   | MUL |                     |                               |
| endpoint_id | int(10) unsigned | NO   |     | NULL                |                               |
| ts          | int(11)          | YES  |     | NULL                |                               |
| t_create    | datetime         | NO   |     | NULL                |                               |
| t_modify    | timestamp        | NO   |     | current_timestamp() | on update current_timestamp() |
+-------------+------------------+------+-----+---------------------+-------------------------------+
*/
type EndpointTag struct {
	Id         uint      `xorm:"id notnull int pk autoincr"`
	Tag        string    `xorm:"tag notnull varchar(255)"`
	EndpointId uint      `xorm:"endpoint_id notnull int"`
	Timestamp  int       `xorm:"ts notnull int"`
	Created    time.Time `xorm:"t_create notnull datetime created"`
	Updated    time.Time `xorm:"t_modify notnull datetime updated"`
}
