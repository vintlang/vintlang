"use client";

import LearnItem from "@/components/docs/learn-item";
import SectionHeader from "@/components/docs/SectionHeader";
import { fetchMarkdown } from "@/lib/utils";
import { useParams } from "next/navigation";
import React from "react";
import Markdown from "react-markdown";

const Page = () => {
  const params = useParams<{ item: string }>();
  console.log(params);
  return (
    <div className="space-y-10">
      <LearnItem item={params.item} />
    </div>
  );
};

export default Page;
