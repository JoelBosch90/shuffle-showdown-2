import { OpenIdConnectProvider, PolicyDocument, PolicyStatement, Effect, Role, WebIdentityPrincipal } from "aws-cdk-lib/aws-iam";
import { Stack } from "aws-cdk-lib/core/lib/stack";

export const createGitHubActionsDeploymentRole = async (stack: Stack, thumbprint: string = ''): Promise<Role> => {
  const oidcProvider = new OpenIdConnectProvider(stack, 'GitHubOIDCProvider', {
    url: 'https://token.actions.githubusercontent.com',
    thumbprints: ['74f3a68f16524f15424927704c9506f55a9316bd', thumbprint],
    clientIds: ['sts.amazonaws.com'],
  });

  const webIdentityPrincipal = new WebIdentityPrincipal(
    oidcProvider.openIdConnectProviderArn,
    {
      StringLike: {
        'token.actions.githubusercontent.com:sub': 'repo:JoelBosch90/shuffle-showdown-2:*',
        'token.actions.githubusercontent.com:aud': 'sts.amazonaws.com',
      },
    },
  );

  const assumeRolePolicyStatement = new PolicyStatement({
    effect: Effect.ALLOW,
    actions: ['sts:AssumeRole'],
    resources: [`arn:aws:iam::${stack.account}:role/cdk-*`],
  });
  const ssmPolicyStatement = new PolicyStatement({
    effect: Effect.ALLOW,
    actions: ['ssm:GetParameter'],
    resources: [`arn:aws:ssm:${stack.region}:${stack.account}:parameter/cdk-*`],
  });

  const policyDocument = new PolicyDocument({
    statements: [assumeRolePolicyStatement, ssmPolicyStatement],
  });

  const deploymentRole = new Role(stack, 'GitHubActionsDeploymentRole', {
    assumedBy: webIdentityPrincipal,
    roleName: 'github-actions-deployment-role',
    inlinePolicies: { 'github-actions-deployment-policy': policyDocument },
  });

  return deploymentRole;
};