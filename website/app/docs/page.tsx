import { Metadata } from "next"
import { DocsHeader } from "@/components/docs/docs-header"
import { Features } from "@/components/docs/features"
import { Core } from "@/components/docs/core"

export const metadata: Metadata = {
  title: "Documentation - VintLang",
  description: "Documentation for VintLang - A modern programming language built with Go",
}

export default function DocsPage() {
  return (
    <div className="space-y-10">
      <DocsHeader />
      <Features />
      <Core />
    </div>
  )
}

