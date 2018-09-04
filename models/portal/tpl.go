package portal

/*
DESC tpl;
+-------------+------------------+------+-----+---------------------+----------------+
| Field       | Type             | Null | Key | Default             | Extra          |
+-------------+------------------+------+-----+---------------------+----------------+
| id          | int(10) unsigned | NO   | PRI | NULL                | auto_increment |
| tpl_name    | varchar(255)     | NO   | UNI |                     |                |
| parent_id   | int(10) unsigned | NO   |     | 0                   |                |
| action_id   | int(10) unsigned | NO   |     | 0                   |                |
| create_user | varchar(64)      | NO   | MUL |                     |                |
| create_at   | timestamp        | NO   |     | current_timestamp() |                |
+-------------+------------------+------+-----+---------------------+----------------+
*/
