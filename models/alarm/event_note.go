package alarm

import "time"

/*
DESC event_note;
+--------------+------------------+------+-----+---------+----------------+
| Field        | Type             | Null | Key | Default | Extra          |
+--------------+------------------+------+-----+---------+----------------+
| id           | mediumint(9)     | NO   | PRI | NULL    | auto_increment |
| event_caseId | varchar(50)      | YES  | MUL | NULL    |                |
| note         | varchar(300)     | YES  |     | NULL    |                |
| case_id      | varchar(20)      | YES  |     | NULL    |                |
| status       | varchar(15)      | YES  |     | NULL    |                |
| timestamp    | timestamp        | YES  |     | NULL    |                |
| user_id      | int(10) unsigned | YES  | MUL | NULL    |                |
+--------------+------------------+------+-----+---------+----------------+
*/
type EventNote struct {
	Id          uint      `xorm:"'id' notnull int pk autoincr"`
	EventCaseId string    `xorm:"'event_caseId' null varchar(50)"`
	Note        string    `xorm:"'note' notnull varchar(300)"`
	CaseId      uint      `xorm:"'case_id' null int"`
	Status      string    `xorm:"'status' null varchar(15)"`
	Time        time.Time `xorm:"'timestamp' null datetime"`
	UserId      uint      `xorm:"'user_id' null int"`
}
