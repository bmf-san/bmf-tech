import { test, expect } from '@playwright/test';

// Note: pagination links in generated HTML use absolute prod URLs
// (e.g. https://bmf-tech.com/ja/page/2/) so we cannot click them during tests.
// Instead we navigate directly to paginated URLs and assert content.

// JA: 584 articles / 20 per page = 30 pages
const JA_LAST_PAGE = 30;

// ── JA root pagination ────────────────────────────────────────────────────────

test.describe('JA root pagination', () => {
  test('page 1 loads with articles', async ({ page }) => {
    await page.goto('/ja/');
    await expect(page.locator('a.card').first()).toBeVisible();
  });

  test('page 2 loads with articles', async ({ page }) => {
    await page.goto('/ja/page/2/');
    await expect(page.locator('a.card').first()).toBeVisible();
  });

  test('last page loads with articles', async ({ page }) => {
    await page.goto(`/ja/page/${JA_LAST_PAGE}/`);
    await expect(page.locator('a.card').first()).toBeVisible();
  });

  test('beyond last page returns 404', async ({ page }) => {
    const res = await page.goto(`/ja/page/${JA_LAST_PAGE + 1}/`);
    expect(res?.status()).toBe(404);
  });
});

// ── Tag pagination (/tags/golang/ has 3 pages) ────────────────────────────────

test.describe('Tag pagination /ja/tags/golang/', () => {
  test('page 1 loads with articles', async ({ page }) => {
    await page.goto('/ja/tags/golang/');
    await expect(page.locator('a.card').first()).toBeVisible();
  });

  test('page 2 loads with articles', async ({ page }) => {
    await page.goto('/ja/tags/golang/page/2/');
    await expect(page.locator('a.card').first()).toBeVisible();
  });

  test('page 3 (last) loads with articles', async ({ page }) => {
    await page.goto('/ja/tags/golang/page/3/');
    await expect(page.locator('a.card').first()).toBeVisible();
  });

  test('page 4 (beyond last) returns 404', async ({ page }) => {
    const res = await page.goto('/ja/tags/golang/page/4/');
    expect(res?.status()).toBe(404);
  });
});

// ── Category pagination (/categories/アーキテクチャ/ has 5 pages) ──────────────

test.describe('Category pagination /ja/categories/アーキテクチャ/', () => {
  const BASE = '/ja/categories/アーキテクチャ/';

  test('page 1 loads with articles', async ({ page }) => {
    await page.goto(BASE);
    await expect(page.locator('a.card').first()).toBeVisible();
  });

  test('page 2 loads with articles', async ({ page }) => {
    await page.goto(`${BASE}page/2/`);
    await expect(page.locator('a.card').first()).toBeVisible();
  });

  test('page 5 (last) loads with articles', async ({ page }) => {
    await page.goto(`${BASE}page/5/`);
    await expect(page.locator('a.card').first()).toBeVisible();
  });

  test('page 6 (beyond last) returns 404', async ({ page }) => {
    const res = await page.goto(`${BASE}page/6/`);
    expect(res?.status()).toBe(404);
  });
});
