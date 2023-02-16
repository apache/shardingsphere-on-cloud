package model

type (
	Backup struct {
		ID                string `json:"id"`
		BackupMode        string `json:"backup-mode"`
		Wal               string `json:"wal"`
		CompressAlg       string `json:"compress-alg"`
		CompressLevel     int    `json:"compress-level"`
		FromReplica       string `json:"from-replica"`
		BlockSize         int    `json:"block-size"`
		XlogBlockSize     int    `json:"xlog-block-size"`
		ChecksumVersion   int    `json:"checksum-version"`
		ProgramVersion    string `json:"program-version"`
		ServerVersion     string `json:"server-version"`
		CurrentTli        int    `json:"current-tli"`
		ParentTli         int    `json:"parent-tli"`
		StartLsn          string `json:"start-lsn"`
		StopLsn           string `json:"stop-lsn"`
		StartTime         string `json:"start-time"`
		EndTime           string `json:"end-time"`
		RecoveryXid       int    `json:"recovery-xid"`
		RecoveryTime      string `json:"recovery-time"`
		RecoveryName      string `json:"recovery-name"`
		DataBytes         int    `json:"data-bytes"`
		WalBytes          int    `json:"wal-bytes"`
		UncompressedBytes int    `json:"uncompressed-bytes"`
		PgdataBytes       int    `json:"pgdata-bytes"`
		Status            string `json:"status"`
		ContentCrc        int64  `json:"content-crc"`
	}

	BackupList struct {
		Instance string   `json:"instance"`
		List     []Backup `json:"backups"`
	}
)
