package portal

/*
DESC grp_host;
+---------+------------------+------+-----+---------+-------+
| Field   | Type             | Null | Key | Default | Extra |
+---------+------------------+------+-----+---------+-------+
| grp_id  | int(10) unsigned | NO   | MUL | NULL    |       |
| host_id | int(10) unsigned | NO   | MUL | NULL    |       |
+---------+------------------+------+-----+---------+-------+
*/
type GrpHost struct {
	GrpId  uint `xorm:"'grp_id' notnull int"`
	HostId uint `xorm:"'host_id' notnull int"`
}

func (s *GrpHost) TableName() string {
	return "grp_host"
}
