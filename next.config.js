module.exports = {
  // Set `apps/docs` as the main Next.js app
  basePath: "/",
  distDir: "apps/docs/.next",

  // Redirect all requests to `docs`
  async redirects() {
    return [
      {
        source: "/",
        destination: "/docs",
        permanent: true,
      },
    ];
  },
};
