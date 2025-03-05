package postgresql

import (
	"log/slog"

	"github.com/google/uuid"
	"github.com/ibm/opentalaria/protocol"
	"github.com/ibm/opentalaria/utils"
	"github.com/lib/pq"
)

func (p *Plugin) AddTopic(topic protocol.CreatableTopic) utils.KError {
	slog.Debug("Create", "topic", topic.Name, "configs", topic.Configs)

	statement := `
		INSERT INTO topics (topic_id, topic_name, num_partitions, replication_factor)
		VALUES ($1, $2, $3, $4)`

	_, err := p.db.Exec(statement, uuid.New(), topic.Name, topic.NumPartitions, topic.ReplicationFactor)

	returnErr := utils.ErrNoError

	if err, ok := err.(*pq.Error); ok {
		// 23505 is a unique constraint violation. In our case the topic already exists
		// https://www.postgresql.org/docs/9.3/errcodes-appendix.html
		if err.Code == "23505" {
			returnErr = utils.ErrTopicAlreadyExists
		} else {
			returnErr = utils.ErrInvalidRequest
		}
	}

	return returnErr
}
