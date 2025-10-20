import SectionHeader from "@/components/docs/SectionHeader";
import { getCategorizedDocs, getLearnItems } from "@/lib/docs";
import { Metadata } from "next";

export const metadata: Metadata = {
  title: "Learn - VintLang",
  description:
    "Learn how to use the VintLang programming language with comprehensive documentation and examples.",
};

export default async function LearnPage() {
  const categorizedDocs = await getCategorizedDocs();

  return (
    <div className="space-y-10">
      <SectionHeader title="Learn VintLang" />
      <div className="prose max-w-none">
        <p className="text-lg text-muted-foreground mb-8">
          Explore comprehensive documentation for VintLang features, from basic
          syntax to advanced modules.
        </p>
      </div>

      {Object.entries(categorizedDocs).map(([category, items]) => (
        <div key={category} className="space-y-4">
          <h2 className="text-xl font-bold text-foreground border-b pb-2">
            {category}
          </h2>
          <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
            {items.map((item) => (
              <div
                key={item.title}
                className="p-4 border rounded-lg hover:shadow-lg transition-all hover:border-border/80 bg-card"
              >
                <a
                  href={item.href}
                  className="block focus:outline-none focus-visible:ring-2 focus-visible:ring-ring rounded"
                >
                  <h3 className="text-lg font-semibold text-foreground hover:underline mb-2">
                    {item.title}
                  </h3>
                  <p className="text-sm text-muted-foreground line-clamp-2">
                    {item.description}
                  </p>
                </a>
              </div>
            ))}
          </div>
        </div>
      ))}
    </div>
  );
}
