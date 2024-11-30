"use client";

import Link from "next/link";
import { usePathname } from "next/navigation";
import { cn } from "@/lib/utils";
import { Button } from "@/components/ui/button";
import { Search } from "@/components/docs/search";
import {
  Hand,
  FileText,
  BookOpen,
  HelpCircle,
  Github,
  ArrowDownToLine,
} from "lucide-react";

const navigation = [
  {
    name: "Introduction",
    href: "#",
    icon: Hand,
  },
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
  // {
  //   name: "Repository",
  //   href: "https://github.com/ekilie/vint-lang",
  //   icon: Github,
  // },
];

export function Sidebar() {
  const pathname = usePathname();

  return (
    <div className="flex h-full flex-col p-14">
      <div className="mb-6">
        <Link href="/" className="flex items-center space-x-2">
          <span className="text-9xl text-neutral-300 font-bold">Vint</span>
        </Link>
        <p className="mt-1 text-sm text-muted-foreground">
          Modern Programming made simple
        </p>
      </div>
      {/* <Search /> */}
      <nav className="mt-4 flex-1">
        {navigation.map((item) => {
          const isActive = pathname === item.href;
          return (
            <Button
              key={item.name}
              variant={isActive ? "secondary" : "ghost"}
              className={cn(
                "mb-2 w-full p-6 justify-start shadow-lg rounded-lg border-2",
                isActive ? "bg-accent" : "hover:bg-accent/50"
              )}
              asChild
            >
              <Link href={item.href}>
                <item.icon className="mr-2 h-4 w-4" />
                {item.name}
              </Link>
            </Button>
          );
        })}
      </nav>
      <div className="mt-auto space-y-2 text-xs text-muted-foreground">
        <p>Language by Tachera Sasi</p>
        <p>
          From <Link href="https://ekilie.com">ekilie</Link>
        </p>
        {/* <p className="flex items-center">
          site by{" "}
          <span className="ml-1">
            <span className="text-[#977196]">T</span>
            <span className="text-[#98aa98]">a</span>
            <span className="text-[#00FF00]">c</span>
            <span className="text-[#FF1CF7]">h</span>
            <span className="text-[#00FF00]">era</span>
          </span>
        </p> */}
      </div>
    </div>
  );
}
