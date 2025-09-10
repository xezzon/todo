/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
  rewrites() {
    return [
      {
        source: '/api/TodoService/:path*',
        destination: `http://localhost:8080/TodoService/:path*`
      }
    ]
  },
};

export default nextConfig;
