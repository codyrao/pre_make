/**
 * @Author root
 * @Description //TODO $
 * @Date $ $
 **/
package utils

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func ReadFileOneTime(path string) string {
	byte, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("read file error:" + err.Error())
		os.Exit(0)
	}
	return string(byte)
}

type CheckType struct {
	InitEmptyJudge bool
	VarName        string
}

var checkType CheckType

func A() {

}
func (checkType CheckType) A() {

}
func ReadFileRow(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	row := 0
	scanner.Scan()
	header := scanner.Text()
	if !strings.HasPrefix(header, "// Package") {
		fmt.Printf("文件%s没有加头注释\n\n", path)
	}
	lastLine := ""
	for scanner.Scan() {

		row++
		line := scanner.Text()
		if checkType.InitEmptyJudge {
			if IsInitEmpty(line) {
				fmt.Printf("文件%s在第%d行的变量%s未判空\n\n", path, row-1, checkType.VarName)
			}
			checkType.InitEmptyJudge = false
		}

		if strings.Contains(line, ":= new(") || strings.Contains(line, ":= make(") || strings.Contains(line, ":= New") {
			checkType.InitEmptyJudge = true
			checkType.VarName = strings.Split(strings.Replace(line, "\t", "", -1), " :=")[0]

		}

		if strings.Contains(line, "%") {
			if IsFormatError(line) {
				fmt.Printf("文件%s在第%d行的格式化变量数不匹配\n\n", path, row)
			}
		}

		if strings.HasPrefix(line, "func") {
			if !strings.HasPrefix(lastLine, "//") {
				fmt.Printf("文件%s在第%d行的函数没有加注释\n\n", path, row)
			}
			var Info := strings.Split(line[6:], "(")[1]
			if strings.Contains(varInfo, "*") {
				checkType.InitEmptyJudge = true
				checkType.VarName = strings.Split(varInfo, " ")[0]
			}
		}
		lastLine = line

	}

}

func IsInitEmpty(line string) bool {

	if strings.Contains(line, "nil == "+checkType.VarName) || strings.Contains(line, checkType.VarName+" == nil") {

		return false
	}

	if strings.Contains(line, "nil != "+checkType.VarName) || strings.Contains(line, checkType.VarName+" != nil") {
		return false
	}

	return true
}

func IsFormatError(line string) bool {
	if strings.Count(line, ",") == strings.Count(line, "%") {
		return false
	}
	return true
}
