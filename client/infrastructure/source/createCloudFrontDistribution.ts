import { Stack, Duration } from 'aws-cdk-lib';
import { AllowedMethods, CachePolicy, Distribution, OriginProtocolPolicy, OriginRequestPolicy, ViewerProtocolPolicy } from 'aws-cdk-lib/aws-cloudfront';
import { HttpOrigin } from 'aws-cdk-lib/aws-cloudfront-origins';
import { Bucket } from 'aws-cdk-lib/aws-s3';
import { createCertificate } from './createCertificate';
import { isLocalEnvironment } from './isLocalEnvironment';
import { createApiGatewayOrigin } from './createApiGatewayOrigin';

export const createCloudFrontDistribution = (stack: Stack, bucket: Bucket, domainName: string): Distribution => {
  const certificate = createCertificate(stack);
  const viewerProtocolPolicy = ViewerProtocolPolicy.REDIRECT_TO_HTTPS;

  const distribution = new Distribution(stack, 'WebsiteDistribution', {
    domainNames: [domainName],
    defaultBehavior: {
      origin: new HttpOrigin(bucket.bucketWebsiteDomainName, {
        protocolPolicy: OriginProtocolPolicy.HTTP_ONLY,
        httpPort: 80,
        httpsPort: 443,
      }),
      allowedMethods: AllowedMethods.ALLOW_GET_HEAD,
      compress: true,
      viewerProtocolPolicy,
    },
    additionalBehaviors: isLocalEnvironment() ? {} : {
      '/api/*': {
        origin: createApiGatewayOrigin(),
        allowedMethods: AllowedMethods.ALLOW_ALL,
        cachePolicy: CachePolicy.CACHING_DISABLED,
        originRequestPolicy: OriginRequestPolicy.ALL_VIEWER,
        viewerProtocolPolicy,
      },
    },
    errorResponses: [
      {
        httpStatus: 403,
        responseHttpStatus: 200,
        responsePagePath: '/index.html',
        ttl: Duration.seconds(0),
      },
      {
        httpStatus: 404,
        responseHttpStatus: 200,
        responsePagePath: '/index.html',
        ttl: Duration.seconds(0),
      },
    ],
    certificate,
  });

  return distribution;
};