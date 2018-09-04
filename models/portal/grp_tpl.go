package portal

/*
DESC grp_tpl;
+-----------+------------------+------+-----+---------+-------+
| Field     | Type             | Null | Key | Default | Extra |
+-----------+------------------+------+-----+---------+-------+
| grp_id    | int(10) unsigned | NO   | MUL | NULL    |       |
| tpl_id    | int(10) unsigned | NO   | MUL | NULL    |       |
| bind_user | varchar(64)      | NO   |     |         |       |
+-----------+------------------+------+-----+---------+-------+
*/
type GrpTpl struct {
	GrpId    uint   `xorm:"grp_id notnull int"`
	TplId    uint   `xorm:"tpl_id notnull int"`
	BindUser string `xorm:"bind_user notnull varchar(64)"`
}
