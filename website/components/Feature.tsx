"use client";

import React, { ForwardRefExoticComponent, RefAttributes } from "react";
import { motion } from "motion/react";
import { LucideProps } from "lucide-react";

interface VintLangFeature {
  name: string;
  description: string;
  icon: ForwardRefExoticComponent<
    Omit<LucideProps, "ref"> & RefAttributes<SVGSVGElement>
  >;
}

interface FeatureProps {
  feature: VintLangFeature;
  index: number;
}

const Feature: React.FC<FeatureProps> = ({ feature, index }) => {
  return (
    <motion.div
      className="rounded-lg bg-neutral-900 p-6 text-center md:text-left flex flex-col md:flex-row items-center shadow-lg hover:shadow-xl transition-shadow duration-300 cursor-pointer"
      initial={{ opacity: 0, y: 30 }}
      animate={{ opacity: 1, y: 0 }}
      transition={{ duration: 0.6, delay: index * 0.1 }}
    >
      <div className="flex items-center px-1 justify-center h-16 w-16 rounded-full bg-neutral-100 text-neutral-900 dark:bg-neutral-800 dark:text-neutral-100 shadow-md">
        <feature.icon className="w-8 h-8" aria-label={`${feature.name} icon`} />
      </div>
      <div className="mt-4 md:mt-0 md:ml-6 text-center md:text-left">
        <h3 className="text-xl font-bold text-white">{feature.name}</h3>
        <p className="mt-2 text-sm text-neutral-400">{feature.description}</p>
      </div>
    </motion.div>
  );
};

export default Feature;
