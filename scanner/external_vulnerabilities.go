package scanner

type ExternalVulnerabilities struct {
  Critcal int `json:"critcal"`
  High int `json:"high"`
  Medium int `json:"medium"`
  Low int `json:"low"`
}
