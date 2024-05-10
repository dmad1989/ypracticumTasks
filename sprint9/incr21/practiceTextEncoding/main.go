package main

import (
	"bytes"
	"fmt"
	"io"
	"log"

	"golang.org/x/text/encoding/ianaindex"
	"golang.org/x/text/transform"
)

func ConvertToUTF8(source []byte, coding string) (string, error) {
	// выбираем кодировку из каталога IANA
	e, err := ianaindex.IANA.Encoding(coding)
	if err != nil {
		return ``, err
	}
	// конструируем кодировщик transform.Reader,
	// способный читать и трансформировать байты на лету,
	// для этого понадобятся любой io.Reader (здесь делаем из слайса байт)
	// и кодировщик для трансформации байт (берём Decoder выбранной из IANA кодировки)
	toutf8 := transform.NewReader(bytes.NewReader(source), e.NewDecoder())
	// читаем всё и перекодируем на лету
	decBytes, err := io.ReadAll(toutf8)
	if err != nil {
		return ``, err
	}
	return string(decBytes), nil
}

func main() {
	dest, err := ConvertToUTF8([]byte{207, 240, 232, 226, 229, 242}, `windows-1251`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dest)
}
