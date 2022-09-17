// Copyright 2016-2022, Pulumi Corporation.  All rights reserved.
//go:build nodejs || all

package ints

import (
	"path/filepath"
	"runtime"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/stretchr/testify/assert"
)

// TestEmptyNodeJS simply tests that we can run an empty NodeJS project.
func TestEmptyNodeJS(t *testing.T) {
	integration.ProgramTest(t, &integration.ProgramTestOptions{
		Dir:          filepath.Join("empty", "nodejs"),
		Dependencies: []string{"@pulumi/pulumi"},
		Quick:        true,
	})
}

// Tests that stack references work in Node.
func TestStackReferenceNodeJS(t *testing.T) {
	if runtime.GOOS == WindowsOS {
		t.Skip("Temporarily skipping test on Windows - pulumi/pulumi#3811")
	}

	opts := &integration.ProgramTestOptions{
		RequireService: true,

		Dir:          filepath.Join("stack_reference", "nodejs"),
		Dependencies: []string{"@pulumi/pulumi"},
		Quick:        true,
		EditDirs: []integration.EditDir{
			{
				Dir:      "step1",
				Additive: true,
			},
			{
				Dir:      "step2",
				Additive: true,
			},
		},
	}
	integration.ProgramTest(t, opts)
}

// Test remote component construction in Node.
func TestConstructNode(t *testing.T) {
	if runtime.GOOS == WindowsOS {
		t.Skip("Temporarily skipping test on Windows")
	}
	t.Parallel()

	testDir := "construct_component"
	runComponentSetup(t, testDir)

	tests := []struct {
		componentDir          string
		expectedResourceCount int
	}{
		{
			componentDir:          "testcomponent",
			expectedResourceCount: 9,
		},
		{
			componentDir:          "testcomponent-python",
			expectedResourceCount: 9,
		},
		{
			componentDir:          "testcomponent-go",
			expectedResourceCount: 8, // One less because no dynamic provider.
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.componentDir, func(t *testing.T) {
			pathEnv := pathEnv(t,
				buildTestProvider(t, filepath.Join("..", "testprovider")),
				filepath.Join(testDir, test.componentDir))
			integration.ProgramTest(t,
				optsForConstructNode(t, test.expectedResourceCount, pathEnv))
		})
	}
}

func optsForConstructNode(t *testing.T, expectedResourceCount int, env ...string) *integration.ProgramTestOptions {
	return &integration.ProgramTestOptions{
		Env:          env,
		Dir:          filepath.Join("construct_component", "nodejs"),
		Dependencies: []string{"@pulumi/pulumi"},
		Secrets: map[string]string{
			"secret": "this super secret is encrypted",
		},
		Quick: true,
		// verify that additional flags don't cause the component provider hang
		UpdateCommandlineFlags: []string{"--logflow", "--logtostderr"},
		ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
			assert.NotNil(t, stackInfo.Deployment)
			if assert.Equal(t, expectedResourceCount, len(stackInfo.Deployment.Resources)) {
				stackRes := stackInfo.Deployment.Resources[0]
				assert.NotNil(t, stackRes)
				assert.Equal(t, resource.RootStackType, stackRes.Type)
				assert.Equal(t, "", string(stackRes.Parent))

				// Check that dependencies flow correctly between the originating program and the remote component
				// plugin.
				urns := make(map[string]resource.URN)
				for _, res := range stackInfo.Deployment.Resources[1:] {
					assert.NotNil(t, res)

					urns[string(res.URN.Name())] = res.URN
					switch res.URN.Name() {
					case "child-a":
						for _, deps := range res.PropertyDependencies {
							assert.Empty(t, deps)
						}
					case "child-b":
						expected := []resource.URN{urns["a"]}
						assert.ElementsMatch(t, expected, res.Dependencies)
						assert.ElementsMatch(t, expected, res.PropertyDependencies["echo"])
					case "child-c":
						expected := []resource.URN{urns["a"], urns["child-a"]}
						assert.ElementsMatch(t, expected, res.Dependencies)
						assert.ElementsMatch(t, expected, res.PropertyDependencies["echo"])
					case "a", "b", "c":
						secretPropValue, ok := res.Outputs["secret"].(map[string]interface{})
						assert.Truef(t, ok, "secret output was not serialized as a secret")
						assert.Equal(t, resource.SecretSig, secretPropValue[resource.SigKey].(string))
					}
				}
			}
		},
	}
}
