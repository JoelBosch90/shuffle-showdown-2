import { isLocalEnvironment } from "./isLocalEnvironment";

describe('isLocalEnvironment', () => {
  it('returns true for localstack endpoint', () => {
    process.env.AWS_ENDPOINT_URL = 'http://localhost:4566';
    expect(isLocalEnvironment()).toBe(true);
  });

  it('returns true for localhost endpoint', () => {
    process.env.AWS_ENDPOINT_URL = 'http://localhost:3000';
    expect(isLocalEnvironment()).toBe(true);
  });

  it('returns false for non-local environment', () => {
    process.env.AWS_ENDPOINT_URL = 'https://api.example.com';
    expect(isLocalEnvironment()).toBe(false);
  });

  it('returns false when AWS_ENDPOINT_URL is not set', () => {
    delete process.env.AWS_ENDPOINT_URL;
    expect(isLocalEnvironment()).toBe(false);
  });
});