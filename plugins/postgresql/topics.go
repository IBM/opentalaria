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
		INSERT INTO topics (topic_id, topic_name, replication_factor)
		VALUES ($1, $2, $3)`

	_, err := p.db.Exec(statement, uuid.New(), topic.Name, topic.ReplicationFactor)

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

	err = p.CreatePartitions(topic.Name, int(topic.NumPartitions))
	if err != nil {
		slog.Error("error creating partitions", "err", err)
		returnErr = utils.ErrInvalidRequest
	}

	return returnErr
}

func (p *Plugin) DeleteTopic(topic string) utils.KError {
	err := p.deletePartitions(topic)
	if err != nil {
		slog.Error("error deleting partitions for topic", "topic", topic, "err", err)
		return utils.ErrInvalidRequest
	}

	statement := `
		DELETE FROM topics
		where topic_name = $1
	`
	_, err = p.db.Exec(statement, topic)
	if err != nil {
		slog.Error("error deleting topic", "err", err)
		return utils.ErrInvalidRequest
	}

	return utils.ErrNoError
}

func (p *Plugin) ListTopics(topicName []string) ([]protocol.MetadataResponseTopic, error) {
	var rows *sql.Rows
	var err error

	query := "SELECT * from topics"

	if topicName != nil {
		query += " WHERE topic_name = ANY($1)"

		var stmt *sql.Stmt
		stmt, err = p.db.Prepare(query)
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

		result = append(result, getTopicMetadata(topic_id, topic_name, num_partitions, replication_factor, is_internal))
	}

	return result, nil
}

func (p *Plugin) GetTopic(topicName string) (protocol.MetadataResponseTopic, error) {
	query := "SELECT * FROM topics WHERE topic_name = $1"

	stmt, err := p.db.Prepare(query)
	if err != nil {
		return protocol.MetadataResponseTopic{}, err
	}

	row := stmt.QueryRow(topicName)

	var topic_id, topic_name string
	var num_partitions, replication_factor int
	var is_internal bool

	if err := row.Scan(&topic_id, &topic_name, &num_partitions,
		&replication_factor, &is_internal); err != nil {
		return protocol.MetadataResponseTopic{}, err
	}

	return getTopicMetadata(topic_id, topic_name, num_partitions, replication_factor, is_internal), nil
}

func getTopicMetadata(topic_id, topic_name string, num_partitions, replication_factor int, is_internal bool) protocol.MetadataResponseTopic {
	kErr := utils.ErrNoError
	tId, err := uuid.Parse(topic_id)
	if err != nil {
		slog.Error("error parsing topic id", "err", err)

		kErr = utils.ErrInvalidTopic
	}

	partitions := make([]protocol.MetadataResponsePartition, num_partitions)
	for i := range num_partitions {
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
	return protocol.MetadataResponseTopic{
		ErrorCode:                 int16(kErr),
		Name:                      &topic_name,
		TopicID:                   tId,
		IsInternal:                is_internal,
		Partitions:                partitions,
		TopicAuthorizedOperations: 0,
	}
}
