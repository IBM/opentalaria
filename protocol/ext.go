package protocol

import "fmt"

func (r *RequestHeader) String() string {
	var clientID string
	if r.ClientID != nil {
		clientID = *r.ClientID
	}
	return fmt.Sprintf("API Key: %d\nAPI Version: %d\nCorrelation ID: %d\nClient Name: %s",
		r.RequestApiKey,
		r.RequestApiVersion,
		r.CorrelationID,
		clientID)
}
