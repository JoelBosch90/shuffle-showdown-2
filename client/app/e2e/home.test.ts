import { expect, test } from '@playwright/test';
import { config } from 'dotenv';
import { resolve, dirname } from 'path';
import { fileURLToPath } from 'url';

const __dirname = dirname(fileURLToPath(import.meta.url));
config({ path: resolve(__dirname, '../../../.env') });

const apiPath = process.env.PUBLIC_SHUFFLE_SHOWDOWN_API_PATH || 'api';

// TODO: actually test this against the server without mocking.
test.describe('Home page', () => {
  const helloEndpoint = `/${apiPath}/hello`;
  const expectedLoadingMessage = 'loading message...';

  test('shows loading message initially and then updates with server response', async ({ page }) => {
    const message = 'Hello from mocked server!';
    await page.route(helloEndpoint, async route => {
      await route.fulfill({
        status: 200,
        contentType: 'text/plain',
        body: message,
      });
    });

    await page.goto('/');

    await expect(page.locator('p')).toContainText(expectedLoadingMessage);
    await expect(page.locator('p')).toContainText(message);
  });

  test('handles API errors gracefully', async ({ page }) => {
    await page.route(helloEndpoint, async route => {
      await route.fulfill({
        status: 500,
        contentType: 'text/plain',
        body: 'Internal Server Error'
      });
    });

    await page.goto('/');

    await expect(page.locator('p')).toContainText(expectedLoadingMessage);
    await page.waitForTimeout(500);
    await expect(page.locator('p')).toContainText(expectedLoadingMessage);
  });

  test('handles network failures', async ({ page }) => {
    await page.route(helloEndpoint, async route => {
      await route.abort('failed');
    });

    await page.goto('/');

    await expect(page.locator('p')).toContainText(expectedLoadingMessage);
    await page.waitForTimeout(500);
    await expect(page.locator('p')).toContainText(expectedLoadingMessage);
  });
});