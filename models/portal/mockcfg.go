package portal

import "time"

/*
DESC mockcfg;
+----------+---------------------+------+-----+---------------------+-------------------------------+
| Field    | Type                | Null | Key | Default             | Extra                         |
+----------+---------------------+------+-----+---------------------+-------------------------------+
| id       | bigint(20) unsigned | NO   | PRI | NULL                | auto_increment                |
| name     | varchar(255)        | NO   | UNI |                     |                               |
| obj      | varchar(10240)      | NO   |     |                     |                               |
| obj_type | varchar(255)        | NO   |     |                     |                               |
| metric   | varchar(128)        | NO   |     |                     |                               |
| tags     | varchar(1024)       | NO   |     |                     |                               |
| dstype   | varchar(32)         | NO   |     | GAUGE               |                               |
| step     | int(11) unsigned    | NO   |     | 60                  |                               |
| mock     | double              | NO   |     | 0                   |                               |
| creator  | varchar(64)         | NO   |     |                     |                               |
| t_create | datetime            | NO   |     | NULL                |                               |
| t_modify | timestamp           | NO   |     | current_timestamp() | on update current_timestamp() |
+----------+---------------------+------+-----+---------------------+-------------------------------+
*/
type Mockcfg struct {
	Id         uint      `xorm:"id notnull int pk autoincr"`
	Name       string    `xorm:"name notnull uniqe varchar(255)"`
	Object     string    `xorm:"obj notnull varchar(10240)"`
	ObjectType string    `xorm:"obj_type notnull varchar(255)"`
	Metric     string    `xorm:"metric notnull varchar(128)"`
	Tags       string    `xorm:"tags notnull varchar(1024)"`
	Type       string    `xorm:"dstype notnull varchar(32)"`
	Step       uint      `xorm:"step notnull int"`
	Mock       float64   `xorm:"mock notnull double"`
	Creator    string    `xorm:"creator notnull varchar(64)"`
	Created    time.Time `xorm:"t_create notnull datetime created"`
	Updated    time.Time `xorm:"t_modify notnull datetime updated"`
}
