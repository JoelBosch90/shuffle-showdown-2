import { Stack, CfnOutput, RemovalPolicy } from 'aws-cdk-lib';
import { Certificate } from 'aws-cdk-lib/aws-certificatemanager';
import { AllowedMethods, Distribution, OriginProtocolPolicy } from 'aws-cdk-lib/aws-cloudfront';
import { HttpOrigin } from 'aws-cdk-lib/aws-cloudfront-origins';
import { BlockPublicAccess, Bucket } from 'aws-cdk-lib/aws-s3';
import { BucketDeployment, Source } from 'aws-cdk-lib/aws-s3-deployment';

export const createWebsiteBucket = (stack: Stack, bucketName: string): Bucket => {
  const bucket = new Bucket(stack, 'Website', {
    bucketName,
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
  });

  new BucketDeployment(stack, 'DeployWebsite', {
    sources: [Source.asset('../app/build')],
    destinationBucket: bucket,
  });

  const certificateArn = process.env.SHUFFLE_SHOWDOWN_DOMAIN_CERTIFICATE_ARN;
  if (!certificateArn) {
    throw new Error('Certificate ARN is not defined in environment variables.');
  }
  const certificate = Certificate.fromCertificateArn(stack, 'WebsiteCertificate', certificateArn);

  const distribution = new Distribution(stack, 'WebsiteDistribution', {
    domainNames: [bucketName],
    defaultBehavior: {
      origin: new HttpOrigin(bucket.bucketWebsiteDomainName, {
        protocolPolicy: OriginProtocolPolicy.HTTP_ONLY,
        httpPort: 80,
        httpsPort: 443,
      }),
      allowedMethods: AllowedMethods.ALLOW_GET_HEAD,
      compress: true,
    },
    certificate,
  });

  new CfnOutput(stack, 'WebsiteUrl', { value: distribution.domainName });

  return bucket;
};