package main

import (
	"fmt"
	"os"

)

/*
 プログラムのエントリポイント
 */
func main() {

	//実行時パラメータ取得
	fmt.Println(os.Args)

	if len(os.Args) != 3 {
		//実行時パラメータ数エラー
		fmt.Println("Input Argument count Error.")
		usage()
		os.Exit(1)
	}

	////デバッグ出力
	//fmt.Printf("Execute File Name: %s\n", os.Args[0])
	//fmt.Printf("Arg1: %s\n", os.Args[1])
	//fmt.Printf("Arg2: %s\n", os.Args[2])

	if len(os.Args[2]) <= 0 {
		//出力ファイル名異常
		fmt.Println("Input Argument Output Filename Error.")
		usage()
		os.Exit(1)
	}

	if os.Args[1] == "small" {
		//小分類処理
		SmallClassification(os.Args[2])

	} else if os.Args[1] == "middle" {
		//中分類処理
		MiddleClassification(os.Args[2])

	} else if os.Args[1] == "large" {
		//大分類処理
		//未実装
		fmt.Println("Input Argument mode large is Unimplemented.")
		usage()
		os.Exit(1)

	} else {
		//引数エラー
		fmt.Println("Input Argument mode Error.")
		usage()
		os.Exit(1)

	}


}

/*
 使い方表示
 */
func usage() {

	fmt.Println("")
	fmt.Println("*** Usage ***")
	fmt.Println("[How to use]")
	fmt.Println("  QuestionConverter.exe [mode] [OutputFileName] < [InputFileName]")
	fmt.Println("[Example]")
	fmt.Println("  QuestionConverter.exe small Small.xml < small.csv")
	fmt.Println("[mode]")
	fmt.Println("  small, middle, large  (large is Unimplemented)")

}


