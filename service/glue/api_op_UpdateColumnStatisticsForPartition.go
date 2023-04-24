// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates or updates partition statistics of columns. The Identity and Access
// Management (IAM) permission required for this operation is UpdatePartition .
func (c *Client) UpdateColumnStatisticsForPartition(ctx context.Context, params *UpdateColumnStatisticsForPartitionInput, optFns ...func(*Options)) (*UpdateColumnStatisticsForPartitionOutput, error) {
	if params == nil {
		params = &UpdateColumnStatisticsForPartitionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateColumnStatisticsForPartition", params, optFns, c.addOperationUpdateColumnStatisticsForPartitionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateColumnStatisticsForPartitionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateColumnStatisticsForPartitionInput struct {

	// A list of the column statistics.
	//
	// This member is required.
	ColumnStatisticsList []types.ColumnStatistics

	// The name of the catalog database where the partitions reside.
	//
	// This member is required.
	DatabaseName *string

	// A list of partition values identifying the partition.
	//
	// This member is required.
	PartitionValues []string

	// The name of the partitions' table.
	//
	// This member is required.
	TableName *string

	// The ID of the Data Catalog where the partitions in question reside. If none is
	// supplied, the Amazon Web Services account ID is used by default.
	CatalogId *string

	noSmithyDocumentSerde
}

type UpdateColumnStatisticsForPartitionOutput struct {

	// Error occurred during updating column statistics data.
	Errors []types.ColumnStatisticsError

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateColumnStatisticsForPartitionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateColumnStatisticsForPartition{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateColumnStatisticsForPartition{}, middleware.After)
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
	if err = addOpUpdateColumnStatisticsForPartitionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateColumnStatisticsForPartition(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateColumnStatisticsForPartition(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "UpdateColumnStatisticsForPartition",
	}
}
