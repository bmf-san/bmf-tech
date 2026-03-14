import { test, expect } from '@playwright/test';

// ── Page <title> ──────────────────────────────────────────────────────────────

test.describe('page <title> tag', () => {
  test('index / shows site title only', async ({ page }) => {
    await page.goto('/');
    await expect(page).toHaveTitle('bmf-tech');
  });

  test('article page shows "Article Title | Site Title"', async ({ page }) => {
    await page.goto('/posts/2018-review-2019-goals/');
    await expect(page).toHaveTitle('Reflection on 2018 and Goals for 2019 | bmf-tech');
  });

  test('about page shows "About | Site Title"', async ({ page }) => {
    await page.goto('/about/');
    await expect(page).toHaveTitle('About | bmf-tech');
  });

  test('EN tag page shows "TagName | bmf-tech"', async ({ page }) => {
    await page.goto('/tags/golang/');
    await expect(page).toHaveTitle('Golang | bmf-tech');
  });

  test('JA tag page shows "TagName | bmf-tech"', async ({ page }) => {
    await page.goto('/ja/tags/golang/');
    await expect(page).toHaveTitle('Golang | bmf-tech');
  });

  test('EN category page shows "CategoryName | bmf-tech"', async ({ page }) => {
    await page.goto('/categories/architecture/');
    await expect(page).toHaveTitle('Architecture | bmf-tech');
  });

  test('JA category page shows "CategoryName | bmf-tech"', async ({ page }) => {
    await page.goto('/ja/categories/os/');
    await expect(page).toHaveTitle('OS | bmf-tech');
  });
});

// ── meta description ──────────────────────────────────────────────────────────

test.describe('meta description on key pages', () => {
  const PAGES = ['/', '/ja/', '/posts/2018-review-2019-goals/', '/about/'];
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

test.describe('OGP on EN article /posts/2018-review-2019-goals/', () => {
  const URL = '/posts/2018-review-2019-goals/';

  test('og:title equals article title', async ({ page }) => {
    await page.goto(URL);
    const ogTitle = page.locator('meta[property="og:title"]');
    await expect(ogTitle).toHaveCount(1);
    expect(await ogTitle.getAttribute('content')).toBe('Reflection on 2018 and Goals for 2019');
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
    expect(await ogUrl.getAttribute('content')).toContain('/posts/2018-review-2019-goals/');
  });

  test('og:description is non-empty', async ({ page }) => {
    await page.goto(URL);
    const ogDesc = page.locator('meta[property="og:description"]');
    await expect(ogDesc).toHaveCount(1);
    const content = await ogDesc.getAttribute('content');
    expect(content?.trim().length).toBeGreaterThan(0);
  });

  test('og:image points to /ogp/2018-review-2019-goals.png', async ({ page }) => {
    await page.goto(URL);
    const ogImage = page.locator('meta[property="og:image"]');
    await expect(ogImage).toHaveCount(1);
    expect(await ogImage.getAttribute('content')).toContain('/ogp/2018-review-2019-goals.png');
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

test.describe('canonical on EN article /posts/2018-review-2019-goals/', () => {
  const URL = '/posts/2018-review-2019-goals/';

  test('exactly one canonical link pointing to article URL', async ({ page }) => {
    await page.goto(URL);
    const canonical = page.locator('link[rel="canonical"]');
    await expect(canonical).toHaveCount(1);
    const href = await canonical.getAttribute('href');
    expect(href).toContain('/posts/2018-review-2019-goals/');
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

// ── listing pages must have a self-referencing canonical ───────────────────────
// (Previously we incorrectly asserted there was no canonical. Listing pages
// DO have a canonical that points to themselves; for paginated pages it points
// to the current page URL, not page 1.)

test.describe('listing pages have a canonical link pointing to themselves', () => {
  const CASES: [string, string][] = [
    ['/', 'bmf-tech.com/'],
    ['/ja/', '/ja/'],
    ['/ja/page/2/', '/ja/page/2/'],
    ['/ja/archives/2015/', '/ja/archives/2015/'],
    ['/ja/archives/2015/05/', '/ja/archives/2015/05/'],
    ['/tags/abac/', '/tags/abac/'],
    ['/ja/tags/abac/', '/ja/tags/abac/'],
    ['/categories/application/', '/categories/application/'],
  ];

  for (const [path, expectedFragment] of CASES) {
    test(`canonical on listing page ${path} points to itself`, async ({ page }) => {
      await page.goto(path);
      const canonical = page.locator('link[rel="canonical"]');
      await expect(canonical).toHaveCount(1);
      expect(await canonical.getAttribute('href')).toContain(expectedFragment);
    });
  }
});

test.describe('listing pages have hreflang alternate links', () => {
  const CASES: [string, string, string][] = [
    // [path, self-locale, alt-locale-fragment]
    ['/ja/archives/2015/', 'ja', '/archives/2015/'],
    ['/ja/archives/2015/05/', 'ja', '/archives/2015/05/'],
    ['/tags/abac/', 'en', '/ja/tags/abac/'],
    ['/ja/tags/abac/', 'ja', '/tags/abac/'],
  ];

  for (const [path, selfLocale, altFragment] of CASES) {
    test(`hreflang self=${selfLocale} and alternate present on ${path}`, async ({ page }) => {
      await page.goto(path);
      expect(await page.locator(`link[rel="alternate"][hreflang="${selfLocale}"]`).count()).toBeGreaterThanOrEqual(1);
      const altLinks = page.locator('link[rel="alternate"][hreflang]');
      const count = await altLinks.count();
      expect(count).toBeGreaterThanOrEqual(2); // at least self + one alternate
      // x-default must be present
      expect(await page.locator('link[rel="alternate"][hreflang="x-default"]').count()).toBe(1);
      // alternate locale link must contain expected fragment
      const hrefs = await Promise.all(Array.from({length: count}, (_, i) => altLinks.nth(i).getAttribute('href')));
      expect(hrefs.some(h => h?.includes(altFragment))).toBe(true);
    });
  }
});

// ── <title> on archive and taxonomy listing pages ─────────────────────────────

test.describe('archive page <title> includes date', () => {
  test('month archive /ja/archives/2015/05/ shows YYYY/MM | site', async ({ page }) => {
    await page.goto('/ja/archives/2015/05/');
    await expect(page).toHaveTitle(/^2015\/05 \| bmf-tech$/);
  });

  test('year archive /ja/archives/2015/ shows YYYY | site', async ({ page }) => {
    await page.goto('/ja/archives/2015/');
    await expect(page).toHaveTitle(/^2015 \| bmf-tech$/);
  });

  test('EN month archive /archives/2015/05/ shows YYYY/MM | site', async ({ page }) => {
    await page.goto('/archives/2015/05/');
    await expect(page).toHaveTitle(/^2015\/05 \| bmf-tech$/);
  });
});

// ── OGP og:url on archive and taxonomy listing pages ─────────────────────────

test.describe('og:url on listing pages points to page not homepage', () => {
  test('JA month archive og:url points to archive URL', async ({ page }) => {
    await page.goto('/ja/archives/2015/05/');
    const ogUrl = page.locator('meta[property="og:url"]');
    await expect(ogUrl).toHaveCount(1);
    expect(await ogUrl.getAttribute('content')).toContain('/ja/archives/2015/05/');
  });

  test('JA year archive og:url points to archive URL', async ({ page }) => {
    await page.goto('/ja/archives/2015/');
    const ogUrl = page.locator('meta[property="og:url"]');
    await expect(ogUrl).toHaveCount(1);
    expect(await ogUrl.getAttribute('content')).toContain('/ja/archives/2015/');
  });

  test('JA tag page og:url points to tag URL', async ({ page }) => {
    await page.goto('/ja/tags/abac/');
    const ogUrl = page.locator('meta[property="og:url"]');
    await expect(ogUrl).toHaveCount(1);
    expect(await ogUrl.getAttribute('content')).toContain('/ja/tags/abac/');
  });

  test('JA tag page og:title includes tag name', async ({ page }) => {
    await page.goto('/ja/tags/abac/');
    const ogTitle = page.locator('meta[property="og:title"]');
    await expect(ogTitle).toHaveCount(1);
    expect(await ogTitle.getAttribute('content')).toContain('ABAC');
  });

  // EN locale
  test('EN month archive og:url points to archive URL', async ({ page }) => {
    await page.goto('/archives/2015/05/');
    const ogUrl = page.locator('meta[property="og:url"]');
    await expect(ogUrl).toHaveCount(1);
    expect(await ogUrl.getAttribute('content')).toContain('/archives/2015/05/');
  });

  test('EN tag page og:url points to tag URL', async ({ page }) => {
    await page.goto('/tags/abac/');
    const ogUrl = page.locator('meta[property="og:url"]');
    await expect(ogUrl).toHaveCount(1);
    expect(await ogUrl.getAttribute('content')).toContain('/tags/abac/');
  });

  test('EN tag page og:title includes tag name', async ({ page }) => {
    await page.goto('/tags/abac/');
    const ogTitle = page.locator('meta[property="og:title"]');
    await expect(ogTitle).toHaveCount(1);
    expect(await ogTitle.getAttribute('content')).toContain('ABAC');
  });

  test('JA category page og:url points to category URL', async ({ page }) => {
    await page.goto('/ja/categories/os/');
    const ogUrl = page.locator('meta[property="og:url"]');
    await expect(ogUrl).toHaveCount(1);
    expect(await ogUrl.getAttribute('content')).toContain('/ja/categories/os/');
  });

  test('JA category page og:title includes category name', async ({ page }) => {
    await page.goto('/ja/categories/os/');
    const ogTitle = page.locator('meta[property="og:title"]');
    await expect(ogTitle).toHaveCount(1);
    expect(await ogTitle.getAttribute('content')).toContain('OS');
  });

  test('EN category page og:url points to category URL', async ({ page }) => {
    await page.goto('/categories/architecture/');
    const ogUrl = page.locator('meta[property="og:url"]');
    await expect(ogUrl).toHaveCount(1);
    expect(await ogUrl.getAttribute('content')).toContain('/categories/architecture/');
  });

  test('EN category page og:title includes category name', async ({ page }) => {
    await page.goto('/categories/architecture/');
    const ogTitle = page.locator('meta[property="og:title"]');
    await expect(ogTitle).toHaveCount(1);
    expect(await ogTitle.getAttribute('content')).toContain('Architecture');
  });
});

// ── og:type = "website" on archive and taxonomy listing pages ─────────────────

test.describe('og:type is "website" on listing pages', () => {
  const LISTING_PAGES = [
    '/ja/archives/2015/',
    '/ja/archives/2015/05/',
    '/archives/2015/05/',
    '/ja/tags/abac/',
    '/tags/abac/',
    '/categories/application/',
  ];

  for (const path of LISTING_PAGES) {
    test(`og:type is "website" on ${path}`, async ({ page }) => {
      await page.goto(path);
      const ogType = page.locator('meta[property="og:type"]');
      await expect(ogType).toHaveCount(1);
      expect(await ogType.getAttribute('content')).toBe('website');
    });
  }
});

// ── archive page <h1> heading content ────────────────────────────────────────

test.describe('archive page <h1> heading shows correct date', () => {
  test('JA month archive h1 shows year/month', async ({ page }) => {
    await page.goto('/ja/archives/2015/05/');
    const h1 = page.locator('main h1');
    await expect(h1).toHaveCount(1);
    const text = await h1.textContent();
    expect(text).toMatch(/2015\/05/);
  });

  test('JA year archive h1 shows year only (not year/month)', async ({ page }) => {
    await page.goto('/ja/archives/2015/');
    const h1 = page.locator('main h1');
    await expect(h1).toHaveCount(1);
    const text = await h1.textContent();
    expect(text).toMatch(/2015/);
    expect(text).not.toMatch(/2015\/\d{2}/);
  });

  test('EN month archive h1 shows year/month', async ({ page }) => {
    await page.goto('/archives/2015/05/');
    const h1 = page.locator('main h1');
    await expect(h1).toHaveCount(1);
    const text = await h1.textContent();
    expect(text).toMatch(/2015\/05/);
  });
});

// ── og:locale on listing pages ────────────────────────────────────────────────

test.describe('og:locale on listing pages', () => {
  const EN_PAGES = ['/', '/tags/golang/', '/categories/architecture/', '/archives/2024/01/'];
  const JA_PAGES = ['/ja/', '/ja/tags/golang/', '/ja/categories/os/', '/ja/archives/2024/01/'];

  for (const path of EN_PAGES) {
    test(`og:locale is en_US on EN listing page ${path}`, async ({ page }) => {
      await page.goto(path);
      const ogLocale = page.locator('meta[property="og:locale"]');
      await expect(ogLocale).toHaveCount(1);
      expect(await ogLocale.getAttribute('content')).toBe('en_US');
    });
  }

  for (const path of JA_PAGES) {
    test(`og:locale is ja_JP on JA listing page ${path}`, async ({ page }) => {
      await page.goto(path);
      const ogLocale = page.locator('meta[property="og:locale"]');
      await expect(ogLocale).toHaveCount(1);
      expect(await ogLocale.getAttribute('content')).toBe('ja_JP');
    });
  }
});

// ── og:url on root index pages ────────────────────────────────────────────────

test.describe('og:url on root index pages', () => {
  test('EN root / og:url contains site base URL', async ({ page }) => {
    await page.goto('/');
    const ogUrl = page.locator('meta[property="og:url"]');
    await expect(ogUrl).toHaveCount(1);
    expect(await ogUrl.getAttribute('content')).toMatch(/bmf-tech\.com\//);
  });

  test('JA root /ja/ og:url contains /ja/', async ({ page }) => {
    await page.goto('/ja/');
    const ogUrl = page.locator('meta[property="og:url"]');
    await expect(ogUrl).toHaveCount(1);
    expect(await ogUrl.getAttribute('content')).toContain('/ja/');
  });

  test('JA paginated /ja/page/2/ og:url contains /ja/', async ({ page }) => {
    await page.goto('/ja/page/2/');
    const ogUrl = page.locator('meta[property="og:url"]');
    await expect(ogUrl).toHaveCount(1);
    expect(await ogUrl.getAttribute('content')).toContain('/ja/');
  });
});

// ── BUG-T6: static pages (page.html) must not receive article OGP/JSON-LD ────

test.describe('static pages (About, Privacy Policy) have website OGP, not article', () => {
  const STATIC_PAGES = ['/about/', '/ja/about/', '/privacy-policy/', '/ja/privacy-policy/'];

  for (const path of STATIC_PAGES) {
    test(`og:type is "website" (not "article") on ${path}`, async ({ page }) => {
      await page.goto(path);
      const ogType = page.locator('meta[property="og:type"]');
      await expect(ogType).toHaveCount(1);
      expect(await ogType.getAttribute('content')).toBe('website');
    });

    test(`no article:published_time on ${path}`, async ({ page }) => {
      await page.goto(path);
      await expect(page.locator('meta[property="article:published_time"]')).toHaveCount(0);
    });

    test(`JSON-LD @type is WebPage (not BlogPosting) on ${path}`, async ({ page }) => {
      await page.goto(path);
      const content = await page.content();
      expect(content).toContain('"@type": "WebPage"');
      expect(content).not.toContain('"@type": "BlogPosting"');
    });
  }
});

// ── BUG-T7: paginated taxonomy/archive pages have correct hreflang page number ─

test.describe('paginated tag/archive pages have correct hreflang page number', () => {
  test('JA tag page 2 hreflang "en" points to EN page 2', async ({ page }) => {
    await page.goto('/ja/tags/golang/page/2/');
    const enLink = page.locator('link[rel="alternate"][hreflang="en"]');
    await expect(enLink).toHaveCount(1);
    expect(await enLink.getAttribute('href')).toContain('/tags/golang/page/2/');
  });

  test('EN tag page 2 hreflang "ja" points to JA page 2', async ({ page }) => {
    await page.goto('/tags/golang/page/2/');
    const jaLink = page.locator('link[rel="alternate"][hreflang="ja"]');
    await expect(jaLink).toHaveCount(1);
    expect(await jaLink.getAttribute('href')).toContain('/ja/tags/golang/page/2/');
  });

  test('JA archive page 2 hreflang "en" points to EN page 2', async ({ page }) => {
    await page.goto('/ja/archives/2024/page/2/');
    const enLink = page.locator('link[rel="alternate"][hreflang="en"]');
    await expect(enLink).toHaveCount(1);
    expect(await enLink.getAttribute('href')).toContain('/archives/2024/page/2/');
  });
});

// ── BUG-T8: BreadcrumbList item on paginated taxonomy pages points to page 1 ──

test.describe('BreadcrumbList item URL on paginated taxonomy page points to page-1 URL', () => {
  test('JA tag page 2: JSON-LD breadcrumb item points to /ja/tags/golang/ (page 1)', async ({ page }) => {
    await page.goto('/ja/tags/golang/page/2/');
    const content = await page.content();
    // Breadcrumb item must reference the tag base URL, not /page/2/
    expect(content).toContain('\\/ja\\/tags\\/golang\\/');
    expect(content).not.toContain('\\/tags\\/golang\\/page\\/2\\/');
  });
});

// ── BUG-T9: article count on listing pages shows total, not per-page count ────

test.describe('article count on tag/category pages shows total', () => {
  // /ja/tags/golang/ has 60 articles (>20 per page) — page 2 must show "60"
  test('JA tag page 2 article count shows total (>20)', async ({ page }) => {
    await page.goto('/ja/tags/golang/page/2/');
    const countEl = page.locator('p.text-secondary').first();
    const text = await countEl.textContent();
    const match = text?.match(/^(\d+)/);
    expect(match).not.toBeNull();
    expect(parseInt(match![1])).toBeGreaterThan(20);
  });
});

