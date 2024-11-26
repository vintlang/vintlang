'use client'

import { useEffect, useState } from 'react'
import { motion } from 'motion/react'
import { Tabs, TabsContent, TabsList, TabsTrigger } from '@/components/ui/tabs'

const codeExamples = {
  basics: `// Basic VintLang Operations
let name = "VintLang"
s = name.split("") 
for i in s { 
    print(i)
}

// Type conversion
age = "10"
convert(age, "INTEGER")
print(type(age))`,
  
  functions: `// Function Definition in VintLang
let printDetails = func(name, age, height) {
    print("My name is " + name + 
          ", I am " + age + 
          " years old, and my height is " + 
          height + " feet.")
}

// Function call
printDetails("VintLang", "10", "6.0")`,
  
  timeAndNet: `// Time and Network Operations
import time
import net

// Time operations
print(time.format(time.now(), "02-01-2006 15:04:05"))
print(time.add(time.now(), "1h"))

// Network request
let res = net.get("https://example.com")
print(res)`,
}

export default function CodeExample() {
  const [activeTab, setActiveTab] = useState('basics')

  useEffect(()=>{
    console.log(activeTab)
  },[activeTab])

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
          <TabsList className="grid w-full grid-cols-3">
            <TabsTrigger value="basics">Basics</TabsTrigger>
            <TabsTrigger value="functions">Functions</TabsTrigger>
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
                  <code className="text-sm font-mono">{code}</code>
                </pre>
              </motion.div>
            </TabsContent>
          ))}
        </Tabs>
      </motion.div>
    </section>
  )
}

