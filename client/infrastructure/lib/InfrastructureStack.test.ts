import { App, Stack, Aws } from 'aws-cdk-lib'
import { InfrastructureStack } from './InfrastructureStack'

const createWebsiteBucketSpy = jest.spyOn(require('./createWebsiteBucket'), 'createWebsiteBucket')
const createGitHubActionsDeploymentRoleSpy = jest.spyOn(require('./createGitHubActionsDeploymentRole'), 'createGitHubActionsDeploymentRole')

describe('InfrastructureStack', () => {
  const dummyAccount = '123456789012'
  const dummyRegion = 'us-east-1'
  const dummyStackProperties = {
    env: {
      account: dummyAccount,
      region: dummyRegion,
    },
  }
  Object.defineProperty(Aws, 'ACCOUNT_ID', { value: dummyAccount, writable: false })
  Object.defineProperty(Aws, 'REGION', { value: dummyRegion, writable: false })

  it('creates a website bucket with the proper name', () => {
    const app = new App()
    new InfrastructureStack(app, 'TestStack', dummyStackProperties)

    expect(createWebsiteBucketSpy).toHaveBeenCalledWith(
      expect.any(Stack),
      `clientstack-website-${dummyAccount}-${dummyRegion}`,
    )
  })

  it('creates a GitHub Actions deployment role', () => {
    const app = new App()
    new InfrastructureStack(app, 'TestStack', dummyStackProperties)

    expect(createGitHubActionsDeploymentRoleSpy).toHaveBeenCalledWith(
      expect.any(Stack),
      `clientstack-website-${dummyAccount}-${dummyRegion}`,
    )
  })
})