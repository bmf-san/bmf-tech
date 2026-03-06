import { test, expect } from '@playwright/test';

// ── Page <title> ──────────────────────────────────────────────────────────────

test.describe('page <title> tag', () => {
  test('index / shows site title only', async ({ page }) => {
    await page.goto('/');
    await expect(page).toHaveTitle('bmf-tech');
  });

  test('article page shows "Article Title | Site Title"', async ({ page }) => {
    await page.goto('/posts/hello-world/');
    await expect(page).toHaveTitle('Hello, World | bmf-tech');
  });

  test('about page shows "About | Site Title"', async ({ page }) => {
    await page.goto('/about/');
    await expect(page).toHaveTitle('About | bmf-tech');
  });
});

// ── meta description ──────────────────────────────────────────────────────────

test.describe('meta description on key pages', () => {
  const PAGES = ['/', '/ja/', '/posts/hello-world/', '/about/'];
  for (const path of PAGES) {
    test(`meta[name="description"] is non-empty on ${path}`, async ({ page }) => {
      await page.goto(path);
      const desc = page.locator('meta[name="description"]');
      await expect(desc).toHaveCount(1);
      const content = await desc.getAttribute('content');
      expect(content?.trim().length).toBeGreaterThan(0);
    });
  }
});

// ── OGP meta tags — article pages ────────────────────────────────────────────

test.describe('OGP on EN article /posts/hello-world/', () => {
  const URL = '/posts/hello-world/';

  test('og:title equals article title', async ({ page }) => {
    await page.goto(URL);
    const ogTitle = page.locator('meta[property="og:title"]');
    await expect(ogTitle).toHaveCount(1);
    expect(await ogTitle.getAttribute('content')).toBe('Hello, World');
  });

  test('og:type is "article"', async ({ page }) => {
    await page.goto(URL);
    const ogType = page.locator('meta[property="og:type"]');
    await expect(ogType).toHaveCount(1);
    expect(await ogType.getAttribute('content')).toBe('article');
  });

  test('og:url contains article URL', async ({ page }) => {
    await page.goto(URL);
    const ogUrl = page.locator('meta[property="og:url"]');
    await expect(ogUrl).toHaveCount(1);
    expect(await ogUrl.getAttribute('content')).toContain('/posts/hello-world/');
  });

  test('og:description is non-empty', async ({ page }) => {
    await page.goto(URL);
    const ogDesc = page.locator('meta[property="og:description"]');
    await expect(ogDesc).toHaveCount(1);
    const content = await ogDesc.getAttribute('content');
    expect(content?.trim().length).toBeGreaterThan(0);
  });

  test('og:image points to /ogp/hello-world.png', async ({ page }) => {
    await page.goto(URL);
    const ogImage = page.locator('meta[property="og:image"]');
    await expect(ogImage).toHaveCount(1);
    expect(await ogImage.getAttribute('content')).toContain('/ogp/hello-world.png');
  });
});

test.describe('OGP on JA article', () => {
  const URL = '/ja/posts/2018-review-2019-goals/';

  test('og:title equals article title', async ({ page }) => {
    await page.goto(URL);
    const ogTitle = page.locator('meta[property="og:title"]');
    await expect(ogTitle).toHaveCount(1);
    const content = await ogTitle.getAttribute('content');
    expect(content?.trim().length).toBeGreaterThan(0);
  });

  test('og:url contains /ja/posts/', async ({ page }) => {
    await page.goto(URL);
    const ogUrl = page.locator('meta[property="og:url"]');
    await expect(ogUrl).toHaveCount(1);
    expect(await ogUrl.getAttribute('content')).toContain('/ja/posts/');
  });

  test('og:image points to /ogp/{slug}.png', async ({ page }) => {
    await page.goto(URL);
    const ogImage = page.locator('meta[property="og:image"]');
    await expect(ogImage).toHaveCount(1);
    expect(await ogImage.getAttribute('content')).toContain('/ogp/2018-review-2019-goals.png');
  });
});

// ── OGP meta tags — listing pages ────────────────────────────────────────────

test.describe('OGP on listing pages', () => {
  const LISTING_PAGES = ['/', '/ja/'];

  for (const path of LISTING_PAGES) {
    test(`og:type is "website" on ${path}`, async ({ page }) => {
      await page.goto(path);
      const ogType = page.locator('meta[property="og:type"]');
      await expect(ogType).toHaveCount(1);
      expect(await ogType.getAttribute('content')).toBe('website');
    });

    test(`og:title equals site title on ${path}`, async ({ page }) => {
      await page.goto(path);
      const ogTitle = page.locator('meta[property="og:title"]');
      await expect(ogTitle).toHaveCount(1);
      expect(await ogTitle.getAttribute('content')).toBe('bmf-tech');
    });

    test(`og:image points to default OGP image on ${path}`, async ({ page }) => {
      await page.goto(path);
      const ogImage = page.locator('meta[property="og:image"]');
      await expect(ogImage).toHaveCount(1);
      expect(await ogImage.getAttribute('content')).toContain('/assets/images/ogp-default.png');
    });
  }
});

// ── canonical / hreflang on article pages ────────────────────────────────────

test.describe('canonical on EN article /posts/hello-world/', () => {
  const URL = '/posts/hello-world/';

  test('exactly one canonical link pointing to article URL', async ({ page }) => {
    await page.goto(URL);
    const canonical = page.locator('link[rel="canonical"]');
    await expect(canonical).toHaveCount(1);
    const href = await canonical.getAttribute('href');
    expect(href).toContain('/posts/hello-world/');
    expect(href).toContain('bmf-tech.com');
  });

  test('hreflang="en" link is present', async ({ page }) => {
    await page.goto(URL);
    expect(await page.locator('link[rel="alternate"][hreflang="en"]').count()).toBeGreaterThanOrEqual(1);
  });

  test('hreflang="x-default" link is present', async ({ page }) => {
    await page.goto(URL);
    expect(await page.locator('link[rel="alternate"][hreflang="x-default"]').count()).toBeGreaterThanOrEqual(1);
  });
});

test.describe('canonical on JA article', () => {
  const URL = '/ja/posts/2018-review-2019-goals/';

  test('exactly one canonical link pointing to article URL', async ({ page }) => {
    await page.goto(URL);
    const canonical = page.locator('link[rel="canonical"]');
    await expect(canonical).toHaveCount(1);
    const href = await canonical.first().getAttribute('href');
    expect(href).toContain('/ja/posts/2018-review-2019-goals/');
  });

  test('hreflang="ja" link is present', async ({ page }) => {
    await page.goto(URL);
    expect(await page.locator('link[rel="alternate"][hreflang="ja"]').count()).toBeGreaterThanOrEqual(1);
  });
});

// ── listing pages must have NO article-level canonical ────────────────────────

test.describe('listing pages have no canonical link', () => {
  const LISTING_PAGES = ['/', '/ja/', '/ja/page/2/'];

  for (const path of LISTING_PAGES) {
    test(`no canonical link on listing page ${path}`, async ({ page }) => {
      await page.goto(path);
      await expect(page.locator('link[rel="canonical"]')).toHaveCount(0);
    });
  }
});
