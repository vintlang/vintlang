'use client'

import { useEffect, useRef } from 'react'
import { motion, useAnimation, useInView } from 'motion/react'
import { Button } from '@/components/ui/button'
import { ArrowRight, Code2 } from 'lucide-react'
import { AnimatedBackground } from './AnimatedBackground'
import Link from 'next/link'


export default function Hero() {
  const controls = useAnimation()
  const ref = useRef(null)
  const inView = useInView(ref)

  useEffect(() => {
    if (inView) {
      controls.start('visible')
    }
  }, [controls, inView])

  return (
    <section ref={ref} className="relative min-h-screen flex items-center justify-center overflow-hidden">
      <AnimatedBackground />
      <div className="container relative z-10 flex flex-col items-center justify-center min-h-screen py-12 px-4 sm:px-6 lg:px-8">
        <motion.div
          initial={{ opacity: 0, y: -20 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ duration: 0.5 }}
          className="flex items-center gap-2 text-taupe-850 dark:text-taupe-400"
        >
          <Code2 className="w-8 h-8" />
          <span className="text-3xl font-bold">VintLang</span>
        </motion.div>
        <motion.h1
          className="mt-6 text-4xl font-extrabold tracking-tight sm:text-5xl md:text-6xl lg:text-7xl text-center"
          initial="hidden"
          animate={controls}
          variants={{
            hidden: { opacity: 0, y: -20 },
            visible: {
              opacity: 1,
              y: 0,
              transition: {
                delay: 0.2,
                staggerChildren: 0.1,
              },
            },
          }}
        >
          {['Modern', 'Programming', 'Made', 'Simple'].map((word, index) => (
            <motion.span
              key={index}
              className="inline-block"
              variants={{
                hidden: { opacity: 0, y: 20 },
                visible: { opacity: 1, y: 0 },
              }}
            >
              <span className={index % 2 === 0 ? 'text-taupe-800 dark:text-taupe-400' : ''}>
                {word}{' '}
              </span>
            </motion.span>
          ))}
        </motion.h1>
        <motion.p
          className="mt-6 text-xl text-muted-foreground max-w-2xl text-center"
          initial={{ opacity: 0, y: 20 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ duration: 0.5, delay: 0.8 }}
        >
          A powerful programming language built with Go, featuring intuitive syntax, built-in networking, and comprehensive time operations.
        </motion.p>
        <motion.div
          className="mt-10 flex flex-col sm:flex-row items-center justify-center gap-4 sm:gap-x-6"
          initial={{ opacity: 0, y: 20 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ duration: 0.5, delay: 1 }}
        >
          <Button size="lg" className="bg-taupe-900 hover:bg-taupe-800 text-white text-lg px-8 py-6 w-full sm:w-auto">
            <Link href="/docs">Get Started</Link>
            <ArrowRight className="ml-2 h-5 w-5" />
          </Button>
          <Button size="lg" variant="outline" className="text-lg px-8 py-6 w-full sm:w-auto">
            <Link href="#code-example">View Examples</Link>
          </Button>
        </motion.div>
        <motion.div
          className="mt-16 text-sm text-muted-foreground text-center"
          initial={{ opacity: 0 }}
          animate={{ opacity: 1 }}
          transition={{ duration: 0.5, delay: 1.2 }}
        >
          Built with ❤️ by Tachera Sasi, CEO at <a href="https://ekilie.com" target="_blank" rel="noopener noreferrer" className="underline hover:text-taupe-500 dark:hover:text-taupe-400">ekilie.com</a>
        </motion.div>
      </div>
    </section>
  )
}

