// Code generated by smithy-go-codegen DO NOT EDIT.

package computeoptimizer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/computeoptimizer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns the projected metrics of Amazon ECS service recommendations.
func (c *Client) GetECSServiceRecommendationProjectedMetrics(ctx context.Context, params *GetECSServiceRecommendationProjectedMetricsInput, optFns ...func(*Options)) (*GetECSServiceRecommendationProjectedMetricsOutput, error) {
	if params == nil {
		params = &GetECSServiceRecommendationProjectedMetricsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetECSServiceRecommendationProjectedMetrics", params, optFns, c.addOperationGetECSServiceRecommendationProjectedMetricsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetECSServiceRecommendationProjectedMetricsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetECSServiceRecommendationProjectedMetricsInput struct {

	// The timestamp of the last projected metrics data point to return.
	//
	// This member is required.
	EndTime *time.Time

	// The granularity, in seconds, of the projected metrics data points.
	//
	// This member is required.
	Period int32

	// The ARN that identifies the Amazon ECS service. The following is the format of
	// the ARN: arn:aws:ecs:region:aws_account_id:service/cluster-name/service-name
	//
	// This member is required.
	ServiceArn *string

	// The timestamp of the first projected metrics data point to return.
	//
	// This member is required.
	StartTime *time.Time

	// The statistic of the projected metrics.
	//
	// This member is required.
	Stat types.MetricStatistic

	noSmithyDocumentSerde
}

type GetECSServiceRecommendationProjectedMetricsOutput struct {

	// An array of objects that describes the projected metrics.
	RecommendedOptionProjectedMetrics []types.ECSServiceRecommendedOptionProjectedMetric

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetECSServiceRecommendationProjectedMetricsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpGetECSServiceRecommendationProjectedMetrics{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpGetECSServiceRecommendationProjectedMetrics{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetECSServiceRecommendationProjectedMetricsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetECSServiceRecommendationProjectedMetrics(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opGetECSServiceRecommendationProjectedMetrics(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "compute-optimizer",
		OperationName: "GetECSServiceRecommendationProjectedMetrics",
	}
}
