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
type Graph struct {
	Id       uint   `xorm:"'id' notnull int pk autoincr"`
	Title    string `xorm:"'title' notnull char(128)"`
	Hosts    string `xorm:"'hosts' notnull varchar(10240)"`
	Counters string `xorm:"'counters' notnull varchar(1024)"`
	ScreenId uint   `xorm:"'screen_id' notnull int"`
	Timespan uint   `xorm:"'timespan' notnull int default 3600"`
	Type     string `xorm:"'graph_type' notnull char(2) default 'h'"`
	Method   string `xorm:"'method' notnull char(8)"`
	Position uint   `xorm:"'position' notnull int default 0"`
	Tags     string `xorm:"'falcon_tags' notnull varchar(512)"`
}
