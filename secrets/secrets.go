package secrets

import (
	"encoding/json"
	"strings"
)

const (
	// SecretsGetSecrets is a string representation of the current endpoint for getting secrets
	SecretsGetSecrets = "v1/metadata/getSecrets"
)

// Secret is a struct containing matching data for a secret found in text
type Secret struct {
	// Rule - the defined rule that was matched
	Rule string `json:"rule"`
	// Match - the subtext that was matched
	Match string `json:"match"`
	// Confidence - a float value from 0.0 to 1.0 of our trust in the result
	Confidence float32 `json:"confidence"`
}

// UnmarshalJSON custom unmarshal to mask secrets
func (s *Secret) UnmarshalJSON(data []byte) error {
	var v map[string]interface{}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}

	// t, ok := v.(Secret)
	// if !ok {
	// 	return fmt.Errorf("Data was not in the form of a secret")
	// }
	s.Rule = v["rule"].(string)
	s.Match = v["match"].(string)
	s.Confidence = float32(v["confidence"].(float64))

	half := len(s.Match) / 2
	s.Match = s.Match[:half] + strings.Repeat("*", half)
	return nil
}
