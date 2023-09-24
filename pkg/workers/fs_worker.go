package workers

import (
	"syscall"
)

type FSWorker struct {
	Volume string
	fs     syscall.Statfs_t
}

type driveInfo struct {
	Type      int64
	TotalSize uint64
	FreeSpace uint64
	Volume    string // ???
}

func NewFSWorker(volume string) *FSWorker {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(volume, &fs)

	if err != nil {
		return nil
	}

	return &FSWorker{
		Volume: volume,
	}
}

func (w *FSWorker) DriveInfo() driveInfo {
	drive := driveInfo{}
	drive.Type = w.fs.Type
	drive.TotalSize = w.fs.Blocks * uint64(w.fs.Bsize)
	drive.FreeSpace = w.fs.Bfree * uint64(w.fs.Bsize)

	return drive
}
