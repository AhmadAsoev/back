package structs

type TableRequest struct {
	ID                   int    `json:"id"`
	ActNumber            string `json:"actNumber"`
	SenderName           string `json:"senderName"`
	SenderPositon        string `json:"senderPosition"`
	SenderOrganization   string `json:"senderOrganization"`
	Date                 string `json:"date"`
	ReceiverName         string `json:"receiverName"`
	ReceiverPosition     string `json:"receiverPosition"`
	ReceiverOrganization string `json:"receiverOrganization"`
}
