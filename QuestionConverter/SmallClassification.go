package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

/*
小分類カテゴリー出力
 */
func SmallClassification(filename string) {

	//標準入力
	stdinObj := bufio.NewScanner(os.Stdin)

	//出力ファイルオープン
	file, err := os.Create(filename)
	if err != nil {
		// Openエラー処理
	}
	defer file.Close()


	//ヘッダ出力
	writeHeader_SmallClassification(file, err)

	lineCount := 1

	//標準入力処理
	for stdinObj.Scan() {
		if err := stdinObj.Err(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		lineText := stdinObj.Text()

		//明細行出力
		writeDetail_SmallClassification(file, err, lineText, lineCount)
		lineCount++

	}

	//フッタ出力
	writeFooter_SmallClassification(file, err)

}

/*
明細出力処理
 */
func writeDetail_SmallClassification(fileObj *os.File, err error, lineText string, lineCount int) {

	sepData := strings.Split(lineText, ",")

	if len(sepData) <= 0 {
		fmt.Println("File format error. Line:" +  strconv.Itoa(lineCount))
		return
	}

	detail_line := "				<![CDATA[" + strconv.Itoa(lineCount) + "]]>\n"

	detail_foot :=
	`			 <PRESENTATION>
                    <![CDATA[0|4]]>
               </PRESENTATION>
               <OPTIONS>
                    <![CDATA[]]>
               </OPTIONS>
               <DEPENDITEM>
                    <![CDATA[0]]>
               </DEPENDITEM>
               <DEPENDVALUE>
                    <![CDATA[]]>
               </DEPENDVALUE>
          </ITEM>
	`

	fileObj.WriteString("	<ITEM TYPE=\"numeric\" REQUIRED=\"0\">\n")
	fileObj.WriteString("	   <ITEMID>\n")
	fileObj.WriteString(detail_line)
	fileObj.WriteString("	   </ITEMID>\n")
	fileObj.WriteString("	   <ITEMTEXT>\n")
	fileObj.WriteString("     		<![CDATA[" + sepData[0] + "]]>\n")
	fileObj.WriteString("	   </ITEMTEXT>\n")
	fileObj.WriteString("	   <ITEMLABEL>\n")
	fileObj.WriteString("	       <![CDATA[" + sepData[1] + "]]>\n")
	fileObj.WriteString("	   </ITEMLABEL>\n")
	fileObj.WriteString(detail_foot)


}

/*
ヘッダ出力処理
 */
func writeHeader_SmallClassification( fileObj *os.File, err error) {

	headerStr :=
	`<?xml version="1.0" encoding="UTF-8" ?>
	<FEEDBACK VERSION="200701" COMMENT="XML-Importfile for mod/feedback">
		<ITEMS>
	`

	fileObj.WriteString(headerStr)

}

/*
フッタ出力処理
 */
func writeFooter_SmallClassification(fileObj *os.File, err error) {

	footStr :=
	`
		</ITEMS>
	</FEEDBACK>
	`

	fileObj.WriteString(footStr)
}

