package dbtest

import (
	"context"
)

// 测试cassandra的读写
func ReadCassandra(uId int) {
	var (
		userId int
		/*age     int
		first   string
		last    string*/
	)
	ctx := context.Background()
	err = Session.Query(`SELECT uid FROM user1 WHERE uid = ?`,
		uId).WithContext(ctx).Scan(&userId)

	if err != nil {
		//fmt.Println("ReadCassandra err" , err)
		return
	}
	//fmt.Println("ReadCassandra := ", userId, age, first, last)
}

func WCassandra(uid int){

	ctx := context.Background()
	err = Session.Query(`INSERT INTO user1 (uid, age, first ,last) VALUES (?, ?, ? ,?)`,
		uid, 100, "hello world","ddfj").WithContext(ctx).Exec()
	if err != nil {
		//fmt.Println("ReadCassandra err" , err)
		return
	}
	//fmt.Println("ReadCassandra INSERT succ" , err)
}
