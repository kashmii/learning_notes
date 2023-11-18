package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// コマンドライン引数からファイル名を取得
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Usage: go run main.go <filename>")
		return
	}
	filename := args[1]

	// 1ファイルあたりの行数
	linesPerFile := 1000

	// =========
	// ファイルを開く
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// 分割されたファイルのカウンタ
	fileCounter := 1

	// 新しいファイルを作成
	newFile, err := os.Create(fmt.Sprintf("split_%d.txt", fileCounter))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer newFile.Close()

	// ファイルから読み込み
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// 新しいファイルに行を書き込み
		_, err := newFile.WriteString(line + "\n")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		// 指定行数ごとに新しいファイルを作成
		if scanner.LineCount()%linesPerFile == 0 {
			fileCounter++
			newFile.Close()
			newFile, err = os.Create(fmt.Sprintf("split_%d.txt", fileCounter))
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			defer newFile.Close()
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}
}
