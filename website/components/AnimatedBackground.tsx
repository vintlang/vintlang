import { motion } from 'motion/react'

export const AnimatedBackground = () => {
    return (
      <div className="absolute inset-0 overflow-hidden opacity-20 dark:opacity-30">
        <svg className="w-full h-full" xmlns="http://www.w3.org/2000/svg">
          <defs>
            <filter id="goo">
              <feGaussianBlur in="SourceGraphic" stdDeviation="10" result="blur" />
              <feColorMatrix in="blur" mode="matrix" values="1 0 0 0 0  0 1 0 0 0  0 0 1 0 0  0 0 0 18 -8" result="goo" />
            </filter>
          </defs>
          <g filter="url(#goo)">
            {[...Array(20)].map((_, i) => (
              <motion.circle
                key={i}
                cx={Math.random() * 100 + '%'}
                cy={Math.random() * 100 + '%'}
                r={Math.random() * 50 + 10}
                fill={`hsl(${Math.random() * 360}, 70%, 50%)`}
                initial={{ scale: 0 }}
                animate={{
                  scale: [1, 1.2, 1],
                  x: [0, Math.random() * 100 - 50, 0],
                  y: [0, Math.random() * 100 - 50, 0],
                }}
                transition={{
                  duration: Math.random() * 5 + 5,
                  repeat: Infinity,
                  ease: 'easeInOut',
                }}
              />
            ))}
          </g>
        </svg>
      </div>
    )
}