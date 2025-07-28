package master

import (
	"time"

	"github.com/TienMinh25/gogfs/gfstype"
)

type ChunkInfo struct {
	Handle        gfstype.ChunkHandle
	VersionNumber int64
	Size          int64
	Checksum      string
	Locations     []gfstype.ServerID
	Primary       gfstype.ServerID
	LeaseExpiry   time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
