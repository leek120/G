package main

import (
	"os"
	"strings"
)

func main() {

	RenameFiles("E:/ws/gws/", "新建文本", "")
}

func RenameFiles(root, oldStr, newStr string) {
	f, err := os.Open(root)
	if err != nil {
		println("打开目录 " + root + " 失败!")
		return
	}
	defer f.Close()

	if ds, _ := f.Stat(); !ds.IsDirectory() {
		println("路径 " + root + " 不是目录!")
		return
	}

	os.Chdir(root) // 切换根路径!

	for {
		dirs, err1 := f.Readdir(100)
		if err1 != nil || len(dirs) == 0 {
			break
		}
		for _, d := range dirs {
			if !d.IsDirectory() && strings.Contains(d.Name, oldStr) {
				newName := strings.Replace(d.Name, oldStr, newStr, 1)
				if err2 := os.Rename(d.Name, newName); err2 != nil {
					println("rename " + d.Name + " to " + newName + " failed: " + err2)
				} else {
					println("rename " + d.Name + " to " + newName + " ok")
				}
			}
		}
	}

	/* 
		//文件操作实例
		os.Chdir(root) // 改变根路径
		pwd, _ := os.Getwd()
		println(pwd)

		from, to := "renamefrom", "renameto"
		err = os.Rename(from, to) // 重命名
		if err != nil {
			println("rename %q, %q failed: %v", to, from, err)
		} else {
			println("rename %q to %q ok", to, from)
		}

		os.Remove(to) //remove file

		file, err := os.Create(from) //create file
		if err != nil {
			println("open %q failed: %v", to, err)
		}

		if err = file.Close(); err != nil { // colse file
			println("close %q failed: %v", to, err)
		} 
		defer os.Remove(to)
		_, err = os.Stat(to) //get fileInfo
		if err != nil {
			println("stat %q failed: %v", to, err)
		}
	*/
}
