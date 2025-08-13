import SectionHeader from '@/components/docs/SectionHeader';
import { fetchMarkdown } from '@/lib/utils';
import React from 'react'
import Markdown from 'react-markdown';



const LearnItem = async () => {
  const markdown = await fetchMarkdown("docs/bundler.md");
  return (
    <div className="p-6" id="docs">
      <SectionHeader title="Bundler" />
      <Markdown>{markdown}</Markdown>
    </div>
  )
}

export default LearnItem
