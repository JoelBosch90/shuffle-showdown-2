import { Stack, CfnOutput, RemovalPolicy } from 'aws-cdk-lib'
import { BlockPublicAccess, Bucket } from 'aws-cdk-lib/aws-s3'
import { BucketDeployment, Source } from 'aws-cdk-lib/aws-s3-deployment'

export const createWebsiteBucket = (stack: Stack, bucketName: string): Bucket => {
  const bucket = new Bucket(stack, 'Website', {
    bucketName,
    websiteIndexDocument: 'index.html',
    websiteErrorDocument: 'error.html',
    publicReadAccess: true,
    blockPublicAccess: BlockPublicAccess.BLOCK_ACLS,
    removalPolicy: RemovalPolicy.DESTROY,
    autoDeleteObjects: true,
  })

  new BucketDeployment(stack, 'DeployWebsite', {
    sources: [Source.asset('../source/assets')],
    destinationBucket: bucket,
  })

  new CfnOutput(stack, 'WebsiteUrl', {
    value: bucket.bucketWebsiteUrl,
  })

  return bucket
}