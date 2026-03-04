import { test, expect } from '@playwright/test';

test.describe('日本語 nav link', () => {
  test('navigates from / to /ja/', async ({ page }) => {
    await page.goto('/');
    await page.getByRole('link', { name: '日本語' }).click();
    await expect(page).toHaveURL('/ja/');
    await expect(page.locator('html')).toHaveAttribute('lang', 'ja');
  });
});

// gohan generates per-tag and per-category pages but no /tags/ or /categories/ index.
// Test individual taxonomy pages that are known to have articles.
test.describe('Tag page /tags/Go/', () => {
  test('loads and has article list', async ({ page }) => {
    await page.goto('/tags/Go/');
    await expect(page.locator('ul.article-list li').first()).toBeVisible();
  });

  test('article links are present', async ({ page }) => {
    await page.goto('/tags/Go/');
    const links = page.locator('ul.article-list li a');
    expect(await links.count()).toBeGreaterThan(0);
  });
});

test.describe('Category page /categories/OS/', () => {
  test('loads and has article list', async ({ page }) => {
    await page.goto('/categories/OS/');
    await expect(page.locator('ul.article-list li').first()).toBeVisible();
  });

  test('article links are present', async ({ page }) => {
    await page.goto('/categories/OS/');
    const links = page.locator('ul.article-list li a');
    expect(await links.count()).toBeGreaterThan(0);
  });
});

test.describe('About page /about/', () => {
  test('loads and has article content', async ({ page }) => {
    await page.goto('/about/');
    await expect(page.locator('.article-content')).toBeVisible();
  });
});

test.describe('Archives /ja/ → archive pages', () => {
  test('archives index loads', async ({ page }) => {
    const res = await page.goto('/archives/');
    // archives may redirect or render a list
    expect([200, 301, 302]).toContain(res?.status());
  });
});
