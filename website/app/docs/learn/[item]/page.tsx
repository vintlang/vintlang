"use client";

import SectionHeader from '@/components/docs/SectionHeader';
import { fetchMarkdown } from '@/lib/utils';
import { useParams } from 'next/navigation'
import React from 'react'
import Markdown from 'react-markdown';

const Page = () => {
  const params = useParams<{ item: string }>()
  console.log(params)
  const markdown = await fetchMarkdown("docs/bundler.md");
    return (
      <div className="p-6" id="docs">
        <SectionHeader title="Bundler" />
        <Markdown>{markdown}</Markdown>
      </div>
    )
};

export default Page
