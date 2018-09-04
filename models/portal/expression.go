package portal

/*
DESC expression;
+-------------+------------------+------+-----+---------+----------------+
| Field       | Type             | Null | Key | Default | Extra          |
+-------------+------------------+------+-----+---------+----------------+
| id          | int(10) unsigned | NO   | PRI | NULL    | auto_increment |
| expression  | varchar(1024)    | NO   |     | NULL    |                |
| func        | varchar(16)      | NO   |     | all(#1) |                |
| op          | varchar(8)       | NO   |     |         |                |
| right_value | varchar(16)      | NO   |     |         |                |
| max_step    | int(11)          | NO   |     | 1       |                |
| priority    | tinyint(4)       | NO   |     | 0       |                |
| note        | varchar(1024)    | NO   |     |         |                |
| action_id   | int(10) unsigned | NO   |     | 0       |                |
| create_user | varchar(64)      | NO   |     |         |                |
| pause       | tinyint(1)       | NO   |     | 0       |                |
+-------------+------------------+------+-----+---------+----------------+
*/
