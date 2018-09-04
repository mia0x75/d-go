package alarm

/*
DESC events;
+--------------+------------------+------+-----+---------+----------------+
| Field        | Type             | Null | Key | Default | Extra          |
+--------------+------------------+------+-----+---------+----------------+
| id           | int(10)          | NO   | PRI | NULL    | auto_increment |
| event_caseId | varchar(50)      | YES  | MUL | NULL    |                |
| step         | int(10) unsigned | YES  |     | NULL    |                |
| cond         | varchar(200)     | NO   |     | NULL    |                |
| status       | int(3) unsigned  | YES  |     | 0       |                |
| timestamp    | timestamp        | YES  |     | NULL    |                |
+--------------+------------------+------+-----+---------+----------------+
*/
