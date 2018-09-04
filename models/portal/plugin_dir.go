package portal

import "time"

/*
DESC plugin_dir;
+-------------+------------------+------+-----+---------------------+----------------+
| Field       | Type             | Null | Key | Default             | Extra          |
+-------------+------------------+------+-----+---------------------+----------------+
| create_user | varchar(64)      | NO   |     |                     |                |
| create_at   | timestamp        | NO   |     | current_timestamp() |                |
+-------------+------------------+------+-----+---------------------+----------------+
*/
type PluginDir struct {
	Creator string    `xorm:"create_user notnull varchar(64)"`
	Created time.Time `xorm:"create_at notnull datetime created"`
}
