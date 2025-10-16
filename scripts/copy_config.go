package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	// Исходный файл
	src := "configs/default.toml"

	// Целевая папка
	binDir := "bin"

	// Создаем папку bin если ее нет
	if err := os.MkdirAll(binDir, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Error creating bin directory: %v\n", err)
		os.Exit(1)
	}

	// Целевой файл
	dst := filepath.Join(binDir, "default.toml")

	// Читаем исходный файл
	srcFile, err := os.Open(src)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening source file: %v\n", err)
		os.Exit(1)
	}
	defer srcFile.Close()

	// Создаем целевой файл
	dstFile, err := os.Create(dst)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating destination file: %v\n", err)
		os.Exit(1)
	}
	defer dstFile.Close()

	// Копируем содержимое
	if _, err := io.Copy(dstFile, srcFile); err != nil {
		fmt.Fprintf(os.Stderr, "Error copying file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Configuration copied to %s\n", dst)
}
