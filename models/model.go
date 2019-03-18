package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type User struct {
	Id   int
	Name string
	Pwd  string
}

type ArticleType struct {
	Id      int
	Tname   string
	Article []*Article `orm:"reverse(many)"`
}

type Article struct {
	Id          int          `orm:"pk;auto"`
	ArtiName    string       `orm:"size(20)"`
	Atime       time.Time    `orm:"auto_now"`
	Acount      int          `orm:"default(0);null"`
	Acontent    string
	Aimg        string
	ArticleType *ArticleType `orm:"rel(fk)"`
}

func init() {
	orm.RegisterDataBase("default", "mysql", "root:mysqlwsf123@tcp(localhost:3306)/test?charset=utf8")
	orm.RegisterModel(new(User), new(Article), new(ArticleType))
	orm.RunSyncdb("default", false, true)
}
