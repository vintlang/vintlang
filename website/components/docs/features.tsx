export function Features() {
    const features = [
      "Minimal core API (2kb)",
      "Many utilities and extensions",
      "TypeScript oriented",
      "Works with Next.js, Waku, Remix, and React Native",
    ]
  
    return (
      <div className="space-y-4">
        <h2 className="text-2xl font-semibold tracking-tight">Features</h2>
        <ul className="list-disc space-y-2 pl-6">
          {features.map((feature) => (
            <li key={feature} className="text-muted-foreground">
              {feature}
            </li>
          ))}
        </ul>
      </div>
    )
  }
  
  