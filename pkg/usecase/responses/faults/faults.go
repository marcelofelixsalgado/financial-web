package faults

type IFaultMessage interface {
	// AddMessageByErrorCode(faults.ErrorCode) ResponseMessage
	// AddMessageByIssue(faults.Issue, Location, string, string, ...string) ResponseMessage
	GetMessage() FaultMessage
	// Write(http.ResponseWriter)
}

type FaultMessage struct {
	HttpStatusCode int                  `json:"-"`
	ErrorCode      string               `json:"error_code"`
	Message        string               `json:"message"`
	Details        []FaultMessageDetail `json:"details,omitempty"`
}

type FaultMessageDetail struct {
	Issue       string   `json:"issue"`
	Description string   `json:"description"`
	Location    Location `json:"location,omitempty"`
	Field       string   `json:"field,omitempty"`
	Value       string   `json:"value,omitempty"`
}

type Location string

const (
	Body           Location = "body"
	Header         Location = "header"
	QueryParameter Location = "query_parameter"
	PathParameter  Location = "path_parameter"
)

func NewResponseMessage() *FaultMessage {
	return &FaultMessage{}
}

func (responseMessage FaultMessage) GetMessage() FaultMessage {
	return FaultMessage{
		HttpStatusCode: responseMessage.HttpStatusCode,
		ErrorCode:      responseMessage.ErrorCode,
		Message:        responseMessage.Message,
		Details:        responseMessage.Details,
	}
}
