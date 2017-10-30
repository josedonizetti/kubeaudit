package cmd

import (
	"testing"

	"github.com/Shopify/kubeaudit/fakeaudit"
	"github.com/stretchr/testify/assert"
)

func init() {
	fakeaudit.CreateFakeNamespace("fakeDeploymentRANR")
	fakeaudit.CreateFakeDeploymentRunAsNonRoot("fakeDeploymentRANR")
}

func runTest(t *testing.T, file string, function func(Items) []Result, errCode int) {
	assert := assert.New(t)
	resources, err := getKubeResourcesManifest("../fixtures/security_context_nil.yml")
	assert.Nil(err)
	count := len(resources)
	wg.Add(count)
	var results []Result
	for _, resource := range resources {
		current_results := function(resource)
		for _, current_result := range current_results {
			results = append(results, current_result)
		}
	}
	wg.Wait()
	for _, result := range results {
		assert.Equal(result.Occurrences[0].id, errCode)
	}
}

func TestSecurityContextNIL(t *testing.T) {
	runTest(t, "../fixtures/security_context_nil.yml", auditRunAsNonRoot, ErrorSecurityContextNIL)
}

func TestRunAsNonRootNil(t *testing.T) {
	runTest(t, "../fixtures/run_as_non_root_nil.yml", auditRunAsNonRoot, ErrorRunAsNonRootNIL)
}

func TestRunAsNonRootFalse(t *testing.T) {
	runTest(t, "../fixtures/run_as_non_root_false.yml", auditRunAsNonRoot, ErrorRunAsNonRootFalse)
}

func TestRunAsNonRootTrue(t *testing.T) {
	//runTest("../fixtures/run_as_non_root_false.yml", auditRunAsNonRoot, ErrorRunAsNonRootFalse)
}
