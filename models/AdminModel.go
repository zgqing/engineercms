package models

import (
	// "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	// "strconv"
	// "strings"
	"time"
)

type AdminCategory struct {
	Id       int64     `form:"-"`
	ParentId int64     `orm:"null"`
	Title    string    `form:"title;text;title:",valid:"MinSize(1);MaxSize(20)"` //orm:"unique",
	Code     string    `orm:"null"`
	Grade    int       `orm:"null"`
	Created  time.Time `orm:"index","auto_now_add;type(datetime)"`
	Updated  time.Time `orm:"index","auto_now_add;type(datetime)"`
}

type AdminIpsegment struct {
	Id      int64     `form:"-"`
	Title   string    `form:"title;text;title:",valid:"MinSize(1);MaxSize(20)"` //orm:"unique",
	StartIp string    `orm:"not null"`
	EndIp   string    `orm:"null"`
	Iprole  int       `orm:"null"`
	Created time.Time `orm:"index","auto_now_add;type(datetime)"`
	Updated time.Time `orm:"index","auto_now_add;type(datetime)"`
}

type AdminCalenda struct {
	Id        int64     `form:"-"`
	Title     string    `form:"title;text;title:",valid:"MinSize(1);MaxSize(100)"` //orm:"unique",
	starttime time.Time `orm:"not null;type(datetime)"`
	endtime   time.Time `orm:"null;type(datetime)"`
	allday    int8      `orm:"not null;default(0)"`
	color     string    `orm:"null"`
}

// `id` int(11) NOT NULL AUTO_INCREMENT,
//   `title` varchar(100) NOT NULL,
//   `starttime` int(11) NOT NULL,
//   `endtime` int(11) DEFAULT NULL,
//   `allday` tinyint(1) NOT NULL DEFAULT '0',
//   `color` varchar(20) DEFAULT NULL,

func init() {
	orm.RegisterModel(new(AdminCategory), new(AdminIpsegment)) //, new(Article)
	orm.RegisterDriver("sqlite", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "database/engineer.db", 10)
}

//添加
func AddAdminCategory(pid int64, title, code string, grade int) (id int64, err error) {
	o := orm.NewOrm()
	// var category AdminCategory
	// if pid == "" {
	// 	category := &AdminCategory{
	// 		ParentId: 0,
	// 		Title:    title,
	// 		Code:     code,
	// 		Grade:    grade,
	// 		Created:  time.Now(),
	// 		Updated:  time.Now(),
	// 	}
	// 	id, err = o.Insert(category)
	// 	if err != nil {
	// 		return 0, err
	// 	}
	// } else {
	//pid转成64为
	// pidNum, err := strconv.ParseInt(pid, 10, 64)
	// if err != nil {
	// 	return 0, err
	// }
	category := &AdminCategory{
		ParentId: pid,
		Title:    title,
		Code:     code,
		Grade:    grade,
		Created:  time.Now(),
		Updated:  time.Now(),
	}
	id, err = o.Insert(category)
	if err != nil {
		return 0, err
	}
	// }
	return id, nil
}

//修改
func UpdateAdminCategory(cid int64, title, code string, grade int) error {
	o := orm.NewOrm()
	//id转成64为
	// cidNum, err := strconv.ParseInt(cid, 10, 64)
	// if err != nil {
	// 	return err
	// }
	category := &AdminCategory{Id: cid}
	if o.Read(category) == nil {
		category.Title = title
		category.Code = code
		category.Grade = grade
		category.Updated = time.Now()
		_, err := o.Update(category)
		if err != nil {
			return err
		}
	}
	return nil
}

//根据父级id取得所有
//如果父级id为空，则取所有一级category
func GetAdminCategory(pid int64) (categories []*AdminCategory, err error) {
	o := orm.NewOrm()
	categories = make([]*AdminCategory, 0)

	qs := o.QueryTable("AdminCategory") //这个表名AchievementTopic需要用驼峰式，
	// if pid != "" {                      //如果给定父id则进行过滤
	//pid转成64为
	// pidNum, err := strconv.ParseInt(pid, 10, 64)
	// if err != nil {
	// 	return nil, err
	// }
	_, err = qs.Filter("parentid", pid).All(&categories) //而这个字段parentid为何又不用呢
	if err != nil {
		return nil, err
	}

	return categories, err
	// } else { //如果不给定父id（PID=0），则取所有一级
	// _, err = qs.Filter("parentid", 0).All(&categories)
	// if err != nil {
	// 	return nil, err
	// }
	// return categories, err
	// }
}

//根据类别名字title查询所有下级分级category
func GetAdminCategoryTitle(title string) (categories []*AdminCategory, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("AdminCategory")
	var cate AdminCategory
	err = qs.Filter("title", title).One(&cate)
	// if pid != "" {
	// cate := AdminCategory{Title: title}这句无效
	categories = make([]*AdminCategory, 0)
	_, err = qs.Filter("parentid", cate.Id).All(&categories)
	if err != nil {
		return nil, err
	}
	return categories, err
	// } else { //如果不给定父id（PID=0），则取所有一级
	// _, err = qs.Filter("parentid", 0).All(&categories)
	// if err != nil {
	// return nil, err
	// }
	// return categories, err
	// }
}

//根据id查分级
func GetAdminCategorybyId(id int64) (category []*AdminCategory, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("AdminCategory")

	err = qs.Filter("id", id).One(&category)
	if err != nil {
		return nil, err
	}
	return category, err
}

//添加ip地址段
func AddAdminIpsegment(title, startip, endip string, iprole int) (id int64, err error) {
	o := orm.NewOrm()
	ipsegment := &AdminIpsegment{
		Title:   title,
		StartIp: startip,
		EndIp:   endip,
		Iprole:  iprole,
		Created: time.Now(),
		Updated: time.Now(),
	}
	id, err = o.Insert(ipsegment)
	if err != nil {
		return 0, err
	}
	return id, nil
}

//修改Ip地址段
func UpdateAdminIpsegment(cid int64, title, startip, endip string, iprole int) error {
	o := orm.NewOrm()
	ipsegment := &AdminIpsegment{Id: cid}
	if o.Read(ipsegment) == nil {
		ipsegment.Title = title
		ipsegment.StartIp = startip
		ipsegment.EndIp = endip
		ipsegment.Iprole = iprole
		ipsegment.Updated = time.Now()
		_, err := o.Update(ipsegment)
		if err != nil {
			return err
		}
	}
	return nil
}

//查询所有Ip地址段
func GetAdminIpsegment() (ipsegments []*AdminIpsegment, err error) {
	o := orm.NewOrm()
	// ipsegments = make([]*AdminIpsegment, 0)

	qs := o.QueryTable("AdminIpsegment") //这个表名AchievementTopic需要用驼峰式，
	// if pid != "" {                      //如果给定父id则进行过滤
	//pid转成64为
	// pidNum, err := strconv.ParseInt(pid, 10, 64)
	// if err != nil {
	// 	return nil, err
	// }
	_, err = qs.All(&ipsegments)
	if err != nil {
		return nil, err
	}

	return ipsegments, err
	// } else { //如果不给定父id（PID=0），则取所有一级
	// _, err = qs.Filter("parentid", 0).All(&categories)
	// if err != nil {
	// 	return nil, err
	// }
	// return categories, err
	// }
}
