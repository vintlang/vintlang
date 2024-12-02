'use client'

import { motion } from 'motion/react'
import { Button } from '@/components/ui/button'
import Link from 'next/link'

export default function CTA() {
  return (
    <section id="get-started" className="container py-24 sm:py-32">
      <motion.div
        className="bg-taupe-900  rounded-lg px-6 py-16 sm:p-16 text-center"
        initial={{ opacity: 0, y: 20 }}
        animate={{ opacity: 1, y: 0 }}
        transition={{ duration: 0.5 }}
      >
        <h2 className="text-3xl font-bold tracking-tight text-white sm:text-4xl">
          Ready to dive into VintLang?
        </h2>
        <p className="mt-6 text-lg leading-8 text-taupe-700 ">
          Start your journey with VintLang today and experience the power of Swahili-inspired programming.
        </p>
        <div className="mt-10 flex items-center flex-wrap justify-center gap-x-6 gap-y-3">
          <Button size="lg" className="bg-taupe-800 text-white hover:bg-taupe-600 ">
            <Link href="/docs">Get Started</Link>
          </Button>
          <Button size="lg" variant="outline" className=" border-taupe-700 border-white  ">
            <Link href="/docs">Learn More</Link>
          </Button>
        </div>
      </motion.div>
    </section>
  )
}
