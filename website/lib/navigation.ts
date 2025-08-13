import { Hand, ArrowDownToLine, FileText, BookOpen, HelpCircle, PackageCheck, ReplaceAll, GithubIcon } from "lucide-react";

export const navigation = [
    {
      name: "Introduction",
      href: "/docs#",
      icon: Hand,
      blank:false
    },
    {
      name: "Installation",
      href: "/docs#installation",
      icon: ArrowDownToLine,
      blank:false
    },
    {
      name: "Documentation",
      href: "/docs#docs",
      icon: FileText,
      blank:false
    },
    {
      name: "Learn",
      href: "/docs/learn",
      icon: BookOpen,
      blank:false
    },
    {
      name: "Bundler",
      href: "/docs/bundler",
      icon: PackageCheck,
      blank:false
    },
    {
      name: "Changelog",
      href: "/docs/changelog",
      icon: ReplaceAll,
      blank:false
    },
    {
      name: "Github",
      href: "https://github.com/vintlang/vintlang",
      icon: GithubIcon,
      blank:true
    },
    {
      name: "Sponsor",
      href: "/docs#sponsor",
      icon: HelpCircle,
      blank:false
    },
  ];