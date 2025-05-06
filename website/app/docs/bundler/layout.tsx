export const metadata = {
    title: 'VintLang Bundler – Compile & Ship VintLang Code as Binaries',
    description:
      'The VintLang Bundler compiles .vint files into standalone Go executables. Ship portable CLI tools without requiring users to install VintLang or Go.',
    keywords: 'VintLang, VintLang Bundler, .vint compiler, Go binaries, CLI tools, compile VintLang, bundle VintLang, VintLang to binary, standalone executables',
    author: 'Tachera Sasi',
    openGraph: {
      title: 'VintLang Bundler – Compile VintLang Files into Standalone Binaries',
      description:
        'Distribute VintLang programs as self-contained executables with zero runtime dependencies. Perfect for scripting, tooling, and deployment.',
      url: 'https://vintlang.ekilie.com/docs/bundler',
      type: 'website'
    }
  };

export default function BundlerPageLayout({
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
      </head>
      <body>
          {children}
      </body>
    </html>
  );
}
