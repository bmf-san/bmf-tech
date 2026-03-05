import { test, expect } from '@playwright/test';

test.describe('English root index /', () => {
  test('page title', async ({ page }) => {
    await page.goto('/');
    await expect(page).toHaveTitle('bmf-tech');
  });

  test('html lang is en', async ({ page }) => {
    await page.goto('/');
    await expect(page.locator('html')).toHaveAttribute('lang', 'en');
  });

  test('all navigation links are present', async ({ page }) => {
    await page.goto('/');
    const nav = page.locator('.nav-links');
    // Scope to nav to avoid duplicates from article list items
    await expect(nav.getByRole('link', { name: '日本語' })).toBeVisible();
    await expect(nav.getByRole('link', { name: 'Tags' })).toBeVisible();
    await expect(nav.getByRole('link', { name: 'Categories' })).toBeVisible();
    await expect(nav.getByRole('link', { name: 'About' })).toBeVisible();
    await expect(nav.getByRole('link', { name: 'Feed' })).toBeVisible();
  });

  test('article list has items', async ({ page }) => {
    await page.goto('/');
    const items = page.locator('ul.article-list li');
    await expect(items.first()).toBeVisible();
    expect(await items.count()).toBeGreaterThan(0);
  });

  test('pagination info is visible', async ({ page }) => {
    await page.goto('/');
    await expect(page.locator('.pagination-info')).toContainText('/');
  });

});

test.describe('English article /posts/hello-world/', () => {
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
});
