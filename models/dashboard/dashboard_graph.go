package dashboard

/*
DESC dashboard_graph;
+-------------+------------------+------+-----+---------+----------------+
| Field       | Type             | Null | Key | Default | Extra          |
+-------------+------------------+------+-----+---------+----------------+
| id          | int(11) unsigned | NO   | PRI | NULL    | auto_increment |
| title       | char(128)        | NO   |     | NULL    |                |
| hosts       | varchar(10240)   | NO   |     |         |                |
| counters    | varchar(1024)    | NO   |     |         |                |
| screen_id   | int(11) unsigned | NO   | MUL | NULL    |                |
| timespan    | int(11) unsigned | NO   |     | 3600    |                |
| graph_type  | char(2)          | NO   |     | h       |                |
| method      | char(8)          | YES  |     |         |                |
| position    | int(11) unsigned | NO   |     | 0       |                |
| falcon_tags | varchar(512)     | NO   |     |         |                |
+-------------+------------------+------+-----+---------+----------------+
*/
