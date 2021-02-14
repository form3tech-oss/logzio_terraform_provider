module github.com/logzio/logzio_terraform_provider

go 1.12

require (
	github.com/hashicorp/terraform v0.12.6
	github.com/logzio/logzio_terraform_client v1.3.1
	github.com/stretchr/testify v1.3.0
)

replace github.com/logzio/logzio_terraform_client => /home/peternguyen/form3/logzio_client
