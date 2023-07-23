package main

import (
	"bytes"
	"github.com/38888/go-storage/local"
)

func main() {
	// 初始化文件存储系统
	disk, _ := local.Init(local.Config{
		RootDir: "./tests",
		AppUrl:  "http://localhost:8888/tests",
	})

	// 保存文件
	buf := bytes.NewBuffer([]byte{'A', 'B', 'C', 'D', 'E'})
	disk.Put("local/accounts.txt", buf, int64(buf.Len()))

	// 保存本地文件
	disk.PutFile("path/to/file.ext", "local/path/to/local_file.ext")

	// 获取文件内容
	disk.Get("path/to/file.ext")

	// 文件重命名
	disk.Rename("path/to/file1.ext", "path/to/file2.ext")

	// 移动文件
	disk.Copy("path/to/file1.ext", "path/to/file2.ext")

	// 文件是否存在
	disk.Exists("path/to/file.ext")

	// 获取文件大小（字节）
	disk.Size("path/to/file.ext")

	// 删除文件
	disk.Delete("path/to/file.ext")

	// 获取文件URL
	disk.Url("path/to/file.ext")

}
