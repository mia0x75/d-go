package portal

import "time"

/*
DESC cluster;
+-------------+------------------+------+-----+---------------------+-------------------------------+
| Field       | Type             | Null | Key | Default             | Extra                         |
+-------------+------------------+------+-----+---------------------+-------------------------------+
| id          | int(10) unsigned | NO   | PRI | NULL                | auto_increment                |
| grp_id      | int(11)          | NO   |     | NULL                |                               |
| numerator   | varchar(10240)   | NO   |     | NULL                |                               |
| denominator | varchar(10240)   | NO   |     | NULL                |                               |
| endpoint    | varchar(255)     | NO   |     | NULL                |                               |
| metric      | varchar(255)     | NO   |     | NULL                |                               |
| tags        | varchar(255)     | NO   |     | NULL                |                               |
| ds_type     | varchar(255)     | NO   |     | NULL                |                               |
| step        | int(11)          | NO   |     | NULL                |                               |
| last_update | timestamp        | NO   |     | current_timestamp() | on update current_timestamp() |
| creator     | varchar(255)     | NO   |     | NULL                |                               |
+-------------+------------------+------+-----+---------------------+-------------------------------+
*/
type Cluster struct {
	Id          uint      `xorm:"id notnull int pk autoincr"`
	GrpId       int       `xorm:"grp_id notnull int"`
	Numerator   string    `xorm:"numerator notnull varchar(10240)"`
	Denominator string    `xorm:"denominator notnull varchar(10240)"`
	Endpoint    string    `xorm:"endpoint notnull varchar(255)"`
	Metric      string    `xorm:"metric notnull varchar(255)"`
	Tags        string    `xorm:"tags notnull varchar(255)"`
	Type        string    `xorm:"ds_type notnull varchar(255)"`
	Step        int       `xorm:"step notnull int"`
	Creator     string    `xorm:"creator notnull varchar(255)"`
	Updated     time.Time `xorm:"last_update notnull datetime updated"`
}
