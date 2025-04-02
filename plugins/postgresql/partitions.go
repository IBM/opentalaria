package postgresql

func (p *Plugin) CreatePartitions(topicName string, partitionCount int) error {
	statement := `UPDATE topics SET 
					num_partitions = $2 
					WHERE topic_name = $1`

	_, err := p.db.Exec(statement, topicName, partitionCount)

	return err
}
