import { test, expect } from '@playwright/test';

// ── Japanese index /ja/ ───────────────────────────────────────────────────────

test.describe('Japanese index /ja/', () => {
  test('page loads with status 200', async ({ page }) => {
    const res = await page.goto('/ja/');
    expect(res?.status()).toBe(200);
  });

  test('html lang is ja', async ({ page }) => {
    await page.goto('/ja/');
    await expect(page.locator('html')).toHaveAttribute('lang', 'ja');
  });

  test('page title is "bmf-tech"', async ({ page }) => {
    await page.goto('/ja/');
    await expect(page).toHaveTitle('bmf-tech');
  });

  test('navbar is present with Japanese labels', async ({ page }) => {
    await page.goto('/ja/');
    const nav = page.locator('nav.navbar');
    await expect(nav).toBeVisible();
    await expect(nav.getByRole('link', { name: 'ホーム' })).toBeVisible();
    await expect(nav.getByRole('link', { name: 'About' })).toBeVisible();
    await expect(nav.getByRole('link', { name: 'フィード' })).toBeVisible();
  });

  test('locale toggle shows "EN" badge and links to /', async ({ page }) => {
    await page.goto('/ja/');
    const toggle = page.locator('nav.navbar .badge-primary');
    await expect(toggle).toHaveText('EN');
    const href = await toggle.locator('..').getAttribute('href');
    expect(href).toBe('/');
  });

  test('article cards are present', async ({ page }) => {
    await page.goto('/ja/');
    const cards = page.locator('a.card');
    await expect(cards.first()).toBeVisible();
    expect(await cards.count()).toBeGreaterThan(0);
  });

  test('article cards link to /ja/posts/', async ({ page }) => {
    await page.goto('/ja/');
    const href = await page.locator('a.card').first().getAttribute('href');
    expect(href).toMatch(/^\/ja\/posts\//);
  });

  test('pagination is visible with 584 articles (>1 page)', async ({ page }) => {
    await page.goto('/ja/');
    await expect(page.locator('ul.pagination')).toBeVisible();
    // Total pages > 1 (584 articles / 20 per page ≈ 30 pages)
    const items = page.locator('ul.pagination .pagination-item a');
    const count = await items.count();
    expect(count).toBeGreaterThan(1);
  });

  test('footer is present', async ({ page }) => {
    await page.goto('/ja/');
    await expect(page.locator('footer')).toBeVisible();
  });
});

// ── JA paginated index /ja/page/{n}/ ──────────────────────────────────────────

test.describe('Japanese paginated index /ja/page/2/', () => {
  const URL = '/ja/page/2/';

  test('page 2 loads with status 200', async ({ page }) => {
    const res = await page.goto(URL);
    expect(res?.status()).toBe(200);
  });

  test('html lang is ja', async ({ page }) => {
    await page.goto(URL);
    await expect(page.locator('html')).toHaveAttribute('lang', 'ja');
  });

  test('article cards are present', async ({ page }) => {
    await page.goto(URL);
    await expect(page.locator('a.card').first()).toBeVisible();
  });

  test('pagination shows page 2 as active', async ({ page }) => {
    await page.goto(URL);
    await expect(page.locator('.pagination-link.active')).toBeVisible();
  });
});

// ── JA last/beyond page ───────────────────────────────────────────────────────

test.describe('JA paginated index: last page and beyond', () => {
  test('last page loads with articles', async ({ page }) => {
    await page.goto('/ja/');
    const links = page.locator('ul.pagination .pagination-item a');
    const count = await links.count();
    let lastPage = 1;
    for (let i = 0; i < count; i++) {
      const text = await links.nth(i).textContent();
      const n = parseInt(text?.trim() ?? '0');
      if (!isNaN(n) && n > lastPage) lastPage = n;
    }
    await page.goto(`/ja/page/${lastPage}/`);
    await expect(page.locator('a.card').first()).toBeVisible();
  });

  test('beyond last page returns 404', async ({ page }) => {
    const res = await page.goto('/ja/page/9999/');
    expect(res?.status()).toBe(404);
  });
});
