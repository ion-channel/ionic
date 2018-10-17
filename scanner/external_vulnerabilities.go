package scanner

//ExternalVulnerability is a representation of a vulnerability scan provided by the client
type ExternalVulnerability struct {
	Critcal int `json:"critical"`
	High    int `json:"high"`
	Medium  int `json:"medium"`
	Low     int `json:"low"`
}
