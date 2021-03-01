package secrets

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
