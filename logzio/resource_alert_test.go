package logzio

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"io/ioutil"
	"log"
	"testing"
)

func TestAccLogzioAlert_CreateAlert(t *testing.T) {
	alertName := "test_create_alert"
	resourceName := "logzio_alert." + alertName

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: resourceCreateAlert("create_alert.tf", alertName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "title", "hello"),
					resource.TestCheckResourceAttr(resourceName, "severity_threshold_tiers.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "severity_threshold_tiers.0.severity", "HIGH"),
					resource.TestCheckResourceAttr(resourceName, "severity_threshold_tiers.0.threshold", "10"),
				),
			},
		},
	})
}

func TestAccLogzioAlert_CreateAlertMultiple(t *testing.T) {
	alertName := "test_create_alert"

	var checks []resource.TestCheckFunc
	for i := 0; i<10; i++ {
		resourceName := fmt.Sprintf("logzio_alert.%s.%d", alertName, i)

		checks = append(
			checks,
			resource.TestCheckResourceAttr(resourceName, "title", fmt.Sprintf("hello %d", i)),
			resource.TestCheckResourceAttr(resourceName, "severity_threshold_tiers.#", "1"),
			resource.TestCheckResourceAttr(resourceName, "severity_threshold_tiers.0.severity", "HIGH"),
			resource.TestCheckResourceAttr(resourceName, "severity_threshold_tiers.0.threshold", "10"),
		)
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: resourceCreateAlert("create_alert_multiple.tf", alertName),
				Check: resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

func TestAccLogzioAlert_UpdateAlert(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: resourceCreateAlert("create_alert.tf", "test_update_alert"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"logzio_alert.test_update_alert", "title", "hello"),
				),
			},
			resource.TestStep{
				Config: resourceUpdateAlert("test_update_alert"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"logzio_alert.test_update_alert", "title", "updated_alert"),
				),
			},
		},
	})
}

func TestAccLogzioAlert_UpdateFilterWithOrderingDifference(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: resourceCreateAlert("create_alert.tf", "test_update_filter_alert"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"logzio_alert.test_update_filter_alert", "filter", "{\"bool\":{\"must\":[],\"should\":[],\"filter\":[],\"must_not\":[]}}"),
				),
			},
			resource.TestStep{
				Config: resourceUpdateFilterAlert("test_update_filter_alert"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"logzio_alert.test_update_filter_alert", "filter", "{\"bool\":{\"must\":[],\"should\":[],\"filter\":[],\"must_not\":[]}}"),
				),
			},
		},
	})
}

func resourceCreateAlert(fixtureFile, name string) string {
	content, err := ioutil.ReadFile(fmt.Sprintf("testdata/fixtures/%s", fixtureFile))
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf(fmt.Sprintf("%s", content), name)
}

func resourceUpdateAlert(name string) string {
	content, err := ioutil.ReadFile("testdata/fixtures/update_alert.tf")
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf(fmt.Sprintf("%s", content), name)
}

func resourceUpdateFilterAlert(name string) string {
	content, err := ioutil.ReadFile("testdata/fixtures/update_filter_alert.tf")
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf(fmt.Sprintf("%s", content), name)
}