"use client";

import { Clipboard, ClipboardCheck } from "lucide-react";
import Link from "next/link";
import React, { memo, useState } from "react";
import ReactMarkdown from "react-markdown";
import remarkGfm from "remark-gfm";

const NonMemoizedMarkdown = ({ children }: { children: string }) => {
  const [isCopied,setIsCopied] = useState(false)
  const copyToClipboard = (text: string) => {
    navigator.clipboard.writeText(text).then(() => {
      setIsCopied(true)
    });
  };

  const components = {
    h1: ({ node, children, ...props }: any) => (
      <h1
        className="text-3xl font-bold mb-4 mt-6 text-neutral-800 dark:text-neutral-100 border-b border-neutral-300 dark:border-neutral-700 pb-2"
        {...props}
      >
        {children}
      </h1>
    ),
    h2: ({ node, children, ...props }: any) => (
      <h2
        className="text-2xl font-semibold mb-3 mt-5 text-neutral-800 dark:text-neutral-200"
        {...props}
      >
        {children}
      </h2>
    ),
    h3: ({ node, children, ...props }: any) => (
      <h3
        className="text-xl font-medium mb-2 mt-4 text-neutral-800 dark:text-neutral-200"
        {...props}
      >
        {children}
      </h3>
    ),
    p: ({ node, children, ...props }: any) => (
      <p
        className="text-base leading-6 mb-4 text-neutral-700 dark:text-neutral-300"
        {...props}
      >
        {children}
      </p>
    ),
    code: ({ node, inline, className, children, ...props }: any) => {
      const match = /language-(\w+)/.exec(className || "");
      const codeContent = String(children).trim();

      return !inline && match ? (
        <div className="relative group">
          <pre
            {...props}
            className={`${className} text-sm w-[90%] md:max-w-full overflow-x-auto bg-zinc-100 p-4 rounded-lg mt-3 dark:bg-zinc-800`}
          >
            <code className={match[1]}>{children}</code>
          </pre>
          <button
            onClick={() => copyToClipboard(codeContent)}
            className="absolute top-2 right-2 bg-neutral-200 dark:bg-neutral-700 text-sm px-2 py-1 rounded opacity-0 group-hover:opacity-100 transition-opacity"
          >
            {!isCopied ? (<Clipboard />):(<ClipboardCheck />)}
          </button>
        </div>
      ) : (
        <code
          className={`${className} text-sm bg-zinc-100 dark:bg-zinc-800 py-0.5 px-1 rounded-md`}
          {...props}
        >
          {children}
        </code>
      );
    },
    ol: ({ node, children, ...props }: any) => (
      <ol
        className="list-decimal list-outside ml-6 mb-4 text-neutral-700 dark:text-neutral-300"
        {...props}
      >
        {children}
      </ol>
    ),
    ul: ({ node, children, ...props }: any) => (
      <ul
        className="list-disc list-outside ml-6 mb-4 text-neutral-700 dark:text-neutral-300"
        {...props}
      >
        {children}
      </ul>
    ),
    li: ({ node, children, ...props }: any) => (
      <li className="py-1" {...props}>
        {children}
      </li>
    ),
    blockquote: ({ node, children, ...props }: any) => (
      <blockquote
        className="border-l-4 border-neutral-300 dark:border-neutral-700 pl-4 italic text-neutral-600 dark:text-neutral-400 mb-4"
        {...props}
      >
        {children}
      </blockquote>
    ),
    strong: ({ node, children, ...props }: any) => (
      <span className="font-semibold" {...props}>
        {children}
      </span>
    ),
    a: ({ node, children, ...props }: any) => (
      <Link
        className="text-green-500 hover:underline"
        target="_blank"
        rel="noreferrer"
        {...props}
      >
        {children}
      </Link>
    ),
    hr: ({ node, ...props }: any) => (
      <hr
        className="border-t border-neutral-300 dark:border-neutral-700 my-6"
        {...props}
      />
    ),
    table: ({ node, children, ...props }: any) => (
      <table
        className="table-auto border-collapse border border-neutral-300 dark:border-neutral-700 w-full text-left mb-4"
        {...props}
      >
        {children}
      </table>
    ),
    th: ({ node, children, ...props }: any) => (
      <th
        className="border border-neutral-300 dark:border-neutral-700 px-4 py-2 bg-neutral-100 dark:bg-neutral-800 font-semibold"
        {...props}
      >
        {children}
      </th>
    ),
    td: ({ node, children, ...props }: any) => (
      <td
        className="border border-neutral-300 dark:border-neutral-700 px-4 py-2"
        {...props}
      >
        {children}
      </td>
    ),
  };

  return (
    <ReactMarkdown remarkPlugins={[remarkGfm]} components={components}>
      {children}
    </ReactMarkdown>
  );
};

export const Markdown = memo(
  NonMemoizedMarkdown,
  (prevProps, nextProps) => prevProps.children === nextProps.children
);
