package constants

import "time"

const (
	CHUNK_SIZE             = 64 * 1024 * 1024 // 64mb
	DEFAULT_REPLICA_FACTOR = 3
	HEARTBEAT_INTERVAL     = 30 * time.Second
	SERVER_TIMEOUT         = 90 * time.Second
	MAX_RETRIES            = 3
	LEASE_TIMEOUT          = 60 * time.Second
	CHECKSUM_SIZE          = 4 // 4 bytes
)
