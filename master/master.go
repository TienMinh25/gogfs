package master

/*
* Master must manage some metadata, include:
  - mapping filename 				-> many chunks (chunk handle) 	(non-volatile) -> save to disk
  - chunk handle (chunk id) -> struct
  - which chunkserver replicas		(volatile)		 -> not save to disk
  - version number								(non-volatile) -> used to mechanism to check stale replica
  - primary replica
  - lease expire time
*/
type Master struct {
	FileMapping  map[string]*FileMetadata
	ChunkHandles map[string]*ChunkInfo
}

func NewMaster() *Master {
	return &Master{
		FileMapping:  make(map[string]*FileMetadata),
		ChunkHandles: make(map[string]*ChunkInfo),
	}
}
