package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	filename := "file.txt"

	// ファイルオープン
	fp, err := os.Open(filename)
	if err != nil {
		// エラー処理
		log.Fatalf("Error when opening file: %s", err)
		fmt.Println("ファイルがないんじゃーーーい🐶")
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)

	for scanner.Scan() {
		// ここで一行ずつ処理
		fmt.Println("おじさんは" + scanner.Text() + "が大好きだぞ")
	}

	if err = scanner.Err(); err != nil {
		// エラー処理
		log.Fatalf("Error while reading file: %s", err)
	}
}
