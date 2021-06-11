package main

import (
	"fmt"
	"time"

	"github.com/karalabe/hid"
)

func main() {
	devices := hid.Enumerate(0x1038, 0x1290)
	for i := range devices {
		di := devices[i]
		if di.Interface != 0 {
			continue
		}

		fmt.Printf("%#v\n", di)
		d, err := di.Open()
		if err != nil {
			fmt.Printf("%#v\n", err)
			continue
		}

		for {
			for i := 0x40; i <= 0x42; i++ {
				if _, err := d.Write([]byte{byte(i), 0xAA}); err != nil {
					fmt.Printf("%#v\n", err)
					continue
				}

				var r = make([]byte, 4)
				if _, err := d.Read(r); err != nil {
					fmt.Printf("%#v\n", err)
					continue
				}

				fmt.Printf("%x: %#v\n", i, r)
			}
			time.Sleep(time.Second)
		}
	}
}
