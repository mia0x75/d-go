package portal

import "time"

/*
DESC plugin_dir;
+-------------+------------------+------+-----+---------------------+----------------+
| Field       | Type             | Null | Key | Default             | Extra          |
+-------------+------------------+------+-----+---------------------+----------------+
| id          | int(10) unsigned | NO   | PRI | NULL                | auto_increment |
| grp_id      | int(10) unsigned | NO   | MUL | NULL                |                |
| dir         | varchar(255)     | NO   |     | NULL                |                |
| create_user | varchar(64)      | NO   |     |                     |                |
| create_at   | timestamp        | NO   |     | current_timestamp() |                |
+-------------+------------------+------+-----+---------------------+----------------+*/
type PluginDir struct {
	Id      uint      `xorm:"'id' notnull int pk autoincr"`
	GrpId   uint      `xorm:"'grp_id' notnull int"`
	Dir     string    `xorm:"'dir' notnull varchar(255)"`
	Creator string    `xorm:"'create_user' notnull varchar(64)"`
	Created time.Time `xorm:"'create_at' notnull datetime created"`
}

func (s *PluginDir) TableName() string {
	return "plugin_dir"
}
