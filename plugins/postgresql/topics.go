package postgresql

import (
	"database/sql"
	"log/slog"
	"time"

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

	if err1, ok := err.(*pq.Error); ok {
		// 23505 is a unique constraint violation. In our case it means the topic already exists
		// https://www.postgresql.org/docs/9.3/errcodes-appendix.html
		if err1.Code == "23505" {
			returnErr = utils.ErrTopicAlreadyExists
		} else {
			slog.Error("error creating topic", "err", err)
			returnErr = utils.ErrInvalidRequest
		}
	}

	return returnErr
}

func (p *Plugin) DeleteTopic(topic string) utils.KError {
	slog.Debug("delete topic", "name", topic)

	statement := `
		DELETE FROM topics
		where topic_name = $1
	`
	returnErr := utils.ErrNoError

	_, err := p.db.Exec(statement, topic)
	if err != nil {
		slog.Error("error deleting topic", "err", err)
		returnErr = utils.ErrInvalidRequest
	}

	return returnErr
}

func (p *Plugin) ListTopics(topicName []string) ([]protocol.MetadataResponseTopic, error) {

	var rows *sql.Rows
	var err error

	query := "SELECT * from topics"

	if topicName != nil {
		query += " WHERE topic_name = ANY($1)"

		stmt, err := p.db.Prepare(query)
		if err != nil {
			return nil, err
		}

		rows, err = stmt.Query(pq.Array(topicName))
	} else {
		rows, err = p.db.Query(query)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]protocol.MetadataResponseTopic, 0)

	for rows.Next() {
		var topic_id, topic_name string
		var num_partitions, replication_factor int
		var is_internal bool

		if err := rows.Scan(&topic_id, &topic_name, &num_partitions,
			&replication_factor, &is_internal); err != nil {
			return result, err
		}

		kErr := utils.ErrNoError
		tId, err := uuid.Parse(topic_id)
		if err != nil {
			slog.Error("error parsing topic id", "err", err)

			kErr = utils.ErrInvalidTopic
		}

		// TODO: partitions are currently mocked
		partitions := make([]protocol.MetadataResponsePartition, num_partitions)
		for i := 0; i < num_partitions; i++ {
			partitions[i].ErrorCode = int16(utils.ErrNoError)
			partitions[i].PartitionIndex = int32(i)
			partitions[i].LeaderID = 1
			partitions[i].LeaderEpoch = int32(time.Now().Unix())
			// Replicas and isr are currently mocked. We just return whatever value was set when creating the topic.
			partitions[i].ReplicaNodes = []int32{int32(replication_factor)}
			partitions[i].IsrNodes = []int32{int32(replication_factor)}
			partitions[i].OfflineReplicas = []int32{0}
		}

		// TODO: Implement topic authorized operations
		result = append(result, protocol.MetadataResponseTopic{
			ErrorCode:                 int16(kErr),
			Name:                      &topic_name,
			TopicID:                   tId,
			IsInternal:                is_internal,
			Partitions:                partitions,
			TopicAuthorizedOperations: 0,
		})

	}

	return result, nil
}
