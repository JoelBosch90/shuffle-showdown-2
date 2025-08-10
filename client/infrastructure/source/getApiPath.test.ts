import { getApiPath } from "./getApiPath";

describe('getApiPath', () => {
  it('throws an error if the environment variable is not defined', () => {
    delete process.env.SHUFFLE_SHOWDOWN_API_PATH;

    expect(() => getApiPath()).toThrow('API path is not defined');
  });

  it('creates an HttpOrigin with the correct URL', () => {
    process.env.SHUFFLE_SHOWDOWN_API_PATH = 'api';

    const path = getApiPath();

    expect(path).toBe('api');
  });
});