module.exports = {
  ci: {
    collect: {
      staticDistDir: "./public",
      // Sample: top page + one article + one taxonomy page + ja version.
      url: [
        "http://localhost/index.html",
        "http://localhost/ja/index.html",
      ],
      numberOfRuns: 1,
    },
    assert: {
      // Baseline thresholds (intentionally lenient to start).
      // Tighten over time as site improves.
      assertions: {
        "categories:performance": ["warn", { minScore: 0.8 }],
        "categories:accessibility": ["warn", { minScore: 0.9 }],
        "categories:best-practices": ["warn", { minScore: 0.9 }],
        "categories:seo": ["warn", { minScore: 0.9 }],
      },
    },
    upload: {
      target: "temporary-public-storage",
    },
  },
};
