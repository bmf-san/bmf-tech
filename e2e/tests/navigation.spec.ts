import { test, expect } from '@playwright/test';

// ── Nav links ────────────────────────────────────────────────────────────────

test.describe('Nav: locale toggle', () => {
  test('EN page shows JA toggle linking to /ja/', async ({ page }) => {
    await page.goto('/');
    const toggle = page.locator('nav.navbar a:has(.badge-primary)');
    await expect(toggle).toHaveText('JA');
    const href = await toggle.getAttribute('href');
    expect(href).toBe('/ja/');
  });

  test('JA page shows EN toggle linking to /', async ({ page }) => {
    await page.goto('/ja/');
    const toggle = page.locator('nav.navbar a:has(.badge-primary)');
    await expect(toggle).toHaveText('EN');
    const href = await toggle.getAttribute('href');
    expect(href).toBe('/');
  });

  test('clicking JA toggle navigates to /ja/', async ({ page }) => {
    await page.goto('/');
    await page.locator('nav.navbar a:has(.badge-primary)').click();
    await expect(page).toHaveURL('/ja/');
    await expect(page.locator('html')).toHaveAttribute('lang', 'ja');
  });

  test('JA article toggle links to EN translation when one exists', async ({ page }) => {
    await page.goto('/ja/posts/cto-thinking-strategy-leadership/');
    const toggle = page.locator('nav.navbar a:has(.badge-primary)');
    const href = await toggle.getAttribute('href');
    // Has EN translation → links to EN article
    expect(href).toBe('/posts/cto-thinking-strategy-leadership/');
  });

  test('EN article toggle links to JA translation when one exists', async ({ page }) => {
    await page.goto('/posts/cto-thinking-strategy-leadership/');
    const toggle = page.locator('nav.navbar a:has(.badge-primary)');
    const href = await toggle.getAttribute('href');
    // Has JA translation → links to JA article
    expect(href).toBe('/ja/posts/cto-thinking-strategy-leadership/');
  });
});

// ── Tags ─────────────────────────────────────────────────────────────────────

// gohan does not generate /tags/index.html automatically; we create it as a
// page-type content file (content/en/tags.md, slug: tags).
test.describe('Tags index /tags/', () => {
  test('loads a proper HTML page (not a directory listing)', async ({ page }) => {
    await page.goto('/tags/');
    // page.html renders an article.prose wrapper
    await expect(page.locator('article.prose')).toBeVisible();
    // must NOT show the http-server directory listing header
    await expect(page.locator('body')).not.toContainText('Index of /');
  });

  test('contains tag links', async ({ page }) => {
    await page.goto('/tags/');
    const links = page.locator('article.prose a');
    expect(await links.count()).toBeGreaterThan(0);
  });
});

test.describe('Tag page /tags/Golang/', () => {
  test('loads and has article list', async ({ page }) => {
    await page.goto('/tags/Golang/');
    await expect(page.locator('a.card').first()).toBeVisible();
  });

  test('article links point to /ja/posts/ or /posts/', async ({ page }) => {
    await page.goto('/tags/Golang/');
    const href = await page.locator('a.card').first().getAttribute('href');
    expect(href).toMatch(/\/posts\//);
  });
});

test.describe('Tags index /ja/tags/', () => {
  test('loads a proper HTML page (not a directory listing)', async ({ page }) => {
    await page.goto('/ja/tags/');
    // page.html renders an article.prose wrapper
    await expect(page.locator('article.prose')).toBeVisible();
    await expect(page.locator('body')).not.toContainText('Index of /');
  });

  test('contains tag links', async ({ page }) => {
    await page.goto('/ja/tags/');
    const links = page.locator('article.prose a');
    expect(await links.count()).toBeGreaterThan(0);
  });
});

// ── Categories ───────────────────────────────────────────────────────────────

// Same as tags: created as a page-type content file (content/en/categories.md)
test.describe('Categories index /categories/', () => {
  test('loads a proper HTML page (not a directory listing)', async ({ page }) => {
    await page.goto('/categories/');
    await expect(page.locator('article.prose')).toBeVisible();
    await expect(page.locator('body')).not.toContainText('Index of /');
  });

  test('contains category links', async ({ page }) => {
    await page.goto('/categories/');
    const links = page.locator('article.prose a');
    expect(await links.count()).toBeGreaterThan(0);
  });

  test('Tools category link is present', async ({ page }) => {
    await page.goto('/categories/');
    await expect(page.locator('article.prose').getByRole('link', { name: 'Tools' })).toBeVisible();
  });
});

test.describe('Categories index /ja/categories/', () => {
  test('loads a proper HTML page (not a directory listing)', async ({ page }) => {
    await page.goto('/ja/categories/');
    await expect(page.locator('article.prose')).toBeVisible();
    await expect(page.locator('body')).not.toContainText('Index of /');
  });

  test('contains category links', async ({ page }) => {
    await page.goto('/ja/categories/');
    const links = page.locator('article.prose a');
    expect(await links.count()).toBeGreaterThan(0);
  });
});

test.describe('Category page /ja/categories/OS/', () => {
  test('loads and has article list', async ({ page }) => {
    await page.goto('/ja/categories/OS/');
    await expect(page.locator('a.card').first()).toBeVisible();
  });

  test('article links are present', async ({ page }) => {
    await page.goto('/ja/categories/OS/');
    const links = page.locator('a.card');
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
    await expect(page.locator('article.prose')).toBeVisible();
  });

  test('contains profile information', async ({ page }) => {
    await page.goto('/about/');
    await expect(page.locator('article.prose')).toContainText('Kenta Takeuchi');
  });
});

test.describe('About page /ja/about/', () => {
  test('loads successfully', async ({ page }) => {
    const res = await page.goto('/ja/about/');
    expect(res?.status()).toBe(200);
  });

  test('has article-content', async ({ page }) => {
    await page.goto('/ja/about/');
    await expect(page.locator('article.prose')).toBeVisible();
  });

  test('contains Japanese profile text', async ({ page }) => {
    await page.goto('/ja/about/');
    await expect(page.locator('article.prose')).toContainText('シニアプラットフォームエンジニア');
  });

  test('JA nav About link points to /ja/about/', async ({ page }) => {
    await page.goto('/ja/');
    const href = await page.locator('nav.navbar').getByRole('link', { name: 'About' }).getAttribute('href');
    expect(href).toBe('/ja/about/');
  });
});

// ── Archives ─────────────────────────────────────────────────────────────────

test.describe('Archive page /archives/2024/03/', () => {
  test('loads and has article list', async ({ page }) => {
    await page.goto('/archives/2024/03/');
    await expect(page.locator('div.card-body li').first()).toBeVisible();
  });

  test('articles are from 2024-03', async ({ page }) => {
    await page.goto('/archives/2024/03/');
    const firstDate = await page.locator('div.card-body .text-xs.text-secondary').first().textContent();
    expect(firstDate).toMatch(/^2024-03/);
  });
});

// ── JA nav labels ─────────────────────────────────────────────────────────────

test.describe('Nav: JA locale shows Japanese labels', () => {
  test('JA nav shows ホーム and フィード', async ({ page }) => {
    await page.goto('/ja/');
    await expect(page.locator('nav.navbar').getByRole('link', { name: 'ホーム' })).toBeVisible();
    await expect(page.locator('nav.navbar').getByRole('link', { name: 'フィード' })).toBeVisible();
  });

  test('JA ホーム link points to /ja/', async ({ page }) => {
    await page.goto('/ja/');
    const href = await page.locator('nav.navbar').getByRole('link', { name: 'ホーム' }).getAttribute('href');
    expect(href).toBe('/ja/');
  });

  test('EN nav shows Home and Feed (not Japanese labels)', async ({ page }) => {
    await page.goto('/');
    await expect(page.locator('nav.navbar').getByRole('link', { name: 'Home' })).toBeVisible();
    await expect(page.locator('nav.navbar').getByRole('link', { name: 'Feed' })).toBeVisible();
  });
});

// ── JA tag URL on article ────────────────────────────────────────────────────

test.describe('JA article tag links use /ja/tags/ prefix', () => {
  test('tag badge href starts with /ja/tags/', async ({ page }) => {
    await page.goto('/ja/posts/cto-thinking-strategy-leadership/');
    const tagHref = await page.locator('a[href^="/ja/tags/"]').first().getAttribute('href');
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
