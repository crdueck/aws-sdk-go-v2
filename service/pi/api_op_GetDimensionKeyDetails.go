// Code generated by smithy-go-codegen DO NOT EDIT.

package pi

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/pi/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Get the attributes of the specified dimension group for a DB instance or data
// source. For example, if you specify a SQL ID, GetDimensionKeyDetails retrieves
// the full text of the dimension db.sql.statement associated with this ID. This
// operation is useful because GetResourceMetrics and DescribeDimensionKeys don't
// support retrieval of large SQL statement text.
func (c *Client) GetDimensionKeyDetails(ctx context.Context, params *GetDimensionKeyDetailsInput, optFns ...func(*Options)) (*GetDimensionKeyDetailsOutput, error) {
	if params == nil {
		params = &GetDimensionKeyDetailsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDimensionKeyDetails", params, optFns, c.addOperationGetDimensionKeyDetailsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDimensionKeyDetailsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDimensionKeyDetailsInput struct {

	// The name of the dimension group. Performance Insights searches the specified
	// group for the dimension group ID. The following group name values are valid:
	//   - db.query (Amazon DocumentDB only)
	//   - db.sql (Amazon RDS and Aurora only)
	//
	// This member is required.
	Group *string

	// The ID of the dimension group from which to retrieve dimension details. For
	// dimension group db.sql , the group ID is db.sql.id . The following group ID
	// values are valid:
	//   - db.sql.id for dimension group db.sql (Aurora and RDS only)
	//   - db.query.id for dimension group db.query (DocumentDB only)
	//
	// This member is required.
	GroupIdentifier *string

	// The ID for a data source from which to gather dimension data. This ID must be
	// immutable and unique within an Amazon Web Services Region. When a DB instance is
	// the data source, specify its DbiResourceId value. For example, specify
	// db-ABCDEFGHIJKLMNOPQRSTU1VW2X .
	//
	// This member is required.
	Identifier *string

	// The Amazon Web Services service for which Performance Insights returns data.
	// The only valid value is RDS .
	//
	// This member is required.
	ServiceType types.ServiceType

	// A list of dimensions to retrieve the detail data for within the given dimension
	// group. If you don't specify this parameter, Performance Insights returns all
	// dimension data within the specified dimension group. Specify dimension names for
	// the following dimension groups:
	//   - db.sql - Specify either the full dimension name db.sql.statement or the
	//   short dimension name statement (Aurora and RDS only).
	//   - db.query - Specify either the full dimension name db.query.statement or the
	//   short dimension name statement (DocumentDB only).
	RequestedDimensions []string

	noSmithyDocumentSerde
}

type GetDimensionKeyDetailsOutput struct {

	// The details for the requested dimensions.
	Dimensions []types.DimensionKeyDetail

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetDimensionKeyDetailsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetDimensionKeyDetails{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetDimensionKeyDetails{}, middleware.After)
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
	if err = addOpGetDimensionKeyDetailsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetDimensionKeyDetails(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetDimensionKeyDetails(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "pi",
		OperationName: "GetDimensionKeyDetails",
	}
}
