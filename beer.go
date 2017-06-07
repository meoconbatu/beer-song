package beer

import (
	"errors"
	"strconv"
	"strings"
)

const testVersion = 1

var lyricTemplate = "? of beer on the wall, ? of beer.\nTake one down and pass it around, ! of beer on the wall.\n"
var lyric0 = "No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.\n"
var lyric1 = "1 bottle of beer on the wall, 1 bottle of beer.\nTake it down and pass it around, no more bottles of beer on the wall.\n"
var min, max = 0, 99

func Verses(upperBound, lowerBound int) (verses string, err error) {
	if lowerBound > upperBound {
		return "", errors.New("start less than stop")
	}
	for i := lowerBound; i <= upperBound; i++ {
		versei, err := Verse(i)
		if err != nil {
			return "", err
		}
		verses = versei + "\n" + verses
	}
	return verses, nil
}
func Verse(index int) (string, error) {
	if index < min || index > max {
		return "", errors.New("invalid start/stop")
	}
	if index == 0 {
		return lyric0, nil
	}
	if index == 1 {
		return lyric1, nil
	}
	verse := strings.Replace(lyricTemplate, "?", makeBottleString(index), -1)
	verse = strings.Replace(verse, "!", makeBottleString(index-1), -1)
	return verse, nil
	// r := strings.NewReplacer("?", makeBottleString(index), "!", makeBottleString(index-1))
	// return r.Replace(lyricTemplate), nil
}
func makeBottleString(index int) string {
	if index == 1 {
		return strconv.Itoa(index) + " bottle"
	}
	if index == 0 {
		return "no more bottles"
	}
	return strconv.Itoa(index) + " bottles"
}
func Song() string {
	s, _ := Verses(max, min)
	return s
}
