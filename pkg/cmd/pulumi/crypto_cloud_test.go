// Copyright 2016-2022, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	"github.com/stretchr/testify/assert"
	"gocloud.dev/secrets"
	"gocloud.dev/secrets/driver"
)

func deleteFiles(t *testing.T, files map[string]string) {
	for file := range files {
		err := os.Remove(file)
		assert.Nil(t, err, "Should be able to remove the file directory")
	}
}

func createTempFiles(t *testing.T, files map[string]string, f func()) {
	for file, content := range files {
		fileError := os.WriteFile(file, []byte(content), 0600)
		assert.Nil(t, fileError, "should be able to write the file contents")
	}

	defer deleteFiles(t, files)
	f()
}

//nolint:paralleltest
func TestSecretsProviderOverride(t *testing.T) {
	// Don't call t.Parallel because we temporarily modify
	// PULUMI_CLOUD_SECRET_OVERRIDE env var and it may interfere with other
	// tests.

	stackConfigFileName := "Pulumi.TestSecretsProviderOverride.yaml"
	files := make(map[string]string)
	files["Pulumi.yaml"] = "{\"name\":\"test\", \"runtime\":\"dotnet\"}"
	files[stackConfigFileName] = ""

	var stackName = tokens.Name("TestSecretsProviderOverride")

	opener := &mockSecretsKeeperOpener{}
	secrets.DefaultURLMux().RegisterKeeper("test", opener)

	//nolint:paralleltest
	t.Run("without override", func(t *testing.T) {
		createTempFiles(t, files, func() {
			opener.wantURL = "test://foo"
			_, createSecretsManagerError := newCloudSecretsManager(stackName, stackConfigFileName, "test://foo")
			assert.Nil(t, createSecretsManagerError, "Creating the cloud secret manager should succeed")

			_, createSecretsManagerError = newCloudSecretsManager(stackName, stackConfigFileName, "test://bar")
			msg := "newCloudSecretsManager with unexpected secretsProvider URL succeeded, expected an error"
			assert.NotNil(t, createSecretsManagerError, msg)
		})
	})

	//nolint:paralleltest
	t.Run("with override", func(t *testing.T) {
		createTempFiles(t, files, func() {
			opener.wantURL = "test://bar"
			t.Setenv("PULUMI_CLOUD_SECRET_OVERRIDE", "test://bar")

			// Last argument here shouldn't matter anymore, since it gets overridden
			// by the env var. Both calls should succeed.
			msg := "creating the secrets manager should succeed regardless of secrets provider"
			_, createSecretsManagerError := newCloudSecretsManager(stackName, stackConfigFileName, "test://foo")
			assert.Nil(t, createSecretsManagerError, msg)
			_, createSecretsManagerError = newCloudSecretsManager(stackName, stackConfigFileName, "test://bar")
			assert.Nil(t, createSecretsManagerError, msg)
		})
	})
}

type mockSecretsKeeperOpener struct {
	wantURL string
}

func (m *mockSecretsKeeperOpener) OpenKeeperURL(ctx context.Context, u *url.URL) (*secrets.Keeper, error) {
	if m.wantURL != u.String() {
		return nil, fmt.Errorf("got keeper URL: %q, want: %q", u, m.wantURL)
	}
	return secrets.NewKeeper(dummySecretsKeeper{}), nil
}

type dummySecretsKeeper struct {
	driver.Keeper
}

func (k dummySecretsKeeper) Decrypt(ctx context.Context, ciphertext []byte) ([]byte, error) {
	return ciphertext, nil
}

func (k dummySecretsKeeper) Encrypt(ctx context.Context, plaintext []byte) ([]byte, error) {
	return plaintext, nil
}
