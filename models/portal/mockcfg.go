package portal

/*
DESC mockcfg;
+----------+---------------------+------+-----+---------------------+-------------------------------+
| Field    | Type                | Null | Key | Default             | Extra                         |
+----------+---------------------+------+-----+---------------------+-------------------------------+
| id       | bigint(20) unsigned | NO   | PRI | NULL                | auto_increment                |
| name     | varchar(255)        | NO   | UNI |                     |                               |
| obj      | varchar(10240)      | NO   |     |                     |                               |
| obj_type | varchar(255)        | NO   |     |                     |                               |
| metric   | varchar(128)        | NO   |     |                     |                               |
| tags     | varchar(1024)       | NO   |     |                     |                               |
| dstype   | varchar(32)         | NO   |     | GAUGE               |                               |
| step     | int(11) unsigned    | NO   |     | 60                  |                               |
| mock     | double              | NO   |     | 0                   |                               |
| creator  | varchar(64)         | NO   |     |                     |                               |
| t_create | datetime            | NO   |     | NULL                |                               |
| t_modify | timestamp           | NO   |     | current_timestamp() | on update current_timestamp() |
+----------+---------------------+------+-----+---------------------+-------------------------------+
*/
