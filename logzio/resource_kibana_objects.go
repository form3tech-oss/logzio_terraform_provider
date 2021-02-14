package logzio

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/logzio/logzio_terraform_client/kibana_objects"
)

const (
	kibanaObjectsObjects string = "objects"
)

func resourceKibanaObjects() *schema.Resource {
	return &schema.Resource{
		Create: resourceKibanaObjectsCreate,
		Read:   resourceKibanaObjectsRead,
		Update: resourceKibanaObjectsUpdate,
		Delete: resourceKibanaObjectsDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			kibanaObjectsObjects: {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required: true,
			},
		},
	}
}

func kibanaObjectsClient(m interface{}) *kibana_objects.KibanaObjectsClient {
	var client *kibana_objects.KibanaObjectsClient
	client, _ = kibana_objects.New(m.(Config).apiToken, m.(Config).baseUrl)
	return client
}

func resourceKibanaObjectsCreate(d *schema.ResourceData, m interface{}) error {
	objects := d.Get(kibanaObjectsObjects).([]string)
	if len(objects) == 0 {
		return errors.New("'objects' cannot be empty")
	}

	payload := kibana_objects.ImportPayload{}

	for _, object := range objects {
		m, err := kibanaObjectsObjectToMap(object)
		if err != nil {
			return fmt.Errorf("could not convert object '%s' to map: %w", object, err)
		}

		payload.Hits = append(payload.Hits, m)
	}

	err := kibanaObjectsValidateImportPayload(payload)
	if err != nil {
		return err
	}

	results, err := kibanaObjectsClient(m).Import(payload)
	if err != nil {
		return err
	}

	if len(results.Updated) > 0 {
		return fmt.Errorf("unexpected values in import results for 'updated', %v", results.Updated)
	}

	if len(results.Ignored) > 0 {
		return fmt.Errorf("unexpected values in import results for 'ignored', %v", results.Ignored)
	}

	if len(results.Failed) > 0 {
		return fmt.Errorf("failed to create objects, %v", results.Failed)
	}

	return nil
}

func resourceKibanaObjectsRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceKibanaObjectsUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceKibanaObjectsDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}

// objectToMap will unmarshal an Kibana object json string to a map[string]interface{}
func kibanaObjectsObjectToMap(object string) (map[string]interface{}, error) {
	var m map[string]interface{}
	err := json.Unmarshal([]byte(object), &m)
	return m, err
}

// validateImportPayload validates the payload that will be sent to the import
// endpoint. Currently, it only does a check that each Kibana object has a
// name.
func kibanaObjectsValidateImportPayload(payload kibana_objects.ImportPayload) error {
	for _, object := range payload.Hits {
		v, ok := object["name"]
		if !ok {
			return fmt.Errorf("missing 'name' in object %v", object)
		}

		if v == "" {
			return fmt.Errorf("'name' cannot be empty in object %v", object)
		}
	}

	return nil
}
