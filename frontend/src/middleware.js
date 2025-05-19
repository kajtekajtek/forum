import { NextResponse } from 'next/server';

export function middleware(req) {
    const token = req.cookies.get('token')?.value;
    const { pathname } = req.nextUrl;

    if (!token && pathname !== '/login' && pathname !== '/register') {
        const url = req.nextUrl.clone();
        url.pathname = '/login';
        return NextResponse.redirect(url);
    }
    
    return NextResponse.next();
}

export const config = {
    matcher: ['/((?!api|_next/static|_next/image|favicon.ico).*)'],
}
