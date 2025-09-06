import { Stack, Duration } from 'aws-cdk-lib';
import { AllowedMethods, CachePolicy, Distribution, OriginAccessIdentity, OriginProtocolPolicy, OriginRequestPolicy, ViewerProtocolPolicy } from 'aws-cdk-lib/aws-cloudfront';
import { HttpOrigin, S3Origin } from 'aws-cdk-lib/aws-cloudfront-origins';
import { Bucket } from 'aws-cdk-lib/aws-s3';
import { createCertificate } from './createCertificate';
import { isLocalEnvironment } from './isLocalEnvironment';
import { getApiPath } from './getApiPath';
import { getApiGatewayOrigin } from './getApiGatewayOrigin';

/**
 *  This function creates a CloudFront distribution for the website bucket.
 *  It sets up the distribution with the specified domain name and bucket.
 *  @param    {Stack} stack - The CDK stack in which to create the distribution.
 *  @param    {Bucket} bucket - The S3 bucket to be used as the origin for the distribution.
 *  @param    {string} domainName - The domain name for the CloudFront distribution.
 *  @returns  {Distribution} - The created CloudFront distribution.
 */
export const createCloudFrontDistribution = (stack: Stack, bucket: Bucket, domainName?: string): Distribution => {
  const certificate = createCertificate(stack);
  const viewerProtocolPolicy = ViewerProtocolPolicy.REDIRECT_TO_HTTPS;
  const apiPath = `/${getApiPath()}/*`;
  const indexDocument = '/index.html';

  // Create Origin Access Identity for S3 access
  const originAccessIdentity = new OriginAccessIdentity(stack, 'WebsiteOAI', {
    comment: 'OAI for website bucket'
  });

  // Grant CloudFront access to the bucket
  bucket.grantRead(originAccessIdentity);

  const distribution = new Distribution(stack, 'WebsiteDistribution', {
    domainNames: domainName ? [domainName] : undefined,
    defaultBehavior: {
      origin: new S3Origin(bucket, {
        originAccessIdentity
      }),
      allowedMethods: AllowedMethods.ALLOW_GET_HEAD,
      compress: true,
      viewerProtocolPolicy,
    },
    additionalBehaviors: isLocalEnvironment() ? {} : {
      [apiPath]: {
        origin: getApiGatewayOrigin(stack),
        allowedMethods: AllowedMethods.ALLOW_ALL,
        cachePolicy: CachePolicy.CACHING_DISABLED,
        originRequestPolicy: OriginRequestPolicy.ALL_VIEWER_EXCEPT_HOST_HEADER,
        viewerProtocolPolicy,
      },
    },
    errorResponses: [
      {
        httpStatus: 403,
        responseHttpStatus: 200,
        responsePagePath: indexDocument,
        ttl: Duration.seconds(0),
      },
      {
        httpStatus: 404,
        responseHttpStatus: 200,
        responsePagePath: indexDocument,
        ttl: Duration.seconds(0),
      },
    ],
    certificate: domainName ? certificate : undefined,
  });

  return distribution;
};