/**
 * Copyright 2021 IBM Corp.
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

package authenticator

import (
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/secret-utils-lib/pkg/token"
	//"github.com/IBM/secret-utils-lib/pkg/utils"
	"go.uber.org/zap"
)

// ComputeIdentityAuthenticator ...
type ComputeIdentityAuthenticator struct {
	authenticator *core.ContainerAuthenticator
	logger        *zap.Logger
}

// NewComputeIdentityAuthenticator ...
func NewComputeIdentityAuthenticator(profileID string, logger *zap.Logger) *ComputeIdentityAuthenticator {
	logger.Info("Initializing compute identity authenticator")
	defer logger.Info("Initialized compute identity authenticator")
	ca := new(ComputeIdentityAuthenticator)
	ca.authenticator = new(core.ContainerAuthenticator)
	ca.authenticator.IAMProfileID = profileID
	ca.logger = logger
	return ca
}

// GetToken ...
func (ca *ComputeIdentityAuthenticator) GetToken(freshTokenRequired bool) (string, uint64, error) {
	ca.logger.Info("Fetching token using compute identity authenticator")
	var iamtoken string
	var err error
	var tokenlifetime uint64

	if !freshTokenRequired {
		ca.logger.Info("Retreiving existing token")
		iamtoken, err = ca.authenticator.GetToken()
		if err != nil {
			ca.logger.Error("Error fetching token", zap.Error(err))
			// If the error is w.r.t invalid profile ID (which can happen when api key is reset)
			// retry reads the credentials again and fetches the iam token with new credentials
			err = retry(ca, ca.logger, PODIDENTITY, err)
			if err != nil {
				return "", tokenlifetime, err
			}
		}
		tokenlifetime, err = token.CheckTokenLifeTime(iamtoken)
		if err == nil {
			return iamtoken, tokenlifetime, nil
		}
	}

	tokenResponse, err := ca.authenticator.RequestToken()
	if err != nil {
		ca.logger.Error("Error fetching token", zap.Error(err))
		// If the error is w.r.t invalid api key (which can happen when api key is reset)
		// retry reads the credentials again and fetches the iam token with new credentials
		err = retry(ca, ca.logger, PODIDENTITY, err)
		if err != nil {
			return "", tokenlifetime, err
		}
	} else {
		iamtoken = tokenResponse.AccessToken
	}

	tokenlifetime, err = token.CheckTokenLifeTime(iamtoken)
	if err != nil {
		ca.logger.Error("Error fetching tokenlifetime", zap.Error(err))
		return "", tokenlifetime, err
	}

	ca.logger.Info("Successfully fetched IAM token")
	return tokenResponse.AccessToken, tokenlifetime, nil
}

// GetSecret ...
func (ca *ComputeIdentityAuthenticator) GetSecret() string {
	return ca.authenticator.IAMProfileID
}

// SetSecret ...
func (ca *ComputeIdentityAuthenticator) SetSecret(secret string) {
	ca.authenticator.IAMProfileID = secret
}
