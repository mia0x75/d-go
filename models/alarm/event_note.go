package alarm

/*
DESC event_note;
+--------------+------------------+------+-----+---------+----------------+
| Field        | Type             | Null | Key | Default | Extra          |
+--------------+------------------+------+-----+---------+----------------+
| id           | mediumint(9)     | NO   | PRI | NULL    | auto_increment |
| event_caseId | varchar(50)      | YES  | MUL | NULL    |                |
| note         | varchar(300)     | YES  |     | NULL    |                |
| case_id      | varchar(20)      | YES  |     | NULL    |                |
| status       | varchar(15)      | YES  |     | NULL    |                |
| timestamp    | timestamp        | YES  |     | NULL    |                |
| user_id      | int(10) unsigned | YES  | MUL | NULL    |                |
+--------------+------------------+------+-----+---------+----------------+
*/
