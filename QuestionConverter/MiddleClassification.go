package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
	"strconv"
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

	//行カウンタ
	lineCount := 1

	//ブレイクキー
	breakKey := ""

	//標準入力処理
	for stdinObj.Scan() {
		if err := stdinObj.Err(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		lineText := stdinObj.Text()

		sepData := strings.Split(lineText, ",")

		if len(sepData) <= 0 || len(sepData) > 2 {
			fmt.Println("File format error. Line:" +  strconv.Itoa(lineCount))
			return
		}

		middleTitle := sepData[0]
		middleScript := sepData[1]

		if middleTitle != breakKey {
			//条件がブレイクしたので中分類タイトルを出力
			writeDetail_MiddleClassification_Title(file, err, middleTitle, lineCount)

			//ブレイクキーを保存
			breakKey = middleTitle
		}


		//中分類 明細行出力
		writeDetail_MiddleClassification_Script(file, err, middleScript, lineCount)
		lineCount++

	}

	//フッタ出力
	writeFooter_MiddleClassification(file, err)


}

/*
ヘッダ出力処理
 */
func writeHeader_MiddleClassification( fileObj *os.File, err error) {

	headerStr :=
`<?xml version="1.0" encoding="UTF-8" ?>
<FEEDBACK VERSION="200701" COMMENT="XML-Importfile for mod/feedback">
    <ITEMS>
    <ITEM TYPE="label" REQUIRED="0">
        <ITEMID>
            <![CDATA[324]]>
        </ITEMID>
        <ITEMTEXT>
            <![CDATA[]]>
        </ITEMTEXT>
        <ITEMLABEL>
            <![CDATA[]]>
        </ITEMLABEL>
        <PRESENTATION>
            <![CDATA[<ul><li>レベル0: 知識、経験なし</li><li>レベル1: トレーニングを受けた程度の知識あり</li><li>レベル2: サポートがあれば実施できる</li><li>レベル3 :独力で実施できる、またはその経験あり</li><li>レベル4 :他者を指導できる、またはその経験あり</li></ul><hr>]]>
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

	fileObj.WriteString(headerStr)

}

/*
中分類 タイトル出力
 */
func writeDetail_MiddleClassification_Title(fileObj *os.File, err error, middleTitle string, lineCount int) {

	detail_line := "            <![CDATA[" + strconv.Itoa(lineCount) + "]]>\n"

	detail_title := "            <![CDATA[<h3><b>" + middleTitle + "</b></h3>]]>\n"


	detail_foot :=
`        <OPTIONS>
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

	fileObj.WriteString("    <ITEM TYPE=\"numeric\" REQUIRED=\"0\">\n")
	fileObj.WriteString("        <ITEMID>\n")
	fileObj.WriteString(detail_line)
	fileObj.WriteString("        </ITEMID>\n")
	fileObj.WriteString("        <ITEMTEXT>\n")
	fileObj.WriteString("            <![CDATA[]]>\n")
	fileObj.WriteString("        </ITEMTEXT>\n")
	fileObj.WriteString("        <ITEMLABEL>\n")
	fileObj.WriteString("            <![CDATA[]]>\n")
	fileObj.WriteString("        </ITEMLABEL>\n")
	fileObj.WriteString("        <PRESENTATION>\n")
	fileObj.WriteString(detail_title)
	fileObj.WriteString("        </PRESENTATION>\n")
	fileObj.WriteString(detail_foot)

}

/*
中分類 明細出力
 */
func writeDetail_MiddleClassification_Script(fileObj *os.File, err error, middleScript string, lineCount int) {


	detail_line := "            <![CDATA[" + strconv.Itoa(lineCount) + "]]>\n"

	detail_script := "            <![CDATA[<div class=\"editor-indent\" style=\"margin-left:30px;\"><h3><ul><li>" + middleScript + "<br></li></ul></h3></div>]]>\n"


	detail_foot :=
`        <OPTIONS>
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


	fileObj.WriteString("    <ITEM TYPE=\"numeric\" REQUIRED=\"0\">\n")
	fileObj.WriteString("        <ITEMID>\n")
	fileObj.WriteString(detail_line)
	fileObj.WriteString("        </ITEMID>\n")
	fileObj.WriteString("        <ITEMTEXT>\n")
	fileObj.WriteString("            <![CDATA[]]>\n")
	fileObj.WriteString("        </ITEMTEXT>\n")
	fileObj.WriteString("        <ITEMLABEL>\n")
	fileObj.WriteString("            <![CDATA[]]>\n")
	fileObj.WriteString("        </ITEMLABEL>\n")
	fileObj.WriteString("        <PRESENTATION>\n")
	fileObj.WriteString(detail_script)
	fileObj.WriteString("        </PRESENTATION>\n")
	fileObj.WriteString(detail_foot)

}

/*
フッタ出力
 */
func writeFooter_MiddleClassification(fileObj *os.File, err error) {

	footStr :=
`    </ITEMS>
</FEEDBACK>
`

	fileObj.WriteString(footStr)
}