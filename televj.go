package main

import (
	"bufio"
	"golang.org/x/crypto/ssh/terminal"
	"log"
	"os"
)

func main() {
	oldterm, err := terminal.MakeRaw(0)
	fatal(err)
	defer terminal.Restore(0, oldterm)

	anims := loadfiles(os.Args[1])
	inch := run(anims)
	r := bufio.NewReader(os.Stdin)

	log.Println("Running")
	for {
		key, _, err := r.ReadRune()
		fatal(err)
		if key == '\x1b' {
			os.Exit(0)
		}
		inch <- string(key)
	}
}
