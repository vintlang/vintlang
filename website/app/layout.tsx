import "@/app/globals.css";
import { Analytics } from "@vercel/analytics/next";
import { Inter } from "next/font/google";
import { ThemeProvider } from "@/components/theme-provider";

const inter = Inter({ subsets: ["latin"] });

export const metadata = {
  title: "VintLang - Modern Programming Made Simple",
  description:
    "A powerful programming language built with Go, featuring intuitive syntax, built-in networking, and comprehensive time operations.",
  keywords:
    "VintLang, Go Programming Language, Modern Programming, Networking, Time Operations",
  author: "Tachera Sasi",
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en">
      <head>
        <meta charSet="utf-8" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <title>{metadata.title}</title>
        <meta name="description" content={metadata.description} />
        <meta name="keywords" content={metadata.keywords} />
        <meta name="author" content={metadata.author} />
        {/* Open Graph / Social */}
        <meta property="og:title" content={metadata.title} />
        <meta property="og:description" content={metadata.description} />
        <meta property="og:type" content="website" />
        <link rel="icon" href="/favicon.ico" />
      </head>
      <body className={`${inter.className}  mx-auto`}>
        <a href="#content" className="skip-link">
          Skip to content
        </a>
        <ThemeProvider attribute="class" defaultTheme="system" enableSystem>
          <main id="content">{children}</main>
          <Analytics />
        </ThemeProvider>
      </body>
    </html>
  );
}
