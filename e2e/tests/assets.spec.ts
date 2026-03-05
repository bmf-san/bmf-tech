import { test, expect } from '@playwright/test';

// ── CSS loading ───────────────────────────────────────────────────────────────
// Regression: without content/ja/tags.md, gohan served a directory listing
// (text/plain, no stylesheet). Verify sleyt CSS link is present in <head>
// across every major page type.

const PAGES_WITH_CSS = [
  '/',
  '/ja/',
  '/posts/hello-world/',
  '/ja/posts/2018-review-2019-goals/',
  '/tags/',
  '/ja/tags/',
  '/categories/',
  '/ja/categories/',
  '/about/',
  '/archives/2024/03/',
];

for (const path of PAGES_WITH_CSS) {
  test(`CSS stylesheet is loaded on ${path}`, async ({ page }) => {
    await page.goto(path);
    // Sleyt CSS CDN link must be present in <head>
    const cssLink = page.locator('link[rel="stylesheet"][href*="sleyt"]');
    await expect(cssLink).toHaveCount(1);
  });
}

// ── Favicon ───────────────────────────────────────────────────────────────────

test('favicon.ico is served (200)', async ({ request }) => {
  const res = await request.get('/favicon.ico');
  expect(res.status()).toBe(200);
  expect(res.headers()['content-type']).toMatch(/image/);
});

test('favicon link element is present in <head>', async ({ page }) => {
  await page.goto('/');
  await expect(page.locator('link[rel="icon"][href="/favicon.ico"]')).toHaveCount(1);
});

// ── Feed / Atom ───────────────────────────────────────────────────────────────

test('atom.xml is served with XML content-type', async ({ request }) => {
  const res = await request.get('/atom.xml');
  expect(res.status()).toBe(200);
  expect(res.headers()['content-type']).toMatch(/xml/);
  const body = await res.text();
  expect(body).toContain('<feed');
  expect(body).toContain('bmf-tech');
});

test('feed.xml is served with XML content-type', async ({ request }) => {
  const res = await request.get('/feed.xml');
  expect(res.status()).toBe(200);
  expect(res.headers()['content-type']).toMatch(/xml/);
  const body = await res.text();
  expect(body).toContain('bmf-tech');
});

test('feed link element is present in <head>', async ({ page }) => {
  await page.goto('/');
  await expect(
    page.locator('link[rel="alternate"][type="application/atom+xml"]')
  ).toHaveCount(1);
});

// ── Sitemap ───────────────────────────────────────────────────────────────────

test('sitemap.xml is served and contains URLs', async ({ request }) => {
  const res = await request.get('/sitemap.xml');
  expect(res.status()).toBe(200);
  const body = await res.text();
  expect(body).toContain('<urlset');
  expect(body).toContain('bmf-tech.com');
});
