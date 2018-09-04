package alarm

import "time"

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
| closed_note    | varchar(250)     | YES  |     | NULL       |       |
| user_modified  | int(10) unsigned | YES  |     | NULL       |       |
| tpl_creator    | varchar(64)      | YES  |     | NULL       |       |
| expression_id  | int(10) unsigned | YES  |     | NULL       |       |
| strategy_id    | int(10) unsigned | YES  |     | NULL       |       |
| template_id    | int(10) unsigned | YES  |     | NULL       |       |
| process_note   | mediumint(9)     | YES  |     | NULL       |       |
| process_status | varchar(20)      | YES  |     | unresolved |       |
| timestamp      | timestamp        | NO   |     | NULL       |       |
| update_at      | timestamp        | YES  |     | NULL       |       |
| closed_at      | timestamp        | YES  |     | NULL       |       |
+----------------+------------------+------+-----+------------+-------+
*/
type EventCases struct {
	Id             uint      `xorm:"id notnull int pk autoincr"`
	Endpoint       string    `xorm:"endpoint notnull varchar(100)"`
	Metric         string    `xorm:"metric notnull varchar(200)"`
	Func           string    `xorm:"func null varchar(50)"`
	Cond           string    `xorm:"cond notnull varchar(200)"`
	Note           string    `xorm:"note null varchar(500)"`
	MaxStep        uint      `xorm:"max_step null int"`
	CurrentStep    uint      `xorm:"current_step null int"`
	Priority       int       `xorm:"priority notnull int"`
	Status         string    `xorm:"status notnull varchar(20)"`
	CloseNote      string    `xorm:"closed_note null varchar(250)"`
	ModifiedUserId uint      `xorm:"user_modified null int"`
	TplCreator     string    `xorm:"tpl_creator null varchar(64)"`
	ExpressionId   uint      `xorm:"expression_id null int"`
	StrategyId     uint      `xorm:"strategy_id null int"`
	TemplateId     uint      `xorm:"template_id null int"`
	ProcessNote    uint      `xorm:"process_note null mediumint"`
	ProcessStatus  string    `xorm:"hosts null varchar(20)"`
	Timestamp      time.Time `xorm:"timestamp notnull datetime updated"`
	Updated        time.Time `xorm:"update_at notnull datetime updated"`
	Closed         time.Time `xorm:"closed_at notnull datetime"`
}
