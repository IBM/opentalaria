package postgresql

import "github.com/google/uuid"

type PartitionsModel struct {
	PartitionId   uuid.UUID
	TopicId       uuid.UUID
	CurrentOffset int
	PartitionIx   int
}
