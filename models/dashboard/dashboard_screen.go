package dashboard

/*
DESC dashboard_screen;
+-------+------------------+------+-----+---------------------+-------------------------------+
| Field | Type             | Null | Key | Default             | Extra                         |
+-------+------------------+------+-----+---------------------+-------------------------------+
| id    | int(11) unsigned | NO   | PRI | NULL                | auto_increment                |
| pid   | int(11) unsigned | NO   | MUL | 0                   |                               |
| name  | char(128)        | NO   |     | NULL                |                               |
| time  | timestamp        | NO   |     | current_timestamp() | on update current_timestamp() |
+-------+------------------+------+-----+---------------------+-------------------------------+
*/
