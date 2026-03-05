import { test, expect } from '@playwright/test';

// ── OGP / meta tags ───────────────────────────────────────────────────────────
// Note: og:title / og:url / og:description are NOT yet implemented in the
// theme templates. These tests cover what IS currently generated.
//
// TODO: implement OGP meta tags in _partials.html seo block, then add:
//   test('og:title is present', ...)
//   test('og:url is present', ...)
//   test('og:description is present', ...)

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

// ── canonical / hreflang on article pages ────────────────────────────────────
// The article.html template defines a "seo" block with canonical + hreflang.
// These links should be present on individual article pages.

test.describe('canonical link on EN article /posts/hello-world/', () => {
  const URL = '/posts/hello-world/';

  test('canonical link is present and contains the article URL', async ({ page }) => {
    await page.goto(URL);
    const canonical = page.locator('link[rel="canonical"]');
    await expect(canonical).toHaveCount(1);
    const href = await canonical.getAttribute('href');
    expect(href).toContain('/posts/hello-world/');
    expect(href).toContain('bmf-tech.com');
  });

  test('hreflang="en" link is present', async ({ page }) => {
    await page.goto(URL);
    const hreflang = page.locator('link[rel="alternate"][hreflang="en"]');
    expect(await hreflang.count()).toBeGreaterThanOrEqual(1);
  });

  test('hreflang="x-default" link is present', async ({ page }) => {
    await page.goto(URL);
    const hreflang = page.locator('link[rel="alternate"][hreflang="x-default"]');
    expect(await hreflang.count()).toBeGreaterThanOrEqual(1);
  });
});

test.describe('canonical link on JA article', () => {
  const URL = '/ja/posts/2018-review-2019-goals/';

  test('canonical link is present and contains the article URL', async ({ page }) => {
    await page.goto(URL);
    const canonical = page.locator('link[rel="canonical"]');
    expect(await canonical.count()).toBeGreaterThanOrEqual(1);
    const href = await canonical.first().getAttribute('href');
    expect(href).toContain('/ja/posts/2018-review-2019-goals/');
  });

  test('hreflang="ja" link is present', async ({ page }) => {
    await page.goto(URL);
    const hreflang = page.locator('link[rel="alternate"][hreflang="ja"]');
    expect(await hreflang.count()).toBeGreaterThanOrEqual(1);
  });
});

