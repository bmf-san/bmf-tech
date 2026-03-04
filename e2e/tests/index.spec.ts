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

  test('navigation links are visible', async ({ page }) => {
    await page.goto('/');
    await expect(page.getByRole('link', { name: '日本語' })).toBeVisible();
    await expect(page.getByRole('link', { name: 'Tags' })).toBeVisible();
    await expect(page.getByRole('link', { name: 'Categories' })).toBeVisible();
    // scope to nav-links to avoid collision with "About" article in the list
    await expect(page.locator('.nav-links').getByRole('link', { name: 'About' })).toBeVisible();
    await expect(page.getByRole('link', { name: 'Feed' })).toBeVisible();
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
