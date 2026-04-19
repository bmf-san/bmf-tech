import { test, expect } from '@playwright/test';

// ── CSS loading ───────────────────────────────────────────────────────────────
// Verify sleyt CSS link is present in <head> across every major page type.

const PAGES_WITH_CSS = [
  '/',
  '/ja/',
  '/posts/2018-review-2019-goals/',
  '/ja/posts/2018-review-2019-goals/',
  '/tags/golang/',
  '/ja/tags/golang/',
  '/categories/tools/',
  '/ja/categories/ツール/',
  '/about/',
  '/archives/2024/03/',
  '/ja/archives/2026/02/',
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
  await expect(page.locator('link[rel="alternate"][type="application/atom+xml"]')).toHaveCount(1);
});

// ── Sitemap ───────────────────────────────────────────────────────────────────

test('sitemap.xml is served and contains URLs', async ({ request }) => {
  const res = await request.get('/sitemap.xml');
  expect(res.status()).toBe(200);
  const body = await res.text();
  expect(body).toContain('<urlset');
  expect(body).toContain('bmf-tech.com');
});

test('sitemap.xml contains EN and JA homepage URLs', async ({ request }) => {
  const res = await request.get('/sitemap.xml');
  const body = await res.text();
  expect(body).toContain('https://bmf-tech.com/</loc>');
  expect(body).toContain('https://bmf-tech.com/ja/</loc>');
});

test('sitemap.xml contains a known article URL', async ({ request }) => {
  const res = await request.get('/sitemap.xml');
  const body = await res.text();
  expect(body).toContain('/posts/2018-review-2019-goals/</loc>');
  expect(body).toContain('/ja/posts/2018-review-2019-goals/</loc>');
});

test('sitemap.xml article entry has xhtml:link hreflang alternates', async ({ request }) => {
  const res = await request.get('/sitemap.xml');
  const body = await res.text();
  // Each bilingual article should have en and ja xhtml:link alternate entries
  expect(body).toContain('hreflang="en"');
  expect(body).toContain('hreflang="ja"');
  expect(body).toContain('hreflang="x-default"');
});

test('sitemap.xml contains lastmod date on article entries', async ({ request }) => {
  const res = await request.get('/sitemap.xml');
  const body = await res.text();
  expect(body).toMatch(/<lastmod>\d{4}-\d{2}-\d{2}<\/lastmod>/);
});

// ── OGP images ────────────────────────────────────────────────────────────────

const OGP_SLUGS = [
  '2018-review-2019-goals',
  'cto-thinking-strategy-leadership',
  'engineering-in-ai-reflections',
];

for (const slug of OGP_SLUGS) {
  test(`OGP image /ogp/${slug}.png returns HTTP 200`, async ({ request }) => {
    const res = await request.get(`/ogp/${slug}.png`);
    expect(res.status()).toBe(200);
    expect(res.headers()['content-type']).toMatch(/image\/png/);
  });
}
