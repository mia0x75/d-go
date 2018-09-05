package uic

/*
DESC rel_team_user;
+-------+------------------+------+-----+---------+----------------+
| Field | Type             | Null | Key | Default | Extra          |
+-------+------------------+------+-----+---------+----------------+
| id    | int(10) unsigned | NO   | PRI | NULL    | auto_increment |
| tid   | int(10) unsigned | NO   | MUL | NULL    |                |
| uid   | int(10) unsigned | NO   | MUL | NULL    |                |
+-------+------------------+------+-----+---------+----------------+
*/
type RelationTeamUser struct {
	Id     uint `xorm:"'id' notnull int pk autoincr"`
	TeamId uint `xorm:"'tid' notnull int"`
	UserId uint `xorm:"'uid' notnull int"`
}

func (s *RelationTeamUser) TableName() string {
	return "rel_team_user"
}
