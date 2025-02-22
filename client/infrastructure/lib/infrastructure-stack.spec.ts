import * as cdk from 'aws-cdk-lib'
import { InfrastructureStack } from './infrastructure-stack'

describe('InfrastructureStack', () => {
  it('creates a stack', () => {
    const app = new cdk.App()
    const stackName = 'ClientStack'
    const stack = new InfrastructureStack(app, stackName)

    expect(stack).toBeDefined()
  });
});