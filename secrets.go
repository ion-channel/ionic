package ionic

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/ion-channel/ionic/secrets"
)

//GetSecrets takes a text input and returns any matching secrets
func (ic *IonClient) GetSecrets(text string, token string) ([]secrets.Secret, error) {
	b, err := ic.Post(secrets.SecretsGetSecrets, token, nil, *bytes.NewBuffer([]byte(text)), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get secrets: %v", err.Error())
	}

	var s []secrets.Secret
	err = json.Unmarshal(b, &s)
	if err != nil {
		return nil, fmt.Errorf("failed to read response from get secrets: %v", err.Error())
	}

	return s, nil
}
