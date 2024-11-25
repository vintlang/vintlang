'use client'

import { useEffect } from 'react'
import { motion, useAnimation } from 'framer-motion'
import { useInView } from 'react-intersection-observer'
import Header from './header'
import Hero from './hero'
import Features from './features'
import CodeExample from './code-example'
import CTA from './cta'
import Footer from './footer'
import { ThemeProvider } from '@/components/theme-provider'

export function LandingPage() {
  const controls = useAnimation()
  const [ref, inView] = useInView()

  useEffect(() => {
    if (inView) {
      controls.start('visible')
    }
  }, [controls, inView])

  return (
    <ThemeProvider attribute="class" defaultTheme="system" enableSystem>
      <div className="min-h-screen bg-background text-foreground">
        <Header />
        <main className="overflow-hidden">
          <Hero />
          <div className="px-4 sm:px-6 lg:px-8">
            <motion.div
              ref={ref}
              animate={controls}
              initial="hidden"
              variants={{
                visible: { opacity: 1, y: 0 },
                hidden: { opacity: 0, y: 50 },
              }}
              transition={{ duration: 0.5, ease: 'easeOut' }}
            >
              <Features />
            </motion.div>
            <CodeExample />
            <CTA />
          </div>
        </main>
        <Footer />
      </div>
    </ThemeProvider>
  )
}

