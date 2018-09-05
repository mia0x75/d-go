package portal

import "time"

/*
DESC metric;
+-----------+------------------+------+-----+---------------------+----------------+
| Field     | Type             | Null | Key | Default             | Extra          |
+-----------+------------------+------+-----+---------------------+----------------+
| id        | int(10) unsigned | NO   | PRI | NULL                | auto_increment |
| name      | varchar(255)     | NO   | UNI | NULL                |                |
| update_at | timestamp        | NO   |     | current_timestamp() |                |
+-----------+------------------+------+-----+---------------------+----------------+
*/
type Metric struct {
	Id      uint      `xorm:"'id' notnull int pk autoincr"`
	Name    string    `xorm:"'name' notnull varchar(255) unique"`
	Created time.Time `xorm:"'update_at' notnull datetime created"`
}

func (s *Metric) TableName() string {
	return "metric"
}
