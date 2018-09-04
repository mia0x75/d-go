package portal

/*
DESC cluster;
+-------------+------------------+------+-----+---------------------+-------------------------------+
| Field       | Type             | Null | Key | Default             | Extra                         |
+-------------+------------------+------+-----+---------------------+-------------------------------+
| id          | int(10) unsigned | NO   | PRI | NULL                | auto_increment                |
| grp_id      | int(11)          | NO   |     | NULL                |                               |
| numerator   | varchar(10240)   | NO   |     | NULL                |                               |
| denominator | varchar(10240)   | NO   |     | NULL                |                               |
| endpoint    | varchar(255)     | NO   |     | NULL                |                               |
| metric      | varchar(255)     | NO   |     | NULL                |                               |
| tags        | varchar(255)     | NO   |     | NULL                |                               |
| ds_type     | varchar(255)     | NO   |     | NULL                |                               |
| step        | int(11)          | NO   |     | NULL                |                               |
| last_update | timestamp        | NO   |     | current_timestamp() | on update current_timestamp() |
| creator     | varchar(255)     | NO   |     | NULL                |                               |
+-------------+------------------+------+-----+---------------------+-------------------------------+
*/
