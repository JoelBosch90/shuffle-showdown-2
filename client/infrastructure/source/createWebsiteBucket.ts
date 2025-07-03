import { Stack, CfnOutput, RemovalPolicy, Duration } from 'aws-cdk-lib';
import { Certificate } from 'aws-cdk-lib/aws-certificatemanager';
import { AllowedMethods, Distribution, OriginProtocolPolicy, ViewerProtocolPolicy } from 'aws-cdk-lib/aws-cloudfront';
import { HttpOrigin } from 'aws-cdk-lib/aws-cloudfront-origins';
import { BlockPublicAccess, Bucket } from 'aws-cdk-lib/aws-s3';
import { BucketDeployment, Source } from 'aws-cdk-lib/aws-s3-deployment';

export const createWebsiteBucket = (stack: Stack, name: string): Bucket => {
  const bucket = new Bucket(stack, 'Website', {
    bucketName: name.replace(/\:/g, '-'),
    removalPolicy: RemovalPolicy.DESTROY,
    autoDeleteObjects: true,
    publicReadAccess: true,
    blockPublicAccess: new BlockPublicAccess({
      blockPublicAcls: false,
      blockPublicPolicy: false,
      ignorePublicAcls: false,
      restrictPublicBuckets: false,
    }),
    versioned: true,
    websiteIndexDocument: 'index.html',
    websiteErrorDocument: 'index.html',
  });

  new BucketDeployment(stack, 'DeployWebsite', {
    sources: [Source.asset('../app/build')],
    destinationBucket: bucket,
  });

  let websiteUrl = bucket.bucketWebsiteUrl;

  const isLocalEnvironment = process.env.AWS_ENDPOINT_URL?.includes('localhost') || process.env.AWS_ENDPOINT_URL?.includes('localstack')

  // If not in a local environment, set up CloudFront distribution
  if (!isLocalEnvironment) {
    const certificateArn = process.env.SHUFFLE_SHOWDOWN_DOMAIN_CERTIFICATE_ARN;
    if (!certificateArn) {
      throw new Error('Certificate ARN is not defined in environment variables.');
    }
    const certificate = Certificate.fromCertificateArn(stack, 'WebsiteCertificate', certificateArn);

    const distribution = new Distribution(stack, 'WebsiteDistribution', {
      domainNames: [name],
      defaultBehavior: {
        origin: new HttpOrigin(bucket.bucketWebsiteDomainName, {
          protocolPolicy: OriginProtocolPolicy.HTTP_ONLY,
          httpPort: 80,
          httpsPort: 443,
        }),
        allowedMethods: AllowedMethods.ALLOW_GET_HEAD,
        compress: true,
        viewerProtocolPolicy: ViewerProtocolPolicy.REDIRECT_TO_HTTPS,
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

    websiteUrl = distribution.domainName;
  }

  new CfnOutput(stack, 'WebsiteUrl', { value: websiteUrl });

  return bucket;
};