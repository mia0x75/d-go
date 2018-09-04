package dashboard

import "time"

/*
DESC tmp_graph;
+-----------+------------------+------+-----+---------------------+----------------+
| Field     | Type             | Null | Key | Default             | Extra          |
+-----------+------------------+------+-----+---------------------+----------------+
| id        | int(11) unsigned | NO   | PRI | NULL                | auto_increment |
| endpoints | varchar(10240)   | NO   |     |                     |                |
| counters  | varchar(10240)   | NO   |     |                     |                |
| ck        | varchar(32)      | NO   | UNI | NULL                |                |
| time_     | timestamp        | NO   |     | current_timestamp() |                |
+-----------+------------------+------+-----+---------------------+----------------+
*/
type TmpGraph struct {
	Id        uint      `xorm:"id notnull int pk autoincr"`
	Endpoints string    `xorm:"endpoints notnull varchar(10240)"`
	Counters  string    `xorm:"counters notnull varchar(10240)"`
	Ck        string    `xorm:"ck notnull varchar(32) unique"`
	Time      time.Time `xorm:"time_ notnull datetime created"`
}
