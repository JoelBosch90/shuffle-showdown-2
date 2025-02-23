import { Stack, StackProps, Aws } from 'aws-cdk-lib'
import { Construct } from 'constructs'
import { createWebsiteBucket } from './createWebsiteBucket'
import { createGitHubActionsDeploymentRole } from './createGitHubActionsDeploymentRole'

export class InfrastructureStack extends Stack {
  constructor(scope: Construct, id: string, thumbprint: string, props: StackProps = {}) {
    super(scope, id, props)

    const websiteBucketName = `clientstack-website-${Aws.ACCOUNT_ID}-${Aws.REGION}`
    createWebsiteBucket(this, websiteBucketName)
    createGitHubActionsDeploymentRole(this, websiteBucketName)
  }
}
