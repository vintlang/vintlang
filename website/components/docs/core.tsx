export function Core() {
    return (
      <div className="space-y-4">
        <h2 className="text-2xl font-semibold tracking-tight">Core</h2>
        <p className="text-muted-foreground">
          VintLang has a very minimal API, exposing only a few exports from the main{" "}
          <code className="relative rounded bg-muted px-[0.3rem] py-[0.2rem] font-mono text-sm">
            vintlang
          </code>{" "}
          bundle. They are split into four categories below.
        </p>
      </div>
    )
  }
  
  