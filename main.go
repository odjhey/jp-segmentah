package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/go-ego/gse"
)

var (
	text = flag.String("text", `ひっそり遠くから、もしかすると離（はな）し難（がた）いのか。
		黙々（もくもく）と静かに、もしかするととても価値（かち）があるのか。
		僕はまだここで待っている`, "単語セグメンテーションのテキスト")
)

func main() {
	flag.Parse()

	var seg gse.Segmenter
	// seg.LoadDict("../data/dict/dictionary.txt")
	err := seg.LoadDict("jp")
	fmt.Println("load dictionary error: ", err)

	// fmt.Println(seg.Cut(*text, true))
	// fmt.Println(seg.CutAll(*text))
	// fmt.Println(seg.CutSearch(*text))

	// fmt.Println(seg.String(*text, true))
	// fmt.Println(seg.Slice(*text, true))

	// segments := seg.Segment([]byte(*text))
	// fmt.Println(gse.ToString(segments, true))

	// read from filesystem 01.srt
	text2, err := ioutil.ReadFile("01.srt")
	fmt.Println("file read err: ", err)

	// text2 := []byte("運命は神の考えるものだ。人間は人間らしく働けばそれで結構だ。")
	segs := seg.Segment(text2)
	log.Println(gse.ToString(segs))
	// print segs array
	for _, seg := range segs {
		fmt.Println(seg.Token().Pos(), seg.Token().Text(), seg.Token().Freq())
	}

}
