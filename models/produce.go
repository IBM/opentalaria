package models

import "github.com/ibm/opentalaria/utils"

type ProducePartitionResponse struct {
	PartitionIndex int32
	BaseOffset     int
	Error          utils.KError
}
