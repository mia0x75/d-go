package portal

import "time"

/*
DESC host;
+----------------+------------------+------+-----+---------------------+-------------------------------+
| Field          | Type             | Null | Key | Default             | Extra                         |
+----------------+------------------+------+-----+---------------------+-------------------------------+
| id             | int(10) unsigned | NO   | PRI | NULL                | auto_increment                |
| hostname       | varchar(255)     | NO   | UNI |                     |                               |
| ip             | varchar(16)      | NO   |     |                     |                               |
| agent_version  | varchar(16)      | NO   |     |                     |                               |
| plugin_version | varchar(128)     | NO   |     |                     |                               |
| maintain_begin | int(10) unsigned | NO   |     | 0                   |                               |
| maintain_end   | int(10) unsigned | NO   |     | 0                   |                               |
| update_at      | timestamp        | NO   |     | current_timestamp() | on update current_timestamp() |
+----------------+------------------+------+-----+---------------------+-------------------------------+
*/
type Host struct {
	Id            uint      `xorm:"'id' notnull int pk autoincr"`
	Hostname      string    `xorm:"'hostname' notnull varchar(1024) unique"`
	Ip            string    `xorm:"'ip' notnull varchar(16)"`
	AgentVersion  string    `xorm:"'agent_version' notnull varchar(8)"`
	PluginVersion string    `xorm:"'plugin_version' notnull varchar(16)"`
	MaintainBegin int       `xorm:"'maintain_begin' notnull int default 0"`
	MaintainEnd   int       `xorm:"'maintain_end' notnull int default 0"`
	Updated       time.Time `xorm:"'update_at' notnull datetime updated"`
}
