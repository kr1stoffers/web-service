/*
3. Копирование с буфером: Реализуйте свою версию io.Copy, используя буфер
фиксированного размера (например, 32 байта) и цикл.
*/
package main

import (
	"fmt"
	"io"
	"os"
)

func customCopy(dst io.Writer, src io.Reader) (int64, error) {
	buf := make([]byte, 32)
	var written int64

	for {
		nr, rerr := src.Read(buf)
		if nr > 0 {
			nw, werr := dst.Write(buf[:nr])
			if nw > 0 {
				written += int64(nw)
			}
			if werr != nil {
				return written, werr
			}
		}
		if rerr == io.EOF {
			break
		}
		if rerr != nil {
			return written, rerr
		}
	}
	return written, nil
}

func main() {
	srcFile, _ := os.Open("numbers.txt")
	defer srcFile.Close()

	dstFile, _ := os.Create("copy_numbers.txt")
	defer dstFile.Close()

	bytesWritten, err := customCopy(dstFile, srcFile)
	if err != nil {
		fmt.Println("Ошибка копирования:", err)
		return
	}
	fmt.Printf("Скопировано байт: %d\n", bytesWritten)
}
