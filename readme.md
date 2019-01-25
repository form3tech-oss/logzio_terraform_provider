# Logz.io Terraform provider

### Supports CRUD of Logz.io alerts and notification endpoints

This provider is based on the Logz.io client library - https://github.com/jonboydell/logzio_client


##### Obtaining the provider

You can [build from souce](#build-from-source) and copy the output into your Terraform templates folder, Terraform will find it here.

You can [get a release from here](https://github.com/jonboydell/logzio_terraform_provider/releases) and put it into your Terraform templates folder.

You'll need to do a `terraform init` for it to pick up the provider.


##### Using the provider

This simple example will create a Logz.io Slack notification endpoint (you'll need to provide the right URL) and an alert that
is triggered should Logz.io record 10 loglevel:ERROR messages in 5 minutes.  To make this example work you will also need to provide
your Logz.io API token.

```hcl-terraform
provider "logzio" {
  api_token = "${var.api_token}"
}

resource "logzio_endpoint" "my_endpoint" {
  title = "my_endpoint"
  description = "hello"
  endpoint_type = "slack"
  slack {
    url = "${var.slack_url}"
  }
}

resource "logzio_alert" "my_alert" {
  title = "my_other_title"
  query_string = "loglevel:ERROR"
  operation = "GREATER_THAN"
  notification_emails = []
  search_timeframe_minutes = 5
  value_aggregation_type = "NONE"
  alert_notification_endpoints = ["${logzio_endpoint.my_endpoint.id}"]
  suppress_notifications_minutes = 5
  severity_threshold_tiers = [
    {
      "severity" = "HIGH",
      "threshold" = 10
    }
  ]
}
```

##### Doens't work?

Do an [https://github.com/jonboydell/logzio_terraform_provider/issues](issue).

Fix it yourself and do a [https://github.com/jonboydell/logzio_terraform_provider/pulls](PR), please create any fix branches from `develop`.  They'll be merged back into `develop` and go `master` from there.  Releases are from `master`.