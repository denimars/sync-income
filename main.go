package main

import (
	"sync-finance/database"
	"sync-finance/process/command"
	"sync-finance/process/income"
)

func main() {

	mysqlDB := database.MysqlConnection()
	repoIncome := income.NewRepository(mysqlDB)
	logicIncome := income.NewLogic(repoIncome)
	logicIncome.GetData("2024-03-01", "2024-03-31")
	command.Execute()

	// ctx := context.Background()
	// rd, _ := database.RedisConnection()
	// _, err := rd.Set(ctx, "FOO", "BAR", 0).Result()
	// if err != nil {
	// 	log.Println("failed connection")
	// }
	// //create
	// inValue, err := rd.Get(ctx, "FOO").Result()
	// if err != nil {
	// 	log.Println("database not found")
	// }

	// fmt.Println(inValue)

	// //update
	// _, exx := rd.Set(ctx, "FOO", "APA", 0).Result()
	// if exx != nil {
	// 	log.Println("failed connection")
	// }

	// //get
	// inValue, err = rd.Get(ctx, "FOO").Result()
	// if err != nil {
	// 	log.Println("database not found")
	// }
	// fmt.Println(inValue)
	// //delete
	// rd.Del(ctx, "FOO")
	// inValue, err = rd.Get(ctx, "FOO").Result()
	// if err != nil {
	// 	log.Println("database not found")
	// }
	// fmt.Println(inValue)
}
