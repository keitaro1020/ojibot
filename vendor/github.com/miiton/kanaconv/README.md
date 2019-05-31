# kanaconv [![Build Status](https://travis-ci.org/miiton/kanaconv.svg?branch=master)](https://travis-ci.org/miiton/kanaconv) [![GoDoc](https://godoc.org/github.com/miiton/kanaconv?status.svg)](https://godoc.org/github.com/miiton/kanaconv)
日本語のひらがなカタカナ、全角半角を変換するパッケージ written in Go

## Installation

```
go get -u github.com/miiton/kanaconv
```

## Usage

```
import (
    "fmt"
    "github.com/miiton/kanaconv"
)

func main() {
    fmt.Println(kanaconv.HiraganaToKatakana("ひらがなをカタカナに"))
    // ヒラガナヲカタカナニ
    fmt.Println(kanaconv.KatakanaToHiragana("カタカナをひらがなに"))
    // かたかなをひらがなに
    fmt.Println(kanaconv.HankakuToZenkaku("ﾊﾝｶｸｦゼンカクﾆ"))
    // ハンカクヲゼンカクニ
    fmt.Println(kanaconv.ZenkakuToHankaku("ゼンカクヲﾊﾝｶｸニ"))
    // ｾﾞﾝｶｸｦﾊﾝｶｸﾆ
	fmt.Println(kanaconv.SmartConv("ｶﾀｶﾅは全角に統一ＥＩＳＵＵＪＩ＋＝−，記号は半角に統一します"))
    // カタカナは全角に統一EISUUJI+=-,記号は半角に統一します
}
```

## Benchmarks

* MacBook Air (11-inch, Mid 2011)

```
% go test -bench .
PASS
BenchmarkHiraganaToKatakana-4	 1000000	      1657 ns/op
BenchmarkKatakanaToHiragana-4	 1000000	      1597 ns/op
BenchmarkHankakuToZenkaku-4  	 1000000	      1995 ns/op
BenchmarkZenkakuToHankaku-4  	 1000000	      2008 ns/op
BenchmarkSmartConv-4         	 1000000	      2063 ns/op
ok  	github.com/miiton/kanaconv	9.450s
```

## Referenced

* [Go言語で文字列の変換(全角・半角、ひらがな・カタカナ)をする : Serendip - Webデザイン・プログラミング](http://www.serendip.ws/archives/6307)
