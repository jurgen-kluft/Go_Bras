package fileinfo

import (
	"os"
	"syscall"
	"time"
)

const (
	flagsNetBSD = flagHasCTime | FlagHasBTime
)

func timespecToTime(ts syscall.Timespec) time.Time {
	return time.Unix(int64(ts.Sec), int64(ts.Nsec))
}

func getTimespec(fi os.FileInfo) (t Times) {
	stat := fi.Sys().(*syscall.Stat_t)
	t.flags = flagsNetBSD
	t.atime = timespecToTime(stat.Atimespec)
	t.mtime = timespecToTime(stat.Mtimespec)
	t.ctime = timespecToTime(stat.Ctimespec)
	t.btime = timespecToTime(stat.Birthtimespec)
	return t
}
