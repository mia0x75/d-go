package graph

/*
DESC endpoint;
+----------+------------------+------+-----+---------------------+-------------------------------+
| Field    | Type             | Null | Key | Default             | Extra                         |
+----------+------------------+------+-----+---------------------+-------------------------------+
| id       | int(10) unsigned | NO   | PRI | NULL                | auto_increment                |
| endpoint | varchar(255)     | NO   | UNI |                     |                               |
| ts       | int(11)          | YES  |     | NULL                |                               |
| t_create | datetime         | NO   |     | NULL                |                               |
| t_modify | timestamp        | NO   |     | current_timestamp() | on update current_timestamp() |
+----------+------------------+------+-----+---------------------+-------------------------------+
*/
