const { withContentCollections } = require("@content-collections/next");

const securityHeaders = [
  {
    key: "X-Frame-Options",
    value: "SAMEORIGIN",
  },
];

const withBundleAnalyzer = require("@next/bundle-analyzer")({
  enabled: process.env.ANALYZE === "true",
});

/** @type {import('next').NextConfig} */
const nextConfig = {
  pageExtensions: ["tsx", "mdx", "ts", "js"],
  reactStrictMode: true,
  swcMinify: true,
  async headers() {
    return [
      {
        source: "/(.*)",
        headers: securityHeaders,
      },
    ];
  },
  async rewrites() {
    return [
      {
        source: "/changelog/:slug",
        destination: "/changelog#:slug",
      },
      {
        source: "/terms",
        destination: "/policies/terms",
      },
    ];
  },
  async redirects() {
    return [
      {
        source: "/discord",
        destination: "https://discord.gg/fDbezjbJbD",
        permanent: false,
      },
      {
        source: "/github",
        destination: "https://github.com/unkeyed/unkey",
        permanent: false,
      },
      {
        source: "/meet",
        destination: "https://cal.com/team/unkey",
        permanent: false,
      },
    ];
  },
};

module.exports = withContentCollections(withBundleAnalyzer(nextConfig));
