import { test, expect } from '@playwright/test';

test.describe('Japanese index /ja/', () => {
  test('html lang is ja', async ({ page }) => {
    await page.goto('/ja/');
    await expect(page.locator('html')).toHaveAttribute('lang', 'ja');
  });

  test('article list has items', async ({ page }) => {
    await page.goto('/ja/');
    const items = page.locator('ul.article-list li');
    await expect(items.first()).toBeVisible();
    expect(await items.count()).toBeGreaterThan(0);
  });

  test('pagination shows multiple pages (584 articles / 20 per page)', async ({ page }) => {
    await page.goto('/ja/');
    const info = page.locator('.pagination-info');
    await expect(info).toContainText('1 /');
    const text = await info.textContent();
    const totalPages = parseInt(text!.split('/')[1]?.trim() ?? '0');
    expect(totalPages).toBeGreaterThan(1);
  });

  test('page 2 has articles', async ({ page }) => {
    // navigate directly to page 2 — pagination links use absolute prod URLs so clicking them
    // would leave localhost; assert content instead
    await page.goto('/ja/page/2/');
    await expect(page.locator('ul.article-list li').first()).toBeVisible();
    await expect(page.locator('.pagination-info')).toContainText('2 /');
  });

  test('article titles link to article pages', async ({ page }) => {
    await page.goto('/ja/');
    const firstLink = page.locator('ul.article-list li a').first();
    const href = await firstLink.getAttribute('href');
    expect(href).toMatch(/^\/ja\/posts\//);
  });
});
