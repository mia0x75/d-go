package uic

import (
	"time"

	"github.com/go-xorm/xorm"
)

/*
DESC user;
+---------+------------------+------+-----+---------------------+----------------+
| Field   | Type             | Null | Key | Default             | Extra          |
+---------+------------------+------+-----+---------------------+----------------+
| id      | int(10) unsigned | NO   | PRI | NULL                | auto_increment |
| name    | varchar(64)      | NO   | UNI | NULL                |                |
| passwd  | varchar(64)      | NO   |     |                     |                |
| cnname  | varchar(128)     | NO   |     |                     |                |
| email   | varchar(255)     | NO   |     |                     |                |
| phone   | varchar(16)      | NO   |     |                     |                |
| im      | varchar(32)      | NO   |     |                     |                |
| qq      | varchar(16)      | NO   |     |                     |                |
| role    | tinyint(4)       | NO   |     | 0                   |                |
| creator | int(10) unsigned | NO   |     | 0                   |                |
| created | timestamp        | NO   |     | current_timestamp() |                |
+---------+------------------+------+-----+---------------------+----------------+
*/
type User struct {
	Id        uint      `xorm:"'id' notnull int pk autoincr"`
	Login     string    `xorm:"'name' notnull varchar(64) unique"`
	Password  string    `xorm:"'passwd' notnull varchar(64)"` // fixed-length
	Name      string    `xorm:"'cnname' notnull varchar(128)"`
	Email     string    `xorm:"'email' notnull varchar(255)"`
	Phone     string    `xorm:"'phone' notnull varchar(16)"`
	Im        string    `xorm:"'im' notnull varchar(32)"`
	Tim       string    `xorm:"'qq' notnull varchar(16)"`
	Role      int8      `xorm:"'role' notnull int default 0"`
	CreatedBy uint      `xorm:"'creator' notnull int default 0"`
	Created   time.Time `xorm:"'created' notnull datetime created"`
}

func (s *User) TableName() string {
	return "user"
}

func (s *User) BeforeInsert() {
	// s.CreationDate = time.Now().Local()
	// s.UpdatedUnix = s.CreatedUnix
}

func (s *User) BeforeUpdate() {
	// s.UpdatedUnix = time.Now().Unix()
}

func (s *User) AfterSet(colName string, _ xorm.Cell) {
	// switch colName {
	// case "created_unix":
	// 	s.Created = time.Unix(s.CreatedUnix, 0).Local()
	// case "updated_unix":
	// 	s.Updated = time.Unix(s.UpdatedUnix, 0).Local()
	// }
}
