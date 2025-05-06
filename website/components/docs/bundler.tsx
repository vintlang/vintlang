import { fetchMarkdown } from "@/lib/utils";
import SectionHeader from "./SectionHeader";
import { Markdown } from "../Markdown";

export default async function Bundler(){
  const markdown = await fetchMarkdown("docs/bundler.md");
  return (
    <div className="p-6" id="docs">
      <SectionHeader title="Bundler" />
      <Markdown>{markdown}</Markdown>
    </div>
  )
}