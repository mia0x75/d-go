package portal

/*
DESC strategy;
+-------------+------------------+------+-----+---------+----------------+
| Field       | Type             | Null | Key | Default | Extra          |
+-------------+------------------+------+-----+---------+----------------+
| id          | int(10) unsigned | NO   | PRI | NULL    | auto_increment |
| metric      | varchar(128)     | NO   |     |         |                |
| tags        | varchar(256)     | NO   |     |         |                |
| max_step    | int(11)          | NO   |     | 1       |                |
| priority    | tinyint(4)       | NO   |     | 0       |                |
| func        | varchar(16)      | NO   |     | all(#1) |                |
| op          | varchar(8)       | NO   |     |         |                |
| right_value | varchar(64)      | NO   |     | NULL    |                |
| note        | varchar(128)     | NO   |     |         |                |
| run_begin   | varchar(16)      | NO   |     |         |                |
| run_end     | varchar(16)      | NO   |     |         |                |
| tpl_id      | int(10) unsigned | NO   | MUL | 0       |                |
+-------------+------------------+------+-----+---------+----------------+
*/
