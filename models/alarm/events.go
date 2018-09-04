package alarm

import "time"

/*
DESC events;
+--------------+------------------+------+-----+---------+----------------+
| Field        | Type             | Null | Key | Default | Extra          |
+--------------+------------------+------+-----+---------+----------------+
| id           | int(10)          | NO   | PRI | NULL    | auto_increment |
| event_caseId | varchar(50)      | YES  | MUL | NULL    |                |
| step         | int(10) unsigned | YES  |     | NULL    |                |
| cond         | varchar(200)     | NO   |     | NULL    |                |
| status       | int(3) unsigned  | YES  |     | 0       |                |
| timestamp    | timestamp        | YES  |     | NULL    |                |
+--------------+------------------+------+-----+---------+----------------+
*/
type Events struct {
	Id          uint      `xorm:"id notnull int pk autoincr"`
	EventCaseId string    `xorm:"event_caseId null varchar(50)"`
	Step        uint      `xorm:"step null int"`
	Cond        string    `xorm:"cond notnull varchar(200)"`
	Status      uint      `xorm:"status null int"`
	Time        time.Time `xorm:"timestamp null datetime created"`
}
