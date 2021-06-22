package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func Copy() {
	var src string
	var dst string
	var yes bool
	var help bool

	flag.StringVar(&src, "s", "", "Please input the copy target file!	")
	flag.StringVar(&dst, "d", "", "Please input the copy target file!	")
	flag.BoolVar(&yes, "y", false, "Whether to cover (y means to cover)!")
	flag.BoolVar(&help, "h", false, "This is a help:")

	flag.Usage = func() {
		fmt.Println("-s srcfile -d dstfile")
		flag.PrintDefaults()
	}

	flag.Parse()

	if help {
		flag.Usage()
		return
	}

	srcfile, err := os.Open(src)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer srcfile.Close()

	readStr, err := ioutil.ReadAll(srcfile)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = os.Stat(dst)
	if err != nil {
		dstfile, _ := os.OpenFile(dst, os.O_CREATE|os.O_RDWR, 0666)
		writer := bufio.NewWriter(dstfile)
		writer.WriteString(string(readStr))
		writer.Flush()
	} else {
		cover := ""
		fmt.Println("覆盖请输入(y):")
		fmt.Scanln(&cover)
		if cover == "y" {
			fmt.Println("覆盖中")
			dfile, _ := os.OpenFile(dst, os.O_TRUNC|os.O_RDWR, 0666)
			writer := bufio.NewWriter(dfile)
			writer.WriteString(string(readStr))
			writer.Flush()
		}
	}

}

func main() {
	Copy()
}
