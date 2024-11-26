import { clsx, type ClassValue } from "clsx"
import { twMerge } from "tailwind-merge"
import { Clock, Globe, Network, Type, Variable, Wand2 } from 'lucide-react'

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs))
}

export const features = [
  {
    name: 'Simple Syntax',
    description: 'Write clean, readable code with an intuitive syntax inspired by modern programming practices.',
    icon: Type,
  },
  {
    name: 'Built-in Networking',
    description: 'Powerful networking capabilities with the built-in net module for HTTP operations.',
    icon: Network,
  },
  {
    name: 'Time Operations',
    description: 'Comprehensive time manipulation and formatting with the native time module.',
    icon: Clock,
  },
  {
    name: 'Dynamic Typing',
    description: 'Flexible type system with built-in type conversion and checking capabilities.',
    icon: Variable,
  },
  {
    name: 'Go-Powered',
    description: 'Built with Go, ensuring high performance and reliable execution.',
    icon: Wand2,
  },
  {
    name: 'Global Community',
    description: 'Join a growing community of developers building with VintLang.',
    icon: Globe,
  },
]
