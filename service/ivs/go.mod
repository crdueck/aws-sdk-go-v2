module github.com/aws/aws-sdk-go-v2/service/ivs

go 1.15

require (
	github.com/aws/aws-sdk-go-v2 v1.16.12
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.1.19
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.4.13
	github.com/aws/smithy-go v1.13.1
)

replace github.com/aws/aws-sdk-go-v2 => ../../

replace github.com/aws/aws-sdk-go-v2/internal/configsources => ../../internal/configsources/

replace github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 => ../../internal/endpoints/v2/
