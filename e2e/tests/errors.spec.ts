import { test, expect } from '@playwright/test';

// ── 404 responses ─────────────────────────────────────────────────────────────
// Regression for gohan#71: non-HTML 404 responses were sent with HTTP 200
// because injectingResponseWriter never forwarded the status code.

test.describe('Non-existent pages return 404', () => {
  const MISSING_URLS = [
    '/this-page-does-not-exist/',
    '/ja/page/9999/',
    '/ja/tags/NonExistentTag/',
    '/ja/categories/NonExistentCategory/',
    '/posts/non-existent-post/',
  ];

  for (const url of MISSING_URLS) {
    test(`${url} returns 404`, async ({ page }) => {
      const res = await page.goto(url);
      expect(res?.status()).toBe(404);
    });
  }
});
