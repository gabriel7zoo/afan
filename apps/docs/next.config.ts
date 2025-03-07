module.exports = {
  basePath: "/docs", // Forces all pages to be inside /docs
  async redirects() {
    return [
      {
        source: "/",
        destination: "/docs/quickstart/onboarding/onboarding-api", // Change this if needed
        permanent: true,
      },
    ];
  },
};
