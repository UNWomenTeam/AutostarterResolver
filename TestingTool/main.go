package main

import (
	"fmt"
	"github.com/xela07ax/toolsXela/tp"
	"path/filepath"
	"time"
)

func main() {
	fmt.Println("Hola! +w1")
	dir, err := tp.BinDir()
	tp.Fck(err)
	fmt.Println(dir)
	fileName := "hellower.md"
	f, err := tp.OpenWriteFile(filepath.Join(dir, fileName))
	tp.Fck(err)
	writed, err := f.Write([]byte(fmt.Sprintf("Its ok [%v]\n", time.Now())))
	fmt.Println(writed)
	tp.Fck(err)
	fmt.Println("Good by! +w1")
}
