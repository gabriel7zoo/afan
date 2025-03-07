import { NextRequest, NextResponse } from 'next/server';

export function middleware(req: NextRequest) {
  const url = req.nextUrl;

  // Allow requests only under /docs/
  if (!url.pathname.startsWith('/docs')) {
    return new NextResponse(null, { status: 403 }); // Forbidden
  }

  return NextResponse.next();
}

// Apply middleware to all routes
export const config = {
  matcher: '/:path*', // Matches all routes
};
