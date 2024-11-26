'use client'

import { motion } from 'motion/react'
import { Button } from '@/components/ui/button'

export default function CTA() {
  return (
    <section id="get-started" className="container py-24 sm:py-32">
      <motion.div
        className="bg-taupe-500 dark:bg-taupe-600 rounded-lg px-6 py-16 sm:p-16 text-center"
        initial={{ opacity: 0, y: 20 }}
        animate={{ opacity: 1, y: 0 }}
        transition={{ duration: 0.5 }}
      >
        <h2 className="text-3xl font-bold tracking-tight text-white sm:text-4xl">
          Ready to dive into VintLang?
        </h2>
        <p className="mt-6 text-lg leading-8 text-taupe-100">
          Start your journey with VintLang today and experience the power of Swahili-inspired programming.
        </p>
        <div className="mt-10 flex items-center justify-center gap-x-6">
          <Button size="lg" className="bg-white text-taupe-600 hover:bg-taupe-50">
            Get Started
          </Button>
          <Button size="lg" variant="outline" className="text-white border-white hover:bg-taupe-600">
            Learn More
          </Button>
        </div>
      </motion.div>
    </section>
  )
}

