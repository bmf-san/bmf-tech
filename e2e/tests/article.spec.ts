import { test, expect } from '@playwright/test';

const ARTICLE_URL = '/ja/posts/2018-review-2019-goals/';

test.describe('Article page', () => {
  test('page loads successfully', async ({ page }) => {
    const res = await page.goto(ARTICLE_URL);
    expect(res?.status()).toBe(200);
  });

  test('html lang is ja', async ({ page }) => {
    await page.goto(ARTICLE_URL);
    await expect(page.locator('html')).toHaveAttribute('lang', 'ja');
  });

  test('h1 is visible', async ({ page }) => {
    await page.goto(ARTICLE_URL);
    // scope to article header to avoid strict-mode violation from markdown headings
    await expect(page.locator('article header h1').first()).toBeVisible();
  });

  test('article-meta shows date', async ({ page }) => {
    await page.goto(ARTICLE_URL);
    await expect(page.locator('.article-meta time')).toBeVisible();
  });

  test('article-content is present and not empty', async ({ page }) => {
    await page.goto(ARTICLE_URL);
    const content = page.locator('.article-content');
    await expect(content).toBeVisible();
    const text = await content.textContent();
    expect(text?.trim().length).toBeGreaterThan(0);
  });

  test('site nav is present on article page', async ({ page }) => {
    await page.goto(ARTICLE_URL);
    await expect(page.locator('.site-header')).toBeVisible();
    await expect(page.locator('.site-footer')).toBeVisible();
  });
});
