package portal

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
