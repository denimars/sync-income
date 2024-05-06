package main

import (
	"sync-finance/database"
	"sync-finance/process/income"
)

func main() {
	// gmt8TimeString := "2024-05-06T12:30:45"

	// Parse the time string into a time object, assuming it's in the format "2006-01-02T15:04:05"
	// gm8Location, _ := time.LoadLocation("Asia/Makassar")
	// gmt8Time, err := time.ParseInLocation("2006-01-02T15:04:05", gmt8TimeString, gm8Location)
	// if err != nil {
	// 	fmt.Println("Error parsing time string:", err)
	// 	return
	// }

	// // Convert the GMT+8 time to UTC
	// fmt.Println(gmt8Time)
	// utcTime := gmt8Time.UTC()

	// // Print the UTC time
	// fmt.Println("UTC Time:", utcTime.Format("2006-01-02T15:04:05"))

	mysqlDB := database.MysqlConnection()
	repoIncome := income.NewRepository(mysqlDB)
	logicIncome := income.NewLogic(repoIncome)
	logicIncome.GetData()

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
