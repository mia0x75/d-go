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
type Expression struct {
	Id         uint   `xorm:"id notnull int pk autoincr"`
	Expression string `xorm:"expression notnull varchar(1024)"`
	Func       string `xorm:"func notnull varchar(16)"`
	Op         string `xorm:"op notnull varchar(8)"`
	RightValue string `xorm:"right_value notnull varchar(16)"`
	MaxStep    int    `xorm:"max_step notnull int"`
	Priority   int8   `xorm:"priority notnull tinyint"`
	Note       string `xorm:"note notnull varchar(1024)"`
	ActionId   uint   `xorm:"action_id notnull int"`
	Creator    string `xorm:"create_user notnull varchar(64)"`
	Pause      int8   `xorm:"pause notnull tinyint"`
}
