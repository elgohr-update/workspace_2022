// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package snippets

// [START healthcare_update_resource]
import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"

	healthcare "google.golang.org/api/healthcare/v1"
)

// updateFHIRResource updates an FHIR resource to be active or not.
func updateFHIRResource(w io.Writer, projectID, location, datasetID, fhirStoreID, resourceType, fhirResourceID string, active bool) error {
	ctx := context.Background()

	healthcareService, err := healthcare.NewService(ctx)
	if err != nil {
		return fmt.Errorf("healthcare.NewService: %v", err)
	}

	fhirService := healthcareService.Projects.Locations.Datasets.FhirStores.Fhir

	// The following payload works with a Patient resource and is not
	// intended to work with other types of FHIR resources. If necessary,
	// supply a new payload with data that corresponds to the FHIR resource
	// you are updating.
	payload := map[string]interface{}{
		"resourceType": resourceType,
		"id":           fhirResourceID,
		"active":       active,
	}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("json.Encode: %v", err)
	}

	name := fmt.Sprintf("projects/%s/locations/%s/datasets/%s/fhirStores/%s/fhir/%s/%s", projectID, location, datasetID, fhirStoreID, resourceType, fhirResourceID)

	call := fhirService.Update(name, bytes.NewReader(jsonPayload))
	call.Header().Set("Content-Type", "application/fhir+json;charset=utf-8")
	resp, err := call.Do()
	if err != nil {
		return fmt.Errorf("Update: %v", err)
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("could not read response: %v", err)
	}

	if resp.StatusCode > 299 {
		return fmt.Errorf("Update: status %d %s: %s", resp.StatusCode, resp.Status, respBytes)
	}
	fmt.Fprintf(w, "%s", respBytes)

	return nil
}

// [END healthcare_update_resource]
