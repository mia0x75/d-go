package models

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
//varchar(25) notnull unique 'usr_name'
type User struct {
	Id        uint      `xorm:"id notnull int pk autoincr"`
	Login     string    `xorm:"name notnull varchar(64) unique"`
	Password  string    `xorm:"passwd notnull varchar(64)"` // fixed-length
	Name      string    `xorm:"cnname notnull varchar(128)"`
	Email     string    `xorm:"email notnull varchar(255)"`
	Phone     string    `xorm:"phone notnull varchar(16)"`
	Im        string    `xorm:"im notnull varchar(32)"`
	Tim       string    `xorm:"qq notnull varchar(16)"`
	Role      int8      `xorm:"role notnull int"`
	CreatedBy uint      `xorm:"creator notnull int"`
	Created   time.Time `xorm:"created notnull datetime created"`
}

func (u *User) TableName() string {
	return "user"
}

func (u *User) BeforeInsert() {
	// u.CreationDate = time.Now().Local()
	// u.UpdatedUnix = u.CreatedUnix
}

func (u *User) BeforeUpdate() {
	// u.UpdatedUnix = time.Now().Unix()
}

func (u *User) AfterSet(colName string, _ xorm.Cell) {
	// switch colName {
	// case "created_unix":
	// 	u.Created = time.Unix(u.CreatedUnix, 0).Local()
	// case "updated_unix":
	// 	u.Updated = time.Unix(u.UpdatedUnix, 0).Local()
	// }
}
