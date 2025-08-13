import Bundler from "@/components/docs/bundler";
import SectionHeader from "@/components/docs/SectionHeader";
import { Metadata } from "next";

export const metadata: Metadata = {
  title: "Learn - VintLang",
  description:
    "Learn how to use the VintLang ",
};

export default function BundlerPage() {
  return (
    <div className="space-y-10">
      <SectionHeader title="Learn" />
    </div>
  );
}