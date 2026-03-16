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

test.describe('Tag page /tags/golang/', () => {
  test('loads and has article list', async ({ page }) => {
    await page.goto('/tags/golang/');
    await expect(page.locator('a.card').first()).toBeVisible();
  });

  test('article links point to /ja/posts/ or /posts/', async ({ page }) => {
    await page.goto('/tags/golang/');
    const href = await page.locator('a.card').first().getAttribute('href');
    expect(href).toMatch(/\/posts\//);
  });
});

// ── Categories ───────────────────────────────────────────────────────────────

test.describe('Category page /ja/categories/os/', () => {
  test('loads and has article list', async ({ page }) => {
    await page.goto('/ja/categories/os/');
    await expect(page.locator('a.card').first()).toBeVisible();
  });

  test('article links are present', async ({ page }) => {
    await page.goto('/ja/categories/os/');
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
    await expect(page.locator('a.card').first()).toBeVisible();
  });

  test('articles are from 2024-03', async ({ page }) => {
    await page.goto('/archives/2024/03/');
    const firstDate = await page.locator('div.card-body .text-xs.text-secondary').first().textContent();
    expect(firstDate).toMatch(/^2024-03/);
  });
});

test.describe('Archive page /ja/archives/2026/02/', () => {
  test('loads and has article list', async ({ page }) => {
    await page.goto('/ja/archives/2026/02/');
    await expect(page.locator('a.card').first()).toBeVisible();
  });

  test('articles are from 2026-02', async ({ page }) => {
    await page.goto('/ja/archives/2026/02/');
    const firstDate = await page.locator('div.card-body .text-xs.text-secondary').first().textContent();
    expect(firstDate).toMatch(/^2026-02/);
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

  test('/ja/tags/cto/ returns 200', async ({ page }) => {
    const res = await page.goto('/ja/tags/cto/');
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
    await expect(page.locator('footer').getByRole('link', { name: 'サイトマップ' })).toBeVisible();
  });
});

// ── Locale switcher on tag/category/archive pages ────────────────────────────

test.describe('Nav: locale toggle on tag pages', () => {
  test('JA tag page toggle links to EN tag page', async ({ page }) => {
    await page.goto('/ja/tags/2phase-commit/');
    const href = await page.locator('nav.navbar a:has(.badge-primary)').getAttribute('href');
    expect(href).toBe('/tags/2phase-commit/');
  });

  test('EN tag page toggle links to JA tag page', async ({ page }) => {
    await page.goto('/tags/2phase-commit/');
    const href = await page.locator('nav.navbar a:has(.badge-primary)').getAttribute('href');
    expect(href).toBe('/ja/tags/2phase-commit/');
  });
});

test.describe('Nav: locale toggle on category pages', () => {
  test('JA category page toggle links to EN category page', async ({ page }) => {
    await page.goto('/ja/categories/os/');
    const href = await page.locator('nav.navbar a:has(.badge-primary)').getAttribute('href');
    expect(href).toBe('/categories/os/');
  });

  test('EN category page toggle links to JA category page', async ({ page }) => {
    await page.goto('/categories/architecture/');
    const href = await page.locator('nav.navbar a:has(.badge-primary)').getAttribute('href');
    expect(href).toBe('/ja/categories/architecture/');
  });
});

test.describe('Nav: locale toggle on archive pages', () => {
  test('JA archive month page toggle links to EN archive month page', async ({ page }) => {
    await page.goto('/ja/archives/2024/01/');
    const href = await page.locator('nav.navbar a:has(.badge-primary)').getAttribute('href');
    expect(href).toBe('/archives/2024/01/');
  });

  test('EN archive month page toggle links to JA archive month page', async ({ page }) => {
    await page.goto('/archives/2024/01/');
    const href = await page.locator('nav.navbar a:has(.badge-primary)').getAttribute('href');
    expect(href).toBe('/ja/archives/2024/01/');
  });

  test('JA archive year page toggle links to EN archive year page', async ({ page }) => {
    await page.goto('/ja/archives/2024/');
    const href = await page.locator('nav.navbar a:has(.badge-primary)').getAttribute('href');
    expect(href).toBe('/archives/2024/');
  });

  test('EN archive year page toggle links to JA archive year page', async ({ page }) => {
    await page.goto('/archives/2024/');
    const href = await page.locator('nav.navbar a:has(.badge-primary)').getAttribute('href');
    expect(href).toBe('/ja/archives/2024/');
  });
});

// ── Locale toggle on paginated index pages ───────────────────────────────────

test.describe('Nav: locale toggle on paginated root index pages', () => {
  test('EN /page/2/ toggle links to /ja/page/2/', async ({ page }) => {
    await page.goto('/page/2/');
    const href = await page.locator('nav.navbar a:has(.badge-primary)').getAttribute('href');
    expect(href).toBe('/ja/page/2/');
  });

  test('JA /ja/page/2/ toggle links to /page/2/', async ({ page }) => {
    await page.goto('/ja/page/2/');
    const href = await page.locator('nav.navbar a:has(.badge-primary)').getAttribute('href');
    expect(href).toBe('/page/2/');
  });
});

// ── Locale toggle on paginated tag/category pages ───────────────────────────

test.describe('Nav: locale toggle on paginated tag pages', () => {
  test('EN /tags/golang/page/2/ toggle links to /ja/tags/golang/page/2/', async ({ page }) => {
    await page.goto('/tags/golang/page/2/');
    const href = await page.locator('nav.navbar a:has(.badge-primary)').getAttribute('href');
    expect(href).toBe('/ja/tags/golang/page/2/');
  });

  test('JA /ja/tags/golang/page/2/ toggle links to /tags/golang/page/2/', async ({ page }) => {
    await page.goto('/ja/tags/golang/page/2/');
    const href = await page.locator('nav.navbar a:has(.badge-primary)').getAttribute('href');
    expect(href).toBe('/tags/golang/page/2/');
  });
});

// ── Sidebar locale filtering ─────────────────────────────────────────────────
// Regression: localeTaxonomyBase must not leak cross-locale links into the sidebar.

test.describe('Sidebar: EN index has only /categories/ links (no /ja/categories/)', () => {
  test('aside has no /ja/categories/ hrefs on EN index', async ({ page }) => {
    await page.goto('/');
    const jaLinks = page.locator('aside a[href*="/ja/categories/"]');
    expect(await jaLinks.count()).toBe(0);
  });

  test('aside has no /ja/tags/ hrefs on EN index', async ({ page }) => {
    await page.goto('/');
    const jaLinks = page.locator('aside a[href*="/ja/tags/"]');
    expect(await jaLinks.count()).toBe(0);
  });

  test('aside has at least one /categories/ link on EN index', async ({ page }) => {
    await page.goto('/');
    const links = page.locator('aside a[href^="/categories/"]');
    expect(await links.count()).toBeGreaterThan(0);
  });
});

test.describe('Sidebar: JA index has only /ja/categories/ links (no EN /categories/)', () => {
  test('aside has no bare /categories/ hrefs on JA index', async ({ page }) => {
    await page.goto('/ja/');
    // A link starting with /categories/ (but not /ja/categories/) would be a locale leak.
    const enLinks = page.locator('aside a[href^="/categories/"]');
    expect(await enLinks.count()).toBe(0);
  });

  test('aside has no bare /tags/ hrefs on JA index', async ({ page }) => {
    await page.goto('/ja/');
    const enLinks = page.locator('aside a[href^="/tags/"]');
    expect(await enLinks.count()).toBe(0);
  });

  test('aside has at least one /ja/categories/ link on JA index', async ({ page }) => {
    await page.goto('/ja/');
    const links = page.locator('aside a[href^="/ja/categories/"]');
    expect(await links.count()).toBeGreaterThan(0);
  });

  test('aside has at least one /ja/tags/ link on JA index', async ({ page }) => {
    await page.goto('/ja/');
    const links = page.locator('aside a[href^="/ja/tags/"]');
    expect(await links.count()).toBeGreaterThan(0);
  });
});
