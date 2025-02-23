import { OpenIdConnectProvider, PolicyDocument, PolicyStatement, Effect, Role, WebIdentityPrincipal } from "aws-cdk-lib/aws-iam"
import { Stack } from "aws-cdk-lib/core/lib/stack"

export const createGitHubActionsDeploymentRole = (stack: Stack, websiteBucketName: string): Role => {
  const oidcProvider = new OpenIdConnectProvider(stack, 'GitHubOIDCProvider', {
    url: 'https://token.actions.githubusercontent.com',
    // See https://github.com/aws-actions/configure-aws-credentials/issues/357#issuecomment-1011642085 to get thumbprint.
    thumbprints: ['74f3a68f16524f15424927704c9506f55a9316bd'],
    clientIds: ['sts.amazonaws.com'],
  })

  const webIdentityPrincipal = new WebIdentityPrincipal(
    oidcProvider.openIdConnectProviderArn,
    {
      StringLike: {
        'token.actions.githubusercontent.com:sub': 'repo:JoelBosch90/shuffle-showdown-2:*',
        'token.actions.githubusercontent.com:aud': 'sts.amazonaws.com',
      },
    }
  )

  const s3PolicyStatement = new PolicyStatement({
    effect: Effect.ALLOW,
    actions: ['s3:ListObjects'],
    resources: [`arn:aws:s3:::${websiteBucketName}`, `arn:aws:s3:::${websiteBucketName}/*`],
  })

  const policyDocument = new PolicyDocument({
    statements: [s3PolicyStatement],
  })

  const deploymentRole = new Role(stack, 'GitHubActionsDeploymentRole', {
    assumedBy: webIdentityPrincipal,
    roleName: 'github-actions-deployment-role',
    inlinePolicies: { 'github-actions-deployment-policy': policyDocument },
  })

  return deploymentRole
}