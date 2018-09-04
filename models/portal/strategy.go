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
type Strategy struct {
	Id         uint   `xorm:"'id' notnull int pk autoincr"`
	Metric     string `xorm:"'metric' notnull varchar(128)"`
	Tags       string `xorm:"'tags' notnull varchar(256)"`
	MaxStep    int    `xorm:"'max_step' notnull int default 1"`
	Priority   int8   `xorm:"'priority' notnull tinyint default 0"`
	Func       string `xorm:"'func' notnull varchar(16) default 'all(#1)'"`
	Op         string `xorm:"'op' notnull varchar(8)"`
	RightValue string `xorm:"'right_value' notnull varchar(64)"`
	Note       string `xorm:"'note' notnull varchar(128)"`
	RunBegin   string `xorm:"'run_begin' notnull varchar(16)"`
	RunEnd     string `xorm:"'run_end' notnull varchar(16)"`
	TplId      uint   `xorm:"'tpl_id' notnull int default 0"`
}
