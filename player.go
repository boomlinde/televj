package main

import (
	"log"
	"os"
	"time"
)

type frame struct {
	index  int
	length int
	data   []byte
}

type anim []frame

func (a anim) Len() int           { return len(a) }
func (a anim) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a anim) Less(i, j int) bool { return a[i].index < a[j].index }

func (f frame) display(tc chan string) {
	go sleep(tc, f.length)
	os.Stdout.Write(f.data)
	os.Stdout.Sync()
}

func sleep(c chan string, t int) {
	time.Sleep(time.Duration(t) * time.Millisecond)
	c <- "next"
}
func (a anim) play(msg chan string) {
	findex := 0
	for m := range msg {
		if m == "next" {
			if len(a) > 0 {
				a[findex].display(msg)
				findex++
				if findex >= len(a) {
					findex -= len(a)
				}
			}
		} else if m == "stop" {
			return
		}
	}
}

func run(anims map[string]anim) chan string {
	a := anim{}
	stop := make(chan string)
	inch := make(chan string)
	go a.play(stop)
	go func() {
		for key := range inch {
			if a, ok := anims[key]; ok {
				stop <- "stop"
				stop = make(chan string)
				go a.play(stop)
				stop <- "next"
			} else {
				log.Printf("Missing anim: '%s'", key)
			}
		}
	}()
	return inch
}
