package graph

import (
	"time"
)

/*
DESC endpoint_counter;
+-------------+------------------+------+-----+---------------------+-------------------------------+
| Field       | Type             | Null | Key | Default             | Extra                         |
+-------------+------------------+------+-----+---------------------+-------------------------------+
| id          | int(10) unsigned | NO   | PRI | NULL                | auto_increment                |
| endpoint_id | int(10) unsigned | NO   | MUL | NULL                |                               |
| counter     | varchar(255)     | NO   |     |                     |                               |
| step        | int(11)          | NO   |     | 60                  |                               |
| type        | varchar(16)      | NO   |     | NULL                |                               |
| ts          | int(11)          | YES  |     | NULL                |                               |
| t_create    | datetime         | NO   |     | NULL                |                               |
| t_modify    | timestamp        | NO   |     | current_timestamp() | on update current_timestamp() |
+-------------+------------------+------+-----+---------------------+-------------------------------+
*/

type EndpointCounter struct {
	Id         uint      `xorm:"id notnull int pk autoincr"`
	EndpointId uint      `xorm:"endpoint_id notnull int"`
	Counter    string    `xorm:"counter notnull varchar(255)"`
	Step       uint      `xorm:"step notnull int"`
	Type       string    `xorm:"type notnull varchar(255)"`
	Timestamp  int       `xorm:"ts notnull int"`
	Created    time.Time `xorm:"t_create notnull datetime created"`
	Updated    time.Time `xorm:"t_modify notnull datetime updated"`
}
