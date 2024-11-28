package main

import (
	"fmt"
	raw "github.com/MRHT-SRProject/LibRawGo/librawgo"
	"os"
)

func getRawThumbnailFromFile(raw_pointer raw.Libraw_data_t, file_path string) raw.Libraw_thumbnail_t {
	raw.Libraw_open_file(raw_pointer, file_path)
	thumbnail := raw_pointer.GetThumbnail()
	return thumbnail
}

func main() {
	lr := raw.Libraw_init(0)
	thumbnail := getRawThumbnailFromFile(lr, "/home/danmac/Pictures/Darktable/20240820_no_name/20240820_0007.CR3")
	output, err := os.Create("test_thumbnail.jpg")
	if err != nil {
		panic(err)
	}

	defer func() {
		err := output.Close()
		if err != nil {
			panic(err)
		}
	}()

	if _, err := output.Write([]byte(thumbnail.GetThumb())); err != nil {
		panic(err)
	}

	fmt.Println("Thumbnail saved.")
}
