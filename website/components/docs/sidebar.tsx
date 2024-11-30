"use client"

import Link from "next/link"
import { usePathname } from "next/navigation"
import { cn } from "@/lib/utils"
import { Button } from "@/components/ui/button"
import { Search } from "@/components/docs/search"
import { FileText, BookOpen, HelpCircle, Github,ArrowDownToLine } from 'lucide-react'

const navigation = [
  {
    name: "Installation",
    href: "#installation",
    icon: ArrowDownToLine,
  },
  {
    name: "Documentation",
    href: "#docs",
    icon: FileText,
  },
  {
    name: "Tutorial",
    href: "#tutorial",
    icon: BookOpen,
  },
  {
    name: "Sponsor",
    href: "#sponsor",
    icon: HelpCircle,
  },
  {
    name: "Repository",
    href: "https://github.com/ekilie/vint-lang",
    icon: Github,
  },
]

export function Sidebar() {
  const pathname = usePathname()

  return (
    <div className="flex h-full flex-col px-6 py-8">
      <div className="mb-8">
        <Link href="/" className="flex items-center space-x-2">
          <span className="text-8xl text-neutral-300 font-bold">Vint</span>
        </Link>
        <p className="mt-1 text-sm text-muted-foreground">
          Primitive and flexible state management for React
        </p>
      </div>
      {/* <Search /> */}
      <nav className="mt-4 flex-1">
        {navigation.map((item) => {
          const isActive = pathname === item.href
          return (
            <Button
              key={item.name}
              variant={isActive ? "secondary" : "ghost"}
              className={cn(
                "mb-1 w-full justify-start shadow-lg",
                isActive ? "bg-accent" : "hover:bg-accent/50"
              )}
              asChild
            >
              <Link href={item.href}>
                <item.icon className="mr-2 h-4 w-4" />
                {item.name}
              </Link>
            </Button>
          )
        })}
      </nav>
      <div className="mt-auto space-y-2 text-xs text-muted-foreground">
        <p>library by Daishi Kato</p>
        <p>art by Jessie Waters</p>
        <p className="flex items-center">
          site by{" "}
          <span className="ml-1">
            <span className="text-[#FF1CF7]">c</span>
            <span className="text-[#00FF00]">o</span>
            <span className="text-[#00FF00]">d</span>
            <span className="text-[#FF1CF7]">y</span>
            <span className="text-[#00FF00]">code</span>
          </span>
        </p>
      </div>
    </div>
  )
}

