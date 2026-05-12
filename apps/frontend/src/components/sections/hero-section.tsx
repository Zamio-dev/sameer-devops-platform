"use client"

import { motion, type Variants } from "motion/react"
import { Terminal, ArrowRight } from "lucide-react"
import Link from "next/link"
import { Button } from "@/components/ui/button"
import { Badge } from "@/components/ui/badge"

const container: Variants = {
  hidden: { opacity: 0 },
  visible: { opacity: 1, transition: { staggerChildren: 0.1, delayChildren: 0.2 } },
}

const item: Variants = {
  hidden: { opacity: 0, y: 20 },
  visible: { opacity: 1, y: 0, transition: { duration: 0.5, ease: "easeOut" } },
}

const badges = [
  "Kubernetes", "GitOps", "ArgoCD", "Terraform",
  "Prometheus", "RHCSA", "Docker", "GitHub Actions",
]

export function HeroSection() {
  return (
    <section className="relative min-h-[90vh] flex items-center justify-center overflow-hidden">

      {/* Grid background */}
      <div
        className="absolute inset-0 opacity-[0.03] pointer-events-none"
        style={{
          backgroundImage:
            "linear-gradient(oklch(0.87 0.26 138.6 / 50%) 1px, transparent 1px)," +
            "linear-gradient(90deg, oklch(0.87 0.26 138.6 / 50%) 1px, transparent 1px)",
          backgroundSize: "50px 50px",
        }}
      />

      {/* Glow orbs */}
      <div className="absolute top-1/4 left-1/4 h-96 w-96 rounded-full bg-primary/5 blur-3xl pointer-events-none" />
      <div className="absolute bottom-1/4 right-1/4 h-96 w-96 rounded-full bg-blue-500/5 blur-3xl pointer-events-none" />

      <motion.div
        variants={container}
        initial="hidden"
        animate="visible"
        className="relative z-10 container mx-auto max-w-4xl px-4 text-center"
      >
        {/* Status badge */}
        <motion.div variants={item} className="mb-6 flex justify-center">
          <Badge variant="outline" className="px-4 py-1.5 text-xs font-mono border-green-500/30 text-green-400 bg-green-500/5">
            <span className="mr-2 inline-block h-1.5 w-1.5 rounded-full bg-green-400 animate-pulse" />
            Platform Status: All Systems Operational
          </Badge>
        </motion.div>

        {/* Heading */}
        <motion.div variants={item} className="mb-6">
          <h1 className="text-5xl font-bold tracking-tight sm:text-6xl md:text-7xl">
            <span className="block text-foreground">Sameer Malik RP</span>
            <span className="block font-mono text-primary mt-2">DevOps Engineer</span>
          </h1>
        </motion.div>

        {/* Terminal prompt */}
        <motion.div variants={item} className="mb-8">
          <div className="inline-block rounded-lg border border-border/50 bg-muted/50 px-4 py-2 font-mono text-sm text-muted-foreground">
            <span className="text-green-400">$</span>{" "}
            <span className="text-primary">kubectl get</span> pods --all-namespaces
            <span className="animate-[blink_1s_step-end_infinite] text-primary ml-1">▊</span>
          </div>
        </motion.div>

        {/* Description */}
        <motion.p variants={item} className="mx-auto mb-10 max-w-2xl text-lg text-muted-foreground leading-relaxed">
          Building enterprise-grade cloud infrastructure. RHCSA certified engineer specialising in{" "}
          <span className="text-foreground font-medium">Kubernetes orchestration</span>,{" "}
          <span className="text-foreground font-medium">GitOps workflows</span>, and{" "}
          <span className="text-foreground font-medium">security-first platform engineering</span>.
        </motion.p>

        {/* CTAs */}
        <motion.div variants={item} className="flex flex-col sm:flex-row gap-4 justify-center mb-12">
          <Link href="/projects">
            <Button size="lg" className="gap-2 w-full sm:w-auto">
              View Projects <ArrowRight className="h-4 w-4" />
            </Button>
          </Link>
          <Link href="/terminal">
            <Button variant="outline" size="lg" className="gap-2 font-mono w-full sm:w-auto border-primary/30 hover:border-primary hover:bg-primary/10">
              <Terminal className="h-4 w-4" />
              ./explore --interactive
            </Button>
          </Link>
        </motion.div>

        {/* Tech badges */}
        <motion.div variants={item} className="flex flex-wrap gap-2 justify-center">
          {badges.map((label) => (
            <Badge key={label} variant="secondary" className="text-xs font-mono px-3 py-1">
              {label}
            </Badge>
          ))}
        </motion.div>
      </motion.div>
    </section>
  )
}
