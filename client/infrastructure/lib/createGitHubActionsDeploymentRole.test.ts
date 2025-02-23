import { App, Stack } from 'aws-cdk-lib'
import { Effect } from "aws-cdk-lib/aws-iam"
import { Match, Template } from 'aws-cdk-lib/assertions'
import { createGitHubActionsDeploymentRole } from './createGitHubActionsDeploymentRole'

describe('createGitHubActionsDeploymentRole', () => {
  // Set dummy environment values for account and region so our ARNs are predictable.
  const dummyAccount = '123456789012'
  const dummyRegion = 'us-east-1'
  const dummyThumbprint = 'abcdefgh1234567890'

  it('creates a role with the correct name', () => {
    const app = new App()
    const stack = new Stack(app, 'TestStack', { env: { account: dummyAccount, region: dummyRegion } })

    createGitHubActionsDeploymentRole(stack, dummyThumbprint)

    const template = Template.fromStack(stack)
    template.hasResourceProperties('AWS::IAM::Role', {
      RoleName: 'github-actions-deployment-role'
    })
  })

  it('uses the provided thumbprint in the OIDC provider', () => {
    const app = new App()
    const stack = new Stack(app, 'TestStack', { env: { account: dummyAccount, region: dummyRegion } })

    createGitHubActionsDeploymentRole(stack, dummyThumbprint)

    const template = Template.fromStack(stack)
    template.hasResourceProperties('Custom::AWSCDKOpenIdConnectProvider', {
      ThumbprintList: Match.arrayWith([
        dummyThumbprint
      ])
    })
  })

  it('includes inline policy with sts:AssumeRole permission', () => {
    const app = new App()
    const stack = new Stack(app, 'TestStack', { env: { account: dummyAccount, region: dummyRegion } })

    createGitHubActionsDeploymentRole(stack, dummyThumbprint)

    const template = Template.fromStack(stack)
    template.hasResourceProperties('AWS::IAM::Role', {
      Policies: [{
        PolicyDocument: {
          Statement: Match.arrayWith([
            Match.objectLike({
              Action: "sts:AssumeRole",
              Effect: Effect.ALLOW,
              Resource: `arn:aws:iam::${dummyAccount}:role/cdk-*`
            })
          ])
        }
      }]
    })
  })

  it('includes inline policy with ssm:GetParameter permission', () => {
    const app = new App()
    const stack = new Stack(app, 'TestStack', { env: { account: dummyAccount, region: dummyRegion } })

    createGitHubActionsDeploymentRole(stack, dummyThumbprint)

    const template = Template.fromStack(stack)
    template.hasResourceProperties('AWS::IAM::Role', {
      Policies: [{
        PolicyDocument: {
          Statement: Match.arrayWith([
            Match.objectLike({
              Action: "ssm:GetParameter",
              Effect: Effect.ALLOW,
              Resource: `arn:aws:ssm:${dummyRegion}:${dummyAccount}:parameter/cdk-*`
            })
          ])
        }
      }]
    })
  })

  it('has the correct trust policy conditions for GitHub OIDC', () => {
    const app = new App()
    const stack = new Stack(app, 'TestStack', { env: { account: dummyAccount, region: dummyRegion } })

    createGitHubActionsDeploymentRole(stack, dummyThumbprint)

    const template = Template.fromStack(stack)
    template.hasResourceProperties('AWS::IAM::Role', {
      AssumeRolePolicyDocument: {
        Statement: Match.arrayWith([
          Match.objectLike({
            Condition: Match.objectLike({
              StringLike: Match.objectLike({
                'token.actions.githubusercontent.com:sub': 'repo:JoelBosch90/shuffle-showdown-2:*',
              })
            }),
            Effect: Effect.ALLOW,
            Principal: Match.objectLike({
              Federated: {
                "Ref": Match.stringLikeRegexp('GitHubOIDCProvider')
              }
            }),
            Action: "sts:AssumeRoleWithWebIdentity"
          })
        ])
      }
    })
  })
})