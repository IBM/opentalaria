package postgresql

import (
	"fmt"
	"log/slog"

	"github.com/google/uuid"
	"github.com/ibm/opentalaria/protocol"
	"github.com/ibm/opentalaria/utils"
)

func (p *Plugin) Produce(req protocol.ProduceRequest) (protocol.ProduceResponse, error) {
	response := protocol.ProduceResponse{}

	for _, topic := range req.TopicData {
		topicResponse := protocol.TopicProduceResponse{}
		topicResponse.Name = topic.Name
		errResponse := utils.ErrNoError

		topicObject, err := p.GetTopic(topic.Name)
		if err != nil {
			errResponse = utils.ErrInvalidRequest
		}

		for _, partition := range topic.PartitionData {
			slog.Debug("Received records", "records", fmt.Sprintf("%+v", partition.Records))

			newOffset, err := p.writeRecords(partition.Records, topicObject, int(partition.Index))
			if err != nil {
				slog.Error("error producing message", "err", err)
				errResponse = utils.ErrInvalidRequest
			}

			topicResponse.PartitionResponses = append(topicResponse.PartitionResponses, protocol.PartitionProduceResponse{
				Version:    response.Version,
				Index:      partition.Index,
				ErrorCode:  int16(errResponse),
				BaseOffset: int64(newOffset),
				// TODO: this needs to be implemented, see documentation for details
				LogAppendTimeMs: -1,
				LogStartOffset:  0,
				// TODO: Don't forget to handle errors when the protocol is fully implemented
			})
		}

		response.Responses = append(response.Responses, topicResponse)
	}

	return response, nil
}

func (p *Plugin) writeRecords(recordBatch protocol.RecordBatch, topic protocol.MetadataResponseTopic, partitionIx int) (int, error) {
	partition, err := p.getPartitionForTopic(*topic.Name, partitionIx)
	if err != nil {
		return 0, err
	}

	startingOffset := partition.CurrentOffset

	query := `INSERT INTO records (record_id,
	topic_id,
	current_offset,
	base_offset,
	batch_length,
	partition_leader_epoch,
	magic,
	crc,
	compression_type,
	timestamp_type,
	is_transactional,
	is_control_batch,
	has_delete_horizon_ms,
	last_offset_delta,
	base_timestamp,
	max_timestamp,
	producer_id,
	producer_epoch,
	base_sequence,
	records_len,
	records,
	partition_id)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22)`

	_, err = p.db.Exec(query,
		uuid.New(),
		topic.TopicID,
		startingOffset,
		recordBatch.BaseOffset,
		recordBatch.BatchLength,
		recordBatch.PartitionLeaderEpoch,
		recordBatch.Magic,
		recordBatch.CRC,
		recordBatch.CompressionType,
		recordBatch.TimestampType,
		recordBatch.IsTransactional,
		recordBatch.IsControlBatch,
		recordBatch.HasDeleteHorizonMs,
		recordBatch.LastOffsetDelta,
		recordBatch.BaseTimestamp,
		recordBatch.MaxTimestamp,
		recordBatch.ProducerId,
		recordBatch.ProducerEpoch,
		recordBatch.BaseSequence,
		recordBatch.RecordsLen,
		recordBatch.Records,
		partition.PartitionId,
	)

	// update the partition offset

	return startingOffset, err
}
