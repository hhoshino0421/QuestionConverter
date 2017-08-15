package main

import (
	"bufio"
	"os"
	"fmt"
)

/*
中分類カテゴリー出力
 */
func MiddleClassification(filename string) {

	//標準入力
	stdinObj := bufio.NewScanner(os.Stdin)

	//出力ファイルオープン
	file, err := os.Create(filename)
	if err != nil {
		// Openエラー処理
	}
	defer file.Close()

	//ヘッダ出力
	writeHeader_MiddleClassification(file, err)

	lineCount := 1

	//標準入力処理
	for stdinObj.Scan() {
		if err := stdinObj.Err(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		//lineText := stdinObj.Text()

		//明細行出力
		//writeDetail_MiddleClassification(file, err, lineText, lineCount)
		lineCount++

	}

	//フッタ出力
	//writeFooter_MiddleClassification(file, err)


}

/*
ヘッダ出力処理
 */
func writeHeader_MiddleClassification( fileObj *os.File, err error) {

	headerStr :=
		`<?xml version=\"1.0\" encoding=\"UTF-8\" ?>
		<FEEDBACK VERSION="200701" COMMENT="XML-Importfile for mod/feedback">
			<ITEMS>
		`

	fileObj.WriteString(headerStr)

}
