package main

import (
	"github.com/boomlinde/teletext"
	"io/ioutil"
	"path"
	"regexp"
	"strconv"
	"strings"
)

func loadfiles(p string) map[string]anim {
	files, err := ioutil.ReadDir(p)
	fatal(err)

	matcher, err := regexp.Compile("[a-z][0-9]+_[0-9]+\\.ttv")
	fatal(err)

	matches := []string{}
	for _, s := range files {
		if matcher.MatchString(strings.ToLower(path.Base(s.Name()))) {
			matches = append(matches, path.Join(p, s.Name()))
		}
	}

	anims := map[string]anim{}
	for _, m := range matches {
		name := path.Base(m)
		fields := strings.Split(name[1:], "_")

		index, err := strconv.Atoi(fields[0])
		fatal(err)

		length, err := strconv.Atoi(fields[1][:len(fields[1])-4])
		fatal(err)

		data, err := ioutil.ReadFile(m)
		fatal(err)

		ttxframe := teletext.ConvertTTV("televj", 100, data)

		key := strings.ToLower(string(name[0]))
		_, ok := anims[key]
		if !ok {
			anims[key] = anim{}
		}
		anims[key] = append(anims[key], frame{index, length, ttxframe.Serialize()})
	}

	return anims
}
