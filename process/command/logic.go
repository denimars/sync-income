package command

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"sync-finance/util"
	"time"
)

func Execute() {
	util.LoadEnv()
	timeNow := time.Now()
	out, err := exec.Command("mysqldump", os.Getenv("MYSQL_DB_NAME"), "-u", os.Getenv("MYSQL_USER"), fmt.Sprintf("-p%v", os.Getenv("MYSQL_PASSWORD"))).Output()
	if err != nil {
		fmt.Printf("%s", err)
	}

	reader := bytes.NewReader(out)

	file, err := os.Create(fmt.Sprintf("%v%v.sql", os.Getenv("MYSQL_DB_NAME"), timeNow.Format("200601021504")))
	if err != nil {
		log.Fatalf("Error creating file: %s", err)
	}
	defer file.Close()

	_, err = io.Copy(file, reader)
	if err != nil {
		log.Fatalf("Error writing to file: %s", err)
	}

	fmt.Println("Backup completed. File saved.")

}
