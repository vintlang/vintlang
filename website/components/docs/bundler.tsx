import { fetchMarkdown } from "@/lib/utils";
import SectionHeader from "./SectionHeader";
import { Markdown } from "../Markdown";

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

  
export default async function Bundler(){
  const markdown = await fetchMarkdown("docs/bundler.md");
  return (
    <div className="p-6" id="docs">
      <SectionHeader title="Bundler" />
      <Markdown>{markdown}</Markdown>
    </div>
  )
}