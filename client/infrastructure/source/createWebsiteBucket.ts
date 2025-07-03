import { Stack, CfnOutput, RemovalPolicy } from 'aws-cdk-lib';
import { BlockPublicAccess, Bucket } from 'aws-cdk-lib/aws-s3';
import { BucketDeployment, Source } from 'aws-cdk-lib/aws-s3-deployment';
import { createCloudFrontDistribution } from './createCloudFrontDistribution';
import { isLocalEnvironment } from './isLocalEnvironment';

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

  // If not in a local environment, set up CloudFront distribution
  if (!isLocalEnvironment()) {
    const distribution = createCloudFrontDistribution(stack, bucket, name);

    websiteUrl = distribution.domainName;
  }

  new CfnOutput(stack, 'WebsiteUrl', { value: websiteUrl });

  return bucket;
};