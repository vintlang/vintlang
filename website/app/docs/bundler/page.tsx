import Bundler from "@/components/docs/bundler";
import { Metadata } from "next";

export const metadata: Metadata = {
  title: "Bundler - VintLang",
  description: "Learn about the VintLang bundler.",
};


export default function BundlerPage() {
    
    return (
        <div className="space-y-10">
            <Bundler />
        </div>
    )

}