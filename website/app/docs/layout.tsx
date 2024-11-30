import { Sidebar } from "@/components/docs/sidebar"
import { ScrollArea } from "@/components/ui/scroll-area"

interface DocsLayoutProps {
  children: React.ReactNode
}

export default function DocsLayout({ children }: DocsLayoutProps) {
  return (
    <div className="flex min-h-screen max-w-screen-xl mx-auto">
      <aside className="fixed left-0 top-0 z-30 h-screen w-[380px] border-r bg-background">
        <ScrollArea className="h-full">
          <Sidebar />
        </ScrollArea>
      </aside>
      <main className="flex-1 pl-[250px]">
        <div className="container max-w-3xl py-6 lg:py-10">
          {children}
        </div>
      </main>
    </div>
  )
}

