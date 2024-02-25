package entitiy

type Wit struct {
	WARNING  string `json:"WARNING"`
	TEXT     string `json:"_text"`
	MSG_ID   string `json:"msg_id"`
	OUTCOMES []struct {
		TEXT       string            `json:"_text"`
		CONFIDENCE interface{}       `json:"confidence"`
		ENTITIES   map[string]string `json:"entities"`
		INTENT     string            `json:"intent"`
	} `json:"outcomes"`
}