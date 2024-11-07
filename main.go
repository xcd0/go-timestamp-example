package main

import (
	"log"
	"os"
	"time"

	"github.com/djherbis/times"
)

func main() {
	log.SetFlags(log.Ltime | log.Lshortfile) // ログの出力書式を設定する

	if len(os.Args) <= 1 {
		log.Printf("argc: %v, argv: %#v", len(os.Args), os.Args)
		return
	}
	for i := 1; i < len(os.Args); i++ {
		check_time(os.Args[i])
	}

}

func check_time(p string) {
	t, err := times.Stat(p)
	if err != nil {
		log.Fatal(err.Error())
	}
	var atime, mtime, ctime, btime time.Time

	atime = t.AccessTime()
	mtime = t.ModTime()

	if t.HasChangeTime() {
		ctime = t.ChangeTime()
	}

	if t.HasBirthTime() {
		btime = t.BirthTime()
	}

	log.Printf("\n\tatime: %v\n\tmtime: %v\n\tctime: %v\n\tbtime: %v",
		atime, mtime, ctime, btime,
	)

}
