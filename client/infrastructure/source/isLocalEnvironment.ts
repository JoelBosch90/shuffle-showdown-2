export const isLocalEnvironment = (): boolean => (
  process.env.AWS_ENDPOINT_URL?.includes('localhost') || process.env.AWS_ENDPOINT_URL?.includes('localstack')
) ?? false;