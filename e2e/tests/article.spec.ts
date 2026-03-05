import { test, expect } from '@playwright/test';

// ── Japanese article ──────────────────────────────────────────────────────────

test.describe('Japanese article page', () => {
  const URL = '/ja/posts/2018-review-2019-goals/';

  test('page loads with status 200', async ({ page }) => {
    const res = await page.goto(URL);
    expect(res?.status()).toBe(200);
  });

  test('html lang is ja', async ({ page }) => {
    await page.goto(URL);
    await expect(page.locator('html')).toHaveAttribute('lang', 'ja');
  });

  test('article title h1 is visible', async ({ page }) => {
    await page.goto(URL);
    // scope to article header to avoid strict-mode violation from markdown headings
    await expect(page.locator('article header h1').first()).toBeVisible();
  });

  test('article-meta shows date', async ({ page }) => {
    await page.goto(URL);
    await expect(page.locator('.article-meta time')).toBeVisible();
  });

  test('article-content is present and not empty', async ({ page }) => {
    await page.goto(URL);
    const content = page.locator('.article-content');
    await expect(content).toBeVisible();
    const text = await content.textContent();
    expect(text?.trim().length).toBeGreaterThan(0);
  });

  test('site header and footer are present', async ({ page }) => {
    await page.goto(URL);
    await expect(page.locator('.site-header')).toBeVisible();
    await expect(page.locator('.site-footer')).toBeVisible();
  });
});

// ── English article ──────────────────────────────────────────────────────────

test.describe('English article page /posts/hello-world/', () => {
  const URL = '/posts/hello-world/';

  test('page loads with status 200', async ({ page }) => {
    const res = await page.goto(URL);
    expect(res?.status()).toBe(200);
  });

  test('html lang is en', async ({ page }) => {
    await page.goto(URL);
    await expect(page.locator('html')).toHaveAttribute('lang', 'en');
  });

  test('article title h1 is visible', async ({ page }) => {
    await page.goto(URL);
    await expect(page.locator('article header h1').first()).toBeVisible();
  });

  test('article-content is present and not empty', async ({ page }) => {
    await page.goto(URL);
    const content = page.locator('.article-content');
    await expect(content).toBeVisible();
    const text = await content.textContent();
    expect(text?.trim().length).toBeGreaterThan(0);
  });

  test('site header and footer are present', async ({ page }) => {
    await page.goto(URL);
    await expect(page.locator('.site-header')).toBeVisible();
    await expect(page.locator('.site-footer')).toBeVisible();
  });
});
