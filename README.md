# Sameer Malik RP — DevOps Portfolio Platform

[![CI/CD Pipeline](https://github.com/YOUR_USERNAME/sameer-devops-platform/actions/workflows/ci.yml/badge.svg)](https://github.com/YOUR_USERNAME/sameer-devops-platform/actions)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

> An enterprise-grade SaaS portfolio platform built to showcase production DevOps engineering capabilities.

## What This Is

This is not a simple portfolio website. This is a **live, running SaaS platform** that demonstrates:

- **Kubernetes** orchestration with k3s locally and EKS in production
- **GitOps** workflows with ArgoCD
- **CI/CD pipelines** with GitHub Actions and Jenkins
- **Infrastructure as Code** with Terraform and Ansible
- **Observability** with Prometheus, Grafana, Loki, and Tempo
- **Security-first** engineering with Trivy, Falco, and OWASP ZAP
- **Cloud-native** patterns: service mesh, horizontal scaling, disaster recovery

## Architecture Overview
┌─────────────────────────────────────────────────────────┐
│                    Cloudflare (CDN/WAF)                  │
└─────────────────────────┬───────────────────────────────┘
│
┌─────────────────────────▼───────────────────────────────┐
│              Kubernetes Cluster (k3s/EKS)                │
│  ┌──────────────┐  ┌──────────────┐  ┌───────────────┐  │
│  │  Next.js UI  │  │  Go Backend  │  │  PostgreSQL   │  │
│  │   Port 3000  │  │   Port 8080  │  │   Port 5432   │  │
│  └──────────────┘  └──────────────┘  └───────────────┘  │
│  ┌──────────────┐  ┌──────────────┐  ┌───────────────┐  │
│  │    Redis     │  │  Prometheus  │  │    Grafana    │  │
│  │   Cache      │  │  Metrics     │  │   Dashboards  │  │
│  └──────────────┘  └──────────────┘  └───────────────┘  │
└─────────────────────────────────────────────────────────┘

## Tech Stack

| Layer | Technology |
|-------|-----------|
| Frontend | Next.js 15, TypeScript, TailwindCSS, shadcn/ui |
| Backend | Go 1.22, Gin, Clean Architecture |
| Database | PostgreSQL 16, Redis 7 |
| Orchestration | Kubernetes (k3s → EKS) |
| GitOps | ArgoCD |
| CI/CD | GitHub Actions, Jenkins |
| IaC | Terraform, Ansible |
| Monitoring | Prometheus, Grafana, Loki, Tempo |
| Security | Trivy, Falco, OWASP ZAP, Sealed Secrets |

## Getting Started

### Prerequisites
- Docker & Docker Compose
- Make
- Git

### Local Development

```bash
# Clone the repository
git clone https://github.com/YOUR_USERNAME/sameer-devops-platform.git
cd sameer-devops-platform

# Copy environment variables
cp .env.example .env
# Edit .env with your values

# Start all services
make up

# View logs
make logs
```

Access the platform:
- Frontend: http://localhost:3000
- Backend API: http://localhost:8080
- Grafana: http://localhost:3001
- pgAdmin: http://localhost:5050

## Documentation

- [Architecture Overview](docs/architecture/architecture.md)
- [Local Development Runbook](docs/runbooks/local-dev.md)
- [ADR-001: Monorepo Structure](docs/adr/ADR-001-monorepo-structure.md)
- [Roadmap](roadmap.md)

## Project Status

🟡 **Phase 1 — Foundation** — In Progress

See [roadmap.md](roadmap.md) for full timeline.

## About the Engineer

**Sameer Malik RP** — Transitioning DevOps Engineer  
RHCSA Certified | AWS | Docker | Kubernetes | Terraform  

[LinkedIn](#) | [Portfolio](#)
