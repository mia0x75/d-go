package portal

/*
DESC action;
+----------------------+------------------+------+-----+---------+----------------+
| Field                | Type             | Null | Key | Default | Extra          |
+----------------------+------------------+------+-----+---------+----------------+
| id                   | int(10) unsigned | NO   | PRI | NULL    | auto_increment |
| uic                  | varchar(255)     | NO   |     |         |                |
| url                  | varchar(255)     | NO   |     |         |                |
| callback             | tinyint(4)       | NO   |     | 0       |                |
| before_callback_sms  | tinyint(4)       | NO   |     | 0       |                |
| before_callback_mail | tinyint(4)       | NO   |     | 0       |                |
| after_callback_sms   | tinyint(4)       | NO   |     | 0       |                |
| after_callback_mail  | tinyint(4)       | NO   |     | 0       |                |
+----------------------+------------------+------+-----+---------+----------------+
*/
type Action struct {
	Id                 uint   `xorm:"id notnull int pk autoincr"`
	Uic                string `xorm:"uic notnull varchar(255)"`
	Url                string `xorm:"url notnull varchar(255)"`
	Callback           int8   `xorm:"callback notnull tinyint"`
	BeforeCallbackSMS  int8   `xorm:"before_callback_sms notnull tinyint"`
	BeforeCallbackMail int8   `xorm:"before_callback_mail notnull tinyint"`
	AfterCallbackSMS   int8   `xorm:"after_callback_sms notnull tinyint"`
	AfterCallbackMail  int8   `xorm:"after_callback_mail notnull tinyint"`
}
