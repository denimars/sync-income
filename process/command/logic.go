package command

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

func Execute() {
	out, err := exec.Command("mysqldump", "db_name", "-u", "root", "-ppassword").Output()
	if err != nil {
		fmt.Printf("%s", err)
	}

	reader := bytes.NewReader(out)

	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Create(fmt.Sprintf("%v%v%v.sql", "backup", "2", "2"))
	if err != nil {
		log.Fatalf("Error creating file: %s", err)
	}
	defer file.Close()

	_, err = io.Copy(file, reader)
	if err != nil {
		log.Fatalf("Error writing to file: %s", err)
	}

	fmt.Println("Backup completed. File saved as backup.sql")

}
