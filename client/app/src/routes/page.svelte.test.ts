import { describe, test, expect, vi, beforeEach, afterEach } from 'vitest';
import '@testing-library/jest-dom/vitest';
import { render, screen, waitFor } from '@testing-library/svelte';
import Page from './+page.svelte';

const { mockStage, mockApiPath } = vi.hoisted(() => ({
  mockStage: 'test',
  mockApiPath: 'api',
}));

vi.mock('$env/static/public', () => ({
  PUBLIC_STAGE: mockStage,
  PUBLIC_SHUFFLE_SHOWDOWN_API_PATH: mockApiPath
}));

describe('/+page.svelte', () => {
  const originalFetch = global.fetch;

  beforeEach(() => {
    global.fetch = vi.fn();
  });

  afterEach(() => {
    global.fetch = originalFetch;
  });

  test('should show loading message initially', () => {
    const message = 'loading message...';
    global.fetch = vi.fn().mockResolvedValue({
      ok: true,
      text: () => Promise.resolve(message)
    });

    render(Page);

    expect(screen.getByText(message)).toBeInTheDocument();
  });

  test('should fetch data and update message on successful response', async () => {
    const message = 'Hello from the server!';
    global.fetch = vi.fn().mockResolvedValue({
      ok: true,
      text: () => Promise.resolve(message)
    });

    render(Page);

    expect(screen.getByText('loading message...')).toBeInTheDocument();
    await waitFor(() => expect(screen.getByText(message)).toBeInTheDocument());
    expect(global.fetch).toHaveBeenCalledWith(`/${mockStage}/${mockApiPath}/hello`);
  });

  test('should keep loading message on failed fetch', async () => {
    global.fetch = vi.fn().mockResolvedValue({
      ok: false,
      status: 404,
      statusText: 'Not Found'
    });

    render(Page);

    await waitFor(() => {
      expect(screen.getByText('loading message...')).toBeInTheDocument();
      expect(global.fetch).toHaveBeenCalledWith(`/${mockStage}/${mockApiPath}/hello`);
    });
  });

  test('should handle fetch errors gracefully', async () => {
    global.fetch = vi.fn().mockRejectedValue(new Error('Network error'));

    render(Page);

    await waitFor(() => {
      expect(screen.getByText('loading message...')).toBeInTheDocument();
    });
  });
});