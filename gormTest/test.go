package main

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
  "fmt"
  "time"
)
type student struct {
  StuNo int `gorm:"primary_key"`
  StuName string
  Sex string
  Time time.Time
  IsLegal int
  Info string `gorm:"default:'i am a student'"`
}

type class struct {
  //列名是字段名的蛇形小写
  ClassNo int
  Number int
}

type subject struct {
  stu_no int
  score int
  subject string
}

func main() {
  db, err := gorm.Open("mysql", "root:123456@/test?charset=utf8&parseTime=True&loc=Local")
   if err != nil {
    panic("连接数据库失败")
  }
  defer db.Close()

  // 全局禁用表名复数
  db.SingularTable(true)



  //create
  /*
  c := class{ClassNo:11111, Number:59}
  fmt.Println(db.Create(&c))
  */


  // insert
  /*
  stu := student{StuNo:2015110818, StuName:"quyi", Sex:"女", Time:time.Now(), IsLegal:1, Info:""}
  if db.NewRecord(stu) == false {// => 主键不为空返回`false`
    //delete
    fmt.Println("delete......")
    db.Delete(&stu)
  }
  fmt.Println("create......")
  db.Create(&stu)
  fmt.Println(db.NewRecord(stu))  // => 创建`user`后返回`false`
  */


  //select
 // var stu_find student
//   var stus_find []student
/*
  //1.查询
  // 获取第一条记录，按主键排序
  db.First(&stu_find)
  fmt.Println(stu_find)

  // 使用主键获取记录
  fmt.Println(db.First(&stu_find, 2015110818))
  fmt.Println(stu_find)

   //2.Where查询条件 (简单SQL)
  //获取第一个匹配记录
  fmt.Println(db.Where("stu_no = ?", 2015110818).First(&stu_find))
  fmt.Println(stu_find)

  // 获取所有匹配记录
  db.Where("sex = ?", "男").Find(&stus_find)
  for _,s := range stus_find {
    fmt.Println(s)
  }

  //IN
  db.Where("stu_no in (?)", []int{1, 3}).Find(&stus_find)
  fmt.Println(stus_find)

  // LIKE
  db.Where("time LIKE ?", "%06-10%").Find(&stus_find)
  fmt.Println(stus_find)

  // AND
  db.Where("sex = ? AND time >= ?", "女", "2017-06-10 12:33").Find(&stus_find)
  fmt.Println(stus_find)

  //3. Where查询条件 (Struct & Map)
  // Struct
  db.Where(&student{StuName: "quyi", Sex: "女"}).First(&stu_find)

  // Map
  db.Where(map[string]interface{}{"stu_name": "quyi", "sex": "女"}).Find(&stus_find)

  // 主键的Slice
  db.Where([]int{1, 2015110818, 4}).Find(&stus_find)

*/
  //4. join
  type Result struct {
    StuName  string
    Sex string
    Subject string
  }
  var results []Result
  //Scan将结果扫描到另一个结构中
  db.Table("student").Select("student.stu_name, student.sex, subject.score, subject.subject").Joins("left join subject on student.stu_no = subject.stu_no").Scan(&results)
  fmt.Println(results)

}