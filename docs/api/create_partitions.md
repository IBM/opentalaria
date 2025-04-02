# Create Partitions API

https://kafka.apache.org/protocol#The_Messages_CreatePartitions

### Caveats
For now, since OpenTalaria does not distribute data across partitions, the timeout_ms parameter in the request is ignored. We make an assumption that the operation is always instantaneous since all we do is update a field in the database. This behavior might change in the future and this document will have to be updated.

The broker_ids in the assignments is also ignored for the same reason. Currently OpenTalaria does not support more than one broker in a cluster.