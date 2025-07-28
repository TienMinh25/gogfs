package master

import (
	"time"

	"github.com/TienMinh25/gogfs/gfstype"
)

type FileMetadata struct {
	Name      string
	Size      int64
	Chunks    []gfstype.ChunkHandle
	IsDeleted bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
