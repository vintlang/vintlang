import '@/app/globals.css';
import { Inter } from 'next/font/google';
import { ThemeProvider } from '@/components/theme-provider';

const inter = Inter({ subsets: ['latin'] });

export const metadata = {
  title: 'VintLang - Modern Programming Made Simple',
  description:
    'A powerful programming language built with Go, featuring intuitive syntax, built-in networking, and comprehensive time operations.',
  keywords: 'VintLang, Go Programming Language, Modern Programming, Networking, Time Operations',
  author: 'Tachera Sasi',
  viewport: 'width=device-width, initial-scale=1.0',
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en">
      <head>
        <title>{metadata.title}</title>
        <meta name="description" content={metadata.description} />
        <meta name="keywords" content={metadata.keywords} />
        <meta name="author" content={metadata.author} />
        <meta name="viewport" content={metadata.viewport} />
      </head>
      <body className={`${inter.className} max-w-screen-xl mx-auto`}>
        <ThemeProvider attribute="class" defaultTheme="system" enableSystem>
          {children}
        </ThemeProvider>
      </body>
    </html>
  );
}
