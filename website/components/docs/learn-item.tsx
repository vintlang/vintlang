import SectionHeader from '@/components/docs/SectionHeader';
import { fetchMarkdown } from '@/lib/utils';
import React from 'react'
import Markdown from 'react-markdown';

interface LearnItemProps { 
    item: string;
}

const LearnItem = async ({ item }: LearnItemProps) => {
  const markdown = await fetchMarkdown(`docs/${item}.md`);
  return (
    <div className="p-6" id="docs">
      <SectionHeader title={item} />
      <Markdown>{markdown}</Markdown>
    </div>
  )
}

export default LearnItem
