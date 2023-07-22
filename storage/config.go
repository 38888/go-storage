package storage

type DiskName string

const (
	Local DiskName = "local" // 本地
	KoDo  DiskName = "qiniu" // 七牛云
	Oss   DiskName = "oss"   // 阿里云
)
