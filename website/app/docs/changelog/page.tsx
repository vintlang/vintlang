import Bundler from "@/components/docs/bundler";
import Changelog from "@/components/docs/changelog";
import { Metadata } from "next";

export const metadata: Metadata = {
  title: "Changelog - VintLang",
  description: "VintLang Changelog",
};

export default function ChangelogPage() {
    
    return (
        <div className="space-y-10">
            <Changelog />
        </div>
    )

}