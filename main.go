package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

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
	// err := seg.LoadDict("jp")
	seg.LoadDict("jp")
	// fmt.Println("load dictionary error: ", err)

	// err = seg.LoadDict("./custom-dict.txt")
	seg.LoadDict("./custom-dict.txt")
	// fmt.Println("load dictionary error: ", err)

	// fmt.Println(seg.Cut(*text, true))
	// fmt.Println(seg.CutAll(*text))
	// fmt.Println(seg.CutSearch(*text))

	// fmt.Println(seg.String(*text, true))
	// fmt.Println(seg.Slice(*text, true))

	// segments := seg.Segment([]byte(*text))
	// fmt.Println(gse.ToString(segments, true))

	// read from filesystem 01.srt
	// text2, err := ioutil.ReadFile("01.srt")
	// text2, err := ioutil.ReadFile("kokoro1.txt")
	// fmt.Println("file read err: ", err)

	// loop at filecontents per line
	// split bytes string into lines at newline
	// texts := bytes.Split(text2, []byte("\n"))

	// read texts from stdin
	// var texts []byte
	// fmt.Scan(&texts)
	// read all texts from stdin
	texts, _ := ioutil.ReadAll(os.Stdin)
	// fmt.Println("std read", err)

	// stdin to string array
	texts2 := strings.Split(string(texts), "\n")
	// fmt.Println(texts2)

	for _, text := range texts2 {
		text := []byte(text)
		segs := seg.Segment(text)
		// print segs array
		for _, seg := range segs {
			// fmt.Println(seg.Token().Pos(), seg.Token().Text(), seg.Token().Freq())
			// fmt.Println("<---", seg.Token().Text())
			fmt.Println(seg.Token().Pos(), "\t", seg.Token().Text())
		}
	}

	// // text2 := []byte("運命は神の考えるものだ。人間は人間らしく働けばそれで結構だ。")
	// segs := seg.Segment(text2)
	// log.Println(gse.ToString(segs))
	// // print segs array
	// for _, seg := range segs {
	// 	fmt.Println(seg.Token().Text())
	// }
}
