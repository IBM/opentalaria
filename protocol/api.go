package protocol

type API interface {
	SetVersion(version int16)
	GetHeaderVersion() int16
	GetKey() int16
}
