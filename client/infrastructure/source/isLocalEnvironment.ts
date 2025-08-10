/**
 *  This function checks if the current environment is local by examining the AWS_ENDPOINT_URL environment variable.
 *  @returns  {boolean} - True if the environment is local, false otherwise.
 */
export const isLocalEnvironment = (): boolean => (
  process.env.AWS_ENDPOINT_URL?.includes('localhost') || process.env.AWS_ENDPOINT_URL?.includes('localstack')
) ?? false;