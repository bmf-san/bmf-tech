import { test, expect } from '@playwright/test';

// Note: pagination links in generated HTML use absolute prod URLs
// (e.g. https://bmf-tech.com/ja/page/2/) so we cannot click them during tests.
// Instead we navigate directly to paginated URLs and assert content.

// ── JA root pagination ────────────────────────────────────────────────────────

test.describe('JA root pagination', () => {
  test('page 1 shows pagination-info "1 / N"', async ({ page }) => {
    await page.goto('/ja/');
    const info = page.locator('.pagination-info');
    await expect(info).toBeVisible();
    const text = await info.textContent();
    expect(text).toMatch(/^1 \/ \d+/);
    // Total pages must be > 1 (584 JA articles / 20 per page ≈ 30 pages)
    const total = parseInt(text!.split('/')[1]?.trim() ?? '0');
    expect(total).toBeGreaterThan(1);
  });

  test('page 2 loads with articles', async ({ page }) => {
    await page.goto('/ja/page/2/');
    await expect(page.locator('ul.article-list li').first()).toBeVisible();
    await expect(page.locator('.pagination-info')).toContainText('2 /');
  });

  test('last page loads with articles', async ({ page }) => {
    // Get total pages from page 1 first
    await page.goto('/ja/');
    const text = await page.locator('.pagination-info').textContent();
    const total = text!.split('/')[1]?.trim();
    await page.goto(`/ja/page/${total}/`);
    await expect(page.locator('ul.article-list li').first()).toBeVisible();
    await expect(page.locator('.pagination-info')).toContainText(`${total} /`);
  });

  test('beyond last page returns 404', async ({ page }) => {
    await page.goto('/ja/');
    const text = await page.locator('.pagination-info').textContent();
    const total = parseInt(text!.split('/')[1]?.trim() ?? '0');
    const res = await page.goto(`/ja/page/${total + 1}/`);
    expect(res?.status()).toBe(404);
  });
});

// ── Tag pagination (/tags/Golang/ has 3 pages) ────────────────────────────────

test.describe('Tag pagination /tags/Golang/', () => {
  test('page 1 shows pagination-info "1 / 3"', async ({ page }) => {
    await page.goto('/tags/Golang/');
    await expect(page.locator('.pagination-info')).toContainText('1 / 3');
  });

  test('page 2 loads with articles', async ({ page }) => {
    await page.goto('/tags/Golang/page/2/');
    await expect(page.locator('ul.article-list li').first()).toBeVisible();
    await expect(page.locator('.pagination-info')).toContainText('2 / 3');
  });

  test('page 3 (last) loads with articles', async ({ page }) => {
    await page.goto('/tags/Golang/page/3/');
    await expect(page.locator('ul.article-list li').first()).toBeVisible();
    await expect(page.locator('.pagination-info')).toContainText('3 / 3');
  });

  test('page 4 (beyond last) returns 404', async ({ page }) => {
    const res = await page.goto('/tags/Golang/page/4/');
    expect(res?.status()).toBe(404);
  });
});

// ── Category pagination (/categories/アーキテクチャ/ has 5 pages) ──────────────

test.describe('Category pagination /categories/アーキテクチャ/', () => {
  const BASE = '/categories/アーキテクチャ/';

  test('page 1 shows pagination-info "1 / 5"', async ({ page }) => {
    await page.goto(BASE);
    await expect(page.locator('.pagination-info')).toContainText('1 / 5');
  });

  test('page 2 loads with articles', async ({ page }) => {
    await page.goto(`${BASE}page/2/`);
    await expect(page.locator('ul.article-list li').first()).toBeVisible();
    await expect(page.locator('.pagination-info')).toContainText('2 / 5');
  });

  test('page 5 (last) loads with articles', async ({ page }) => {
    await page.goto(`${BASE}page/5/`);
    await expect(page.locator('ul.article-list li').first()).toBeVisible();
    await expect(page.locator('.pagination-info')).toContainText('5 / 5');
  });

  test('page 6 (beyond last) returns 404', async ({ page }) => {
    const res = await page.goto(`${BASE}page/6/`);
    expect(res?.status()).toBe(404);
  });
});
