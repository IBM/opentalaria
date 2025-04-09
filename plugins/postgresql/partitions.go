package postgresql

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

func (p *Plugin) CreatePartitions(topicName string, partitionCount int) error {
	topic, err := p.GetTopic(topicName)
	if err != nil {
		return err
	}

	if partitionCount <= len(topic.Partitions) {
		return errors.New("new partition count cannot be lower than existing partition count")
	}

	for i := range partitionCount - len(topic.Partitions) {
		partitionIx := i + len(topic.Partitions) + 1

		statement := `INSERT INTO partitions (partition_id, topic_id, current_offset, partition_ix)
					  VALUES($1, $2, $3, $4)`

		_, err := p.db.Exec(statement, uuid.New(), topic.TopicID.String(), 0, partitionIx)
		if err != nil {
			return fmt.Errorf("error creating partition: %v", err)
		}
	}

	statement := `UPDATE topics SET 
					num_partitions = $2 
					WHERE topic_name = $1`

	_, err = p.db.Exec(statement, topicName, partitionCount)

	return err
}

func (p *Plugin) deletePartitions(topicName string) error {
	topic, err := p.GetTopic(topicName)
	if err != nil {
		return err
	}

	statement := `DELETE FROM partitions WHERE topic_id = $1`
	_, err = p.db.Exec(statement, topic.TopicID.String())

	return err
}
