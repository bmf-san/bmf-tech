import { test, expect } from '@playwright/test';

// ── Nav links ────────────────────────────────────────────────────────────────

test.describe('Nav: locale toggle', () => {
  test('EN page shows JA toggle linking to /ja/', async ({ page }) => {
    await page.goto('/');
    const toggle = page.locator('.nav-links .locale-toggle');
    await expect(toggle).toHaveText('JA');
    const href = await toggle.getAttribute('href');
    expect(href).toBe('/ja/');
  });

  test('JA page shows EN toggle linking to /', async ({ page }) => {
    await page.goto('/ja/');
    const toggle = page.locator('.nav-links .locale-toggle');
    await expect(toggle).toHaveText('EN');
    const href = await toggle.getAttribute('href');
    expect(href).toBe('/');
  });

  test('clicking JA toggle navigates to /ja/', async ({ page }) => {
    await page.goto('/');
    await page.locator('.nav-links .locale-toggle').click();
    await expect(page).toHaveURL('/ja/');
    await expect(page.locator('html')).toHaveAttribute('lang', 'ja');
  });

  test('JA article toggle falls back to / when no EN translation', async ({ page }) => {
    await page.goto('/ja/posts/cto-thinking-strategy-leadership/');
    const toggle = page.locator('.nav-links .locale-toggle');
    const href = await toggle.getAttribute('href');
    // No EN translation → falls back to /
    expect(href).toBe('/');
  });

  test('EN article /posts/hello-world/ toggle links to /ja/', async ({ page }) => {
    await page.goto('/posts/hello-world/');
    const toggle = page.locator('.nav-links .locale-toggle');
    const href = await toggle.getAttribute('href');
    // No JA translation for this article → falls back to /ja/
    expect(href).toBe('/ja/');
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

  test('contains category links', async ({ page }) => {
    await page.goto('/categories/');
    const links = page.locator('.article-content a');
    expect(await links.count()).toBeGreaterThan(0);
  });

  test('Tools category link is present', async ({ page }) => {
    await page.goto('/categories/');
    await expect(page.locator('.article-content').getByRole('link', { name: 'Tools' })).toBeVisible();
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

test.describe('About page /ja/about/', () => {
  test('loads successfully', async ({ page }) => {
    const res = await page.goto('/ja/about/');
    expect(res?.status()).toBe(200);
  });

  test('has article-content', async ({ page }) => {
    await page.goto('/ja/about/');
    await expect(page.locator('.article-content')).toBeVisible();
  });

  test('contains Japanese profile text', async ({ page }) => {
    await page.goto('/ja/about/');
    await expect(page.locator('.article-content')).toContainText('シニアプラットフォームエンジニア');
  });

  test('JA nav About link points to /ja/about/', async ({ page }) => {
    await page.goto('/ja/');
    const href = await page.locator('.nav-links').getByRole('link', { name: 'About' }).getAttribute('href');
    expect(href).toBe('/ja/about/');
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

// ── JA nav labels ─────────────────────────────────────────────────────────────

test.describe('Nav: JA locale shows Japanese labels', () => {
  test('nav shows タグ and カテゴリ on /ja/', async ({ page }) => {
    await page.goto('/ja/');
    await expect(page.locator('.nav-links').getByRole('link', { name: 'タグ' })).toBeVisible();
    await expect(page.locator('.nav-links').getByRole('link', { name: 'カテゴリ' })).toBeVisible();
  });

  test('タグ link points to /ja/tags/', async ({ page }) => {
    await page.goto('/ja/');
    const href = await page.locator('.nav-links').getByRole('link', { name: 'タグ' }).getAttribute('href');
    expect(href).toBe('/ja/tags/');
  });

  test('EN nav does not show タグ/カテゴリ', async ({ page }) => {
    await page.goto('/');
    await expect(page.locator('.nav-links').getByRole('link', { name: 'Tags' })).toBeVisible();
    await expect(page.locator('.nav-links').getByRole('link', { name: 'タグ' })).toHaveCount(0);
  });
});

// ── JA tag URL on article ────────────────────────────────────────────────────

test.describe('JA article tag links use /ja/tags/ prefix', () => {
  test('tag badge href starts with /ja/tags/', async ({ page }) => {
    await page.goto('/ja/posts/cto-thinking-strategy-leadership/');
    const tagHref = await page.locator('.badge-secondary').first().getAttribute('href');
    expect(tagHref).toMatch(/^\/ja\/tags\//);
  });

  test('/ja/tags/CTO/ returns 200', async ({ page }) => {
    const res = await page.goto('/ja/tags/CTO/');
    expect(res?.status()).toBe(200);
  });
});

// ── Sitemap in footer ────────────────────────────────────────────────────────

test.describe('Footer sitemap link', () => {
  test('EN footer has sitemap link', async ({ page }) => {
    await page.goto('/');
    await expect(page.locator('footer').getByRole('link', { name: 'Sitemap' })).toBeVisible();
  });

  test('sitemap link points to /sitemap.xml', async ({ page }) => {
    await page.goto('/');
    const href = await page.locator('footer').getByRole('link', { name: 'Sitemap' }).getAttribute('href');
    expect(href).toBe('/sitemap.xml');
  });

  test('JA footer also has sitemap link', async ({ page }) => {
    await page.goto('/ja/');
    await expect(page.locator('footer').getByRole('link', { name: 'Sitemap' })).toBeVisible();
  });
});
