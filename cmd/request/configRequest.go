package request

type ConfigRequest struct {
	ProfileType string        `json:"profileType"`
	Brand       string        `json:"brand"`
	VisitCount  int           `json:"visitCount"`
	BlockLobs   []interface{} `json:"blockLobs"`
}
