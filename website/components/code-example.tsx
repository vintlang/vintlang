'use client'

import { useEffect, useState } from 'react'
import { motion } from 'motion/react'
import { Tabs, TabsContent, TabsList, TabsTrigger } from '@/components/ui/tabs'
import { codeExamples } from '@/lib/codeExample'
import { Prism as SyntaxHighlighter } from 'react-syntax-highlighter'
import { dracula ,dark ,oneDark } from 'react-syntax-highlighter/dist/esm/styles/prism'


export default function CodeExample() {
  const [activeTab, setActiveTab] = useState('basics')

  useEffect(() => {
    console.log(activeTab)
  }, [activeTab])

  return (
    <section id="code-example" className="container py-24 sm:py-32">
      <h2 className="text-3xl font-bold tracking-tight text-center mb-12">
        VintLang in Action
      </h2>
      <motion.div
        initial={{ opacity: 0, y: 20 }}
        animate={{ opacity: 1, y: 0 }}
        transition={{ duration: 0.5 }}
      >
        <Tabs defaultValue="basics" className="w-full" onValueChange={setActiveTab}>
          <TabsList className="grid w-full grid-cols-4">
            <TabsTrigger value="basics">Basics</TabsTrigger>
            <TabsTrigger value="functions">Functions</TabsTrigger>
            <TabsTrigger value="jsonModule">JSON</TabsTrigger>
            <TabsTrigger value="timeAndNet">Time & Network</TabsTrigger>
          </TabsList>
          {Object.entries(codeExamples).map(([key, code]) => (
            <TabsContent key={key} value={key}>
              <motion.div
                initial={{ opacity: 0 }}
                animate={{ opacity: 1 }}
                transition={{ duration: 0.5 }}
                className="relative"
              >
                <pre className="p-4 rounded-lg bg-muted overflow-x-auto">
                  <SyntaxHighlighter language="javascript" style={oneDark}>
                    {code}
                  </SyntaxHighlighter>
                </pre>
              </motion.div>
            </TabsContent>
          ))}
        </Tabs>
      </motion.div>
    </section>
  )
}
