/**
 * Copyright 2022 IBM Corp.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package utils

import "fmt"

// Error is structure that is defined to locally to represent any error and it implements golang error
type Error struct {
	Description  string
	BackendError string
	Action       string
}

// Error method implements the Error method golang error.
func (err Error) Error() string {
	var errmsg string
	if err.Description != "" {
		errmsg = fmt.Sprintf("Description: %s ", err.Description)
	}
	if err.BackendError != "" {
		errmsg += fmt.Sprintf("BackendError: %s ", err.BackendError)
	}
	if err.Action != "" {
		errmsg += fmt.Sprintf("Action: %s ", err.Action)
	}
	return errmsg
}

const (
	// ErrCredentialsUndefined ...
	ErrCredentialsUndefined = "ibmcloud credentials undefined"

	// ErrInvalidCredentialsFormat ...
	ErrInvalidCredentialsFormat = "ibmcloud credentials are provided in invalid format, unable to parse the credentials"

	// ErrAuthTypeUndefined ...
	ErrAuthTypeUndefined = "IBMCLOUD_AUTHTYPE undefined"

	// ErrUnknownCredentialType ...
	ErrUnknownCredentialType = "Unknown IBMCLOUD_AUTHTYPE provided. IBMCLOUD_AUTHTYPE: %s"

	// ErrAPIKeyNotProvided ...
	ErrAPIKeyNotProvided = "API key is not provided"

	// ErrProfileIDNotProvided ...
	ErrProfileIDNotProvided = "Profile ID is not provided"

	// APIKeyNotFound ...
	APIKeyNotFound = "api key could not be found"

	// UserNotFound ...
	UserNotFound = "user not found or active"

	// ProfileNotFound ...
	ProfileNotFound = "selected trusted profile not eligible for cr token"

	// ErrSecretConfigPathUndefined ...
	ErrSecretConfigPathUndefined = "SECRET_CONFIG_PATH is not defined"

	// ErrEmptyTokenResponse ...
	ErrEmptyTokenResponse = "Empty token response received"
)
