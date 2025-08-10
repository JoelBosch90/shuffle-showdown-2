/**
 *  This function retrieves the API path from environment variables and returns an HttpOrigin.
 *  @throws   {Error} - If the SHUFFLE_SHOWDOWN_API_PATH environment variable is not defined.
 *  @returns  {string} - The created API origin.
 */
export const getApiPath = (): string => {
  const path = process.env.SHUFFLE_SHOWDOWN_API_PATH;
  if (!path) {
    throw new Error('API path is not defined');
  }

  return path;
};