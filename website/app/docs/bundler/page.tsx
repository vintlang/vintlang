import Bundler from "@/components/docs/bundler";
import { Metadata } from "next";

export const metadata: Metadata = {
  title: "Bundler - VintLang",
  description:
    "Learn how to use the VintLang Bundler to compile your .vint source files into standalone Go binaries. This allows you to distribute your VintLang scripts as self-contained executables that can run on any system without requiring the VintLang interpreter or Go to be installed.",
};

export default function BundlerPage() {
  return (
    <div className="space-y-10">
      <Bundler />
    </div>
  );
}