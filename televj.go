package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	/*
		anims := map[string]anim{
			"a": anim{
				frame{0, 500, "a_f1"},
				frame{1, 500, "a_f2"},
				frame{2, 500, "a_f3"},
				frame{3, 100, "a_f4"},
				frame{4, 500, "a_f5"},
				frame{5, 500, "a_f6"},
			},
			"b": anim{
				frame{0, 5000, "b_f1"},
				frame{1, 5000, "b_f2"},
				frame{2, 5000, "b_f3"},
				frame{3, 5000, "b_f4"},
				frame{4, 1000, "b_f5"},
				frame{5, 5000, "b_f6"},
			},
		}
	*/
	anims := loadfiles(os.Args[1])
	inch := run(anims)
	reader := bufio.NewReader(os.Stdin)

	for {
		key, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		if key == "quit\n" {
			os.Exit(0)
		}
		inch <- key[:len(key)-1]
	}
}
