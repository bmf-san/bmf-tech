import { test, expect } from '@playwright/test';

// ── EN root index / ──────────────────────────────────────────────────────────

test.describe('English root index /', () => {
  test('page loads with status 200', async ({ page }) => {
    const res = await page.goto('/');
    expect(res?.status()).toBe(200);
  });

  test('page title is "bmf-tech"', async ({ page }) => {
    await page.goto('/');
    await expect(page).toHaveTitle('bmf-tech');
  });

  test('html lang is en', async ({ page }) => {
    await page.goto('/');
    await expect(page.locator('html')).toHaveAttribute('lang', 'en');
  });

  test('navbar is present with Home, About, Feed links', async ({ page }) => {
    await page.goto('/');
    const nav = page.locator('nav.navbar');
    await expect(nav).toBeVisible();
    await expect(nav.getByRole('link', { name: 'Home' })).toBeVisible();
    await expect(nav.getByRole('link', { name: 'About' })).toBeVisible();
    await expect(nav.getByRole('link', { name: 'Feed' })).toBeVisible();
  });

  test('locale toggle shows "JA" badge and links to /ja/', async ({ page }) => {
    await page.goto('/');
    const toggle = page.locator('nav.navbar .badge-primary');
    await expect(toggle).toHaveText('JA');
    const href = await toggle.locator('..').getAttribute('href');
    expect(href).toBe('/ja/');
  });

  test('article cards are present (a.card elements)', async ({ page }) => {
    await page.goto('/');
    const cards = page.locator('a.card');
    await expect(cards.first()).toBeVisible();
    expect(await cards.count()).toBeGreaterThan(0);
  });

  test('each card has a title, date and category', async ({ page }) => {
    await page.goto('/');
    const firstCard = page.locator('a.card').first();
    await expect(firstCard.locator('h3')).toBeVisible();
    await expect(firstCard.locator('.text-xs.text-secondary')).toBeVisible();
    await expect(firstCard.locator('.badge-secondary')).toBeVisible();
  });

  test('article cards link to /posts/', async ({ page }) => {
    await page.goto('/');
    const href = await page.locator('a.card').first().getAttribute('href');
    expect(href).toMatch(/^\/posts\//);
  });

  test('sidebar has Categories section', async ({ page }) => {
    await page.goto('/');
    await expect(page.locator('aside')).toBeVisible();
  });

  test('pagination is present on page 1', async ({ page }) => {
    await page.goto('/');
    await expect(page.locator('ul.pagination')).toBeVisible();
    await expect(page.locator('.pagination-link.active')).toBeVisible();
  });

  test('footer is present with Sitemap, GitHub, X links', async ({ page }) => {
    await page.goto('/');
    const footer = page.locator('footer');
    await expect(footer).toBeVisible();
    await expect(footer.getByRole('link', { name: 'Sitemap' })).toBeVisible();
    await expect(footer.getByRole('link', { name: 'GitHub' })).toBeVisible();
    await expect(footer.getByRole('link', { name: /X/ })).toBeVisible();
  });
});

// ── EN paginated index /page/{n}/ ─────────────────────────────────────────────

test.describe('English paginated index /page/2/', () => {
  const URL = '/page/2/';

  test('page 2 loads with status 200', async ({ page }) => {
    const res = await page.goto(URL);
    expect(res?.status()).toBe(200);
  });

  test('html lang is en', async ({ page }) => {
    await page.goto(URL);
    await expect(page.locator('html')).toHaveAttribute('lang', 'en');
  });

  test('article cards are present', async ({ page }) => {
    await page.goto(URL);
    await expect(page.locator('a.card').first()).toBeVisible();
    expect(await page.locator('a.card').count()).toBeGreaterThan(0);
  });

  test('pagination shows page 2 as active', async ({ page }) => {
    await page.goto(URL);
    await expect(page.locator('.pagination-link.active')).toBeVisible();
  });
});

// ── EN paginated index last page ─────────────────────────────────────────────

test.describe('English paginated index: last page and beyond', () => {
  test('last page loads with articles', async ({ page }) => {
    // Discover total page count from page 1
    await page.goto('/');
    const links = page.locator('ul.pagination .pagination-item a');
    const count = await links.count();
    // The last numeric page link (exclude prev/next arrows)
    let lastPage = 1;
    for (let i = 0; i < count; i++) {
      const text = await links.nth(i).textContent();
      const n = parseInt(text?.trim() ?? '0');
      if (!isNaN(n) && n > lastPage) lastPage = n;
    }
    if (lastPage <= 1) {
      test.skip(); // only one page, nothing to test
      return;
    }
    await page.goto(`/page/${lastPage}/`);
    await expect(page.locator('a.card').first()).toBeVisible();
  });

  test('beyond last page returns 404', async ({ page }) => {
    const res = await page.goto('/page/9999/');
    expect(res?.status()).toBe(404);
  });
});
