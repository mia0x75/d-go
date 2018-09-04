package uic

/*
DESC session;
+---------+------------------+------+-----+---------+----------------+
| Field   | Type             | Null | Key | Default | Extra          |
+---------+------------------+------+-----+---------+----------------+
| id      | int(10) unsigned | NO   | PRI | NULL    | auto_increment |
| uid     | int(10) unsigned | NO   | MUL | NULL    |                |
| sig     | varchar(32)      | NO   | MUL | NULL    |                |
| expired | int(10) unsigned | NO   |     | NULL    |                |
+---------+------------------+------+-----+---------+----------------+
*/

type Session struct {
	Id      uint   `xorm:"'id' notnull int pk autoincr"`
	UserId  uint   `xorm:"'uid' notnull int"`
	Sig     string `xorm:"'sig' notnull varchar(32)"`
	Expired uint   `xorm:"'expired' notnull int"`
}
