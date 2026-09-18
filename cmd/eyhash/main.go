package main

import (
	"fmt"
	"os"
	"path"
	"slices"

	"github.com/garudaproject/eyhash"
)

var version = "dev"

const banner = `
              __               __
  ___  __  __/ /_  ____ ______/ /_
 / _ \/ / / / __ \/ __ '/ ___/ __ \
/  __/ /_/ / / / / /_/ (__  ) / / /
\___/\__, /_/ /_/\__,_/____/_/ /_/
    /____/ %s
`

func usage() {
	fmt.Print(fmt.Sprintf(banner, version), `
Usage: eyehash <filename/directory> -f/-d [file/directory]
`)
}

func main() {
	var file, mode string
	for i, arg := range os.Args {
		if slices.Contains([]string{"-h", "--help"}, arg) {
			usage()
			return
		}
		switch i {
		case 1:
			file = arg
		case 2:
			mode = arg
		}
	}
	if file == "" || mode == "" {
		usage()
		return
	}
	switch mode {
	case "-f":
		info, err := eyhash.FileInfo(file)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println()
		hashInfo(file, info)
	case "-d":
		dirs, err := os.ReadDir(file)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		for _, dir := range dirs {
			filePath := path.Join(file, dir.Name())
			info, err := eyhash.FileInfo(filePath)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Println()
			if err := hashInfo(filePath, info); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	default:
		usage()
	}
}

func hashInfo(file string, info *eyhash.Info) error {
	md5hash, err := eyhash.MD5File(file)
	if err != nil {
		return err
	}
	sh1hash, err := eyhash.SHA1File(file)
	if err != nil {
		return err
	}
	sh256hash, err := eyhash.SHA256File(file)
	if err != nil {
		return err
	}
	sh512hash, err := eyhash.SHA512File(file)
	if err != nil {
		return err
	}
	fmt.Printf("FileName: %s, Size: %d bytes, ModTime: %s\n", info.Name, info.Size, info.ModTime.Local())
	fmt.Println("MD5:", md5hash)
	fmt.Println("SHA1:", sh1hash)
	fmt.Println("SHA256:", sh256hash)
	fmt.Println("SHA512:", sh512hash)
	return nil
}
