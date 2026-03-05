import { test, expect } from '@playwright/test';

// ── Nav links ────────────────────────────────────────────────────────────────

test.describe('Nav: 日本語 link', () => {
  test('navigates from / to /ja/', async ({ page }) => {
    await page.goto('/');
    await page.locator('.nav-links').getByRole('link', { name: '日本語' }).click();
    await expect(page).toHaveURL('/ja/');
    await expect(page.locator('html')).toHaveAttribute('lang', 'ja');
  });
});

// ── Tags ─────────────────────────────────────────────────────────────────────

// gohan does not generate /tags/index.html automatically; we create it as a
// page-type content file (content/en/tags.md, slug: tags).
test.describe('Tags index /tags/', () => {
  test('loads a proper HTML page (not a directory listing)', async ({ page }) => {
    await page.goto('/tags/');
    // page.html renders an .article-content wrapper
    await expect(page.locator('.article-content')).toBeVisible();
    // must NOT show the http-server directory listing header
    await expect(page.locator('body')).not.toContainText('Index of /');
  });

  test('contains tag links', async ({ page }) => {
    await page.goto('/tags/');
    const links = page.locator('.article-content a');
    expect(await links.count()).toBeGreaterThan(0);
  });
});

test.describe('Tag page /tags/Go/', () => {
  test('loads and has article list', async ({ page }) => {
    await page.goto('/tags/Go/');
    await expect(page.locator('ul.article-list li').first()).toBeVisible();
  });

  test('article links point to /ja/posts/ or /posts/', async ({ page }) => {
    await page.goto('/tags/Go/');
    const href = await page.locator('ul.article-list li a').first().getAttribute('href');
    expect(href).toMatch(/\/posts\//);
  });
});

test.describe('Tags index /ja/tags/', () => {
  test('loads a proper HTML page (not a directory listing)', async ({ page }) => {
    await page.goto('/ja/tags/');
    // page.html renders an .article-content wrapper
    await expect(page.locator('.article-content')).toBeVisible();
    await expect(page.locator('body')).not.toContainText('Index of /');
  });

  test('contains tag links', async ({ page }) => {
    await page.goto('/ja/tags/');
    const links = page.locator('.article-content a');
    expect(await links.count()).toBeGreaterThan(0);
  });
});

// ── Categories ───────────────────────────────────────────────────────────────

// Same as tags: created as a page-type content file (content/en/categories.md)
test.describe('Categories index /categories/', () => {
  test('loads a proper HTML page (not a directory listing)', async ({ page }) => {
    await page.goto('/categories/');
    await expect(page.locator('.article-content')).toBeVisible();
    await expect(page.locator('body')).not.toContainText('Index of /');
  });

  test('contains all 21 category links', async ({ page }) => {
    await page.goto('/categories/');
    const links = page.locator('.article-content a');
    expect(await links.count()).toBeGreaterThanOrEqual(21);
  });

  test('OS category link is present', async ({ page }) => {
    await page.goto('/categories/');
    await expect(page.locator('.article-content').getByRole('link', { name: 'OS' })).toBeVisible();
  });
});

test.describe('Categories index /ja/categories/', () => {
  test('loads a proper HTML page (not a directory listing)', async ({ page }) => {
    await page.goto('/ja/categories/');
    await expect(page.locator('.article-content')).toBeVisible();
    await expect(page.locator('body')).not.toContainText('Index of /');
  });

  test('contains category links', async ({ page }) => {
    await page.goto('/ja/categories/');
    const links = page.locator('.article-content a');
    expect(await links.count()).toBeGreaterThan(0);
  });
});

test.describe('Category page /ja/categories/OS/', () => {
  test('loads and has article list', async ({ page }) => {
    await page.goto('/ja/categories/OS/');
    await expect(page.locator('ul.article-list li').first()).toBeVisible();
  });

  test('article links are present', async ({ page }) => {
    await page.goto('/ja/categories/OS/');
    const links = page.locator('ul.article-list li a');
    expect(await links.count()).toBeGreaterThan(0);
  });
});

// ── Static pages ─────────────────────────────────────────────────────────────

test.describe('About page /about/', () => {
  test('loads successfully', async ({ page }) => {
    const res = await page.goto('/about/');
    expect(res?.status()).toBe(200);
  });

  test('has article-content', async ({ page }) => {
    await page.goto('/about/');
    await expect(page.locator('.article-content')).toBeVisible();
  });

  test('contains profile information', async ({ page }) => {
    await page.goto('/about/');
    await expect(page.locator('.article-content')).toContainText('Kenta Takeuchi');
  });
});

// ── Archives ─────────────────────────────────────────────────────────────────

test.describe('Archive page /archives/2024/03/', () => {
  test('loads and has article list', async ({ page }) => {
    await page.goto('/archives/2024/03/');
    await expect(page.locator('ul.article-list li').first()).toBeVisible();
  });

  test('articles are from 2024-03', async ({ page }) => {
    await page.goto('/archives/2024/03/');
    const firstDate = await page.locator('.article-date').first().textContent();
    expect(firstDate).toMatch(/^2024-03/);
  });
});
