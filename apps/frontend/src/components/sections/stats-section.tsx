"use client"
import { motion } from "motion/react"

const stats = [
  { value: "RHCSA", label: "Red Hat Certified",    sublabel: "System Administrator" },
  { value: "AWS",   label: "Cloud Infrastructure", sublabel: "Production Experience" },
  { value: "k8s",   label: "Kubernetes",            sublabel: "Container Orchestration" },
  { value: "GitOps",label: "Workflow Automation",  sublabel: "ArgoCD + Flux" },
]

export function StatsSection() {
  return (
    <section className="border-t border-border/40 bg-muted/20">
      <div className="container mx-auto max-w-4xl px-4 py-16">
        <div className="grid grid-cols-2 gap-8 md:grid-cols-4">
          {stats.map((stat, i) => (
            <motion.div
              key={stat.value}
              initial={{ opacity: 0, y: 20 }}
              whileInView={{ opacity: 1, y: 0 }}
              viewport={{ once: true }}
              transition={{ duration: 0.5, delay: i * 0.1 }}
              className="text-center"
            >
              <div className="font-mono text-2xl font-bold text-primary mb-1">{stat.value}</div>
              <div className="text-sm font-medium text-foreground">{stat.label}</div>
              <div className="text-xs text-muted-foreground mt-0.5">{stat.sublabel}</div>
            </motion.div>
          ))}
        </div>
      </div>
    </section>
  )
}
