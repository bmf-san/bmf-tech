import { defineConfig, devices } from '@playwright/test';

export default defineConfig({
  testDir: './tests',
  fullyParallel: true,
  forbidOnly: !!process.env.CI,
  retries: process.env.CI ? 2 : 0,
  workers: 1,
  reporter: process.env.CI
    ? [['github'], ['junit', { outputFile: 'results.xml' }]]
    : 'html',
  use: {
    baseURL: 'http://localhost:1313',
    trace: 'on-first-retry',
  },
  projects: [
    {
      name: 'chromium',
      use: { ...devices['Desktop Chrome'] },
    },
  ],
  webServer: {
    command: 'npx http-server ../public -p 1313 -s --cors',
    url: 'http://localhost:1313',
    reuseExistingServer: !process.env.CI,
    timeout: 10 * 1000,
  },
});
