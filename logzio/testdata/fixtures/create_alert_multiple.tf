resource "logzio_alert" "%s" {
  count = 10
  title = "hello ${count.index}"
  query_string = "loglevel:ERROR"
  operation = "GREATER_THAN"
  filter = "{\"bool\":{\"must\":[],\"should\":[],\"filter\":[],\"must_not\":[]}}"
  notification_emails = ["testx@test.com"]
  search_timeframe_minutes = 5
  value_aggregation_type = "NONE"
  alert_notification_endpoints = []
  suppress_notifications_minutes = 5
  severity_threshold_tiers {
    severity = "HIGH"
    threshold = 10
  }
}