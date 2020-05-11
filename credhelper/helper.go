/*
 * Copyright 2000-2020 JetBrains s.r.o.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package credhelper

import (
	"errors"
	"fmt"
	"github.com/docker/docker-credential-helpers/credentials"
	"os"
)

const (
	spaceClientId     = "JB_SPACE_CLIENT_ID"
	spaceClientSecret = "JB_SPACE_CLIENT_SECRET"
)

// spaceCredHelper implements a credentials.Helper interface backed by a Space
// credential store.
type spaceCredHelper struct {
}

// NewSpaceCredentialHelper returns a Docker credential helper which
// specializes in Space authentication schemes.
func NewSpaceCredentialHelper() credentials.Helper {
	return &spaceCredHelper{}
}

// Delete lists all stored credentials and associated usernames.
func (*spaceCredHelper) List() (map[string]string, error) {
	return nil, errors.New("list is unimplemented")
}

// Add adds new third-party credentials to the keychain.
func (*spaceCredHelper) Add(*credentials.Credentials) error {
	return errors.New("add is unimplemented")
}

// Delete removes third-party credentials from the store.
func (*spaceCredHelper) Delete(string) error {
	return errors.New("delete is unimplemented")
}

// Get returns the username and secret to use for a given registry server URL.
func (ch *spaceCredHelper) Get(serverURL string) (string, string, error) {
	return ch.spaceCredentials()
}

func (ch *spaceCredHelper) spaceCredentials() (string, string, error) {
	var clientId = os.Getenv(spaceClientId)
	if len(clientId) == 0 {
		return "", "", helperErr(fmt.Sprintf("missing %s variable\n", spaceClientId), nil)
	}

	var clientSecret = os.Getenv(spaceClientSecret)
	if len(clientSecret) == 0 {
		return "", "", helperErr(fmt.Sprintf("missing %s variable\n", spaceClientSecret), nil)
	}

	return clientId, clientSecret, nil
}

func helperErr(message string, err error) error {
	if err == nil {
		return fmt.Errorf("docker-credential-space/helper: %s", message)
	}
	return fmt.Errorf("docker-credential-space/helper: %s: %v", message, err)
}
