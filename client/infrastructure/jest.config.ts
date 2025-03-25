import type { Config } from '@jest/types';

const config: Config.InitialOptions = {
  preset: 'ts-jest',
  testEnvironment: 'node',
  roots: ['<rootDir>/source'],
  testMatch: ['**/*.test.ts'],
  collectCoverage: true,
  coverageDirectory: "coverage",
  collectCoverageFrom: [
    '**/*.{ts,tsx}',
    '!**/bin/**',
    '!**/node_modules/**',
  ],
  coverageThreshold: {
    global: {
      branches: 100,
      functions: 100,
      lines: 100,
      statements: 100
    }
  },
};

export default config;
