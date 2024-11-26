"use client";

import { features } from "@/lib/utils";
import Feature from "./Feature";

export default function Features() {
  return (
    <section id="features" className="container mx-auto px-6 py-16">
      <h2 className="text-center text-4xl font-extrabold text-white mb-12 tracking-tight">
        Why Choose VintLang?
      </h2>
      <div className="grid gap-10 sm:grid-cols-2 lg:grid-cols-3">
        {features.map((feature, index) => (
          <Feature key={index} feature={feature} index={index} />
        ))}
      </div>
    </section>
  );
}
