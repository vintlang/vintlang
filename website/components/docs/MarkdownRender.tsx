import React from 'react';
import ReactMarkdown from 'react-markdown';
import rehypeHighlight from 'rehype-highlight';

interface MarkdownRendererProps{
    markdown:string
}

export default function MarkdownRenderer({ markdown }:MarkdownRendererProps) {
  return (
    <ReactMarkdown rehypePlugins={[rehypeHighlight]}>
      {markdown}
    </ReactMarkdown>
  );
}
