package alarm

/*
DESC event_cases;
+----------------+------------------+------+-----+------------+-------+
| Field          | Type             | Null | Key | Default    | Extra |
+----------------+------------------+------+-----+------------+-------+
| id             | varchar(50)      | NO   | PRI | NULL       |       |
| endpoint       | varchar(100)     | NO   | MUL | NULL       |       |
| metric         | varchar(200)     | NO   |     | NULL       |       |
| func           | varchar(50)      | YES  |     | NULL       |       |
| cond           | varchar(200)     | NO   |     | NULL       |       |
| note           | varchar(500)     | YES  |     | NULL       |       |
| max_step       | int(10) unsigned | YES  |     | NULL       |       |
| current_step   | int(10) unsigned | YES  |     | NULL       |       |
| priority       | int(6)           | NO   |     | NULL       |       |
| status         | varchar(20)      | NO   |     | NULL       |       |
| timestamp      | timestamp        | NO   |     | NULL       |       |
| update_at      | timestamp        | YES  |     | NULL       |       |
| closed_at      | timestamp        | YES  |     | NULL       |       |
| closed_note    | varchar(250)     | YES  |     | NULL       |       |
| user_modified  | int(10) unsigned | YES  |     | NULL       |       |
| tpl_creator    | varchar(64)      | YES  |     | NULL       |       |
| expression_id  | int(10) unsigned | YES  |     | NULL       |       |
| strategy_id    | int(10) unsigned | YES  |     | NULL       |       |
| template_id    | int(10) unsigned | YES  |     | NULL       |       |
| process_note   | mediumint(9)     | YES  |     | NULL       |       |
| process_status | varchar(20)      | YES  |     | unresolved |       |
+----------------+------------------+------+-----+------------+-------+
*/
