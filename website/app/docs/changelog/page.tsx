import Bundler from "@/components/docs/bundler";
import Changelog from "@/components/docs/changelog";
import { Metadata } from "next";

export const metadata: Metadata = {
  title: "Changelog - VintLang",
  description:
    "Stay up-to-date with the latest features and improvements in VintLang. Our changelog covers updates to language features like declarative statements, repeat loops, and function overloading, as well as enhancements to our compiler, VM, and documentation.",
};

export default function ChangelogPage() {
    
    return (
        <div className="space-y-10">
            <Changelog />
        </div>
    )

}