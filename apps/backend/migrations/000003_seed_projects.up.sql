INSERT INTO projects (slug, title, description, tech_stack, github_url, featured, sort_order)
VALUES
(
    'devops-platform',
    'Enterprise DevOps Portfolio Platform',
    'Production-grade SaaS platform built with Kubernetes, GitOps, CI/CD, and full observability stack.',
    ARRAY['Kubernetes', 'Go', 'Next.js', 'ArgoCD', 'Terraform', 'Prometheus'],
    'https://github.com/YOUR_USERNAME/sameer-devops-platform',
    true,
    1
),
(
    'kubernetes-gitops',
    'Kubernetes GitOps Pipeline',
    'ArgoCD-powered GitOps workflow with automated deployments, rollbacks, and drift detection.',
    ARRAY['Kubernetes', 'ArgoCD', 'Helm', 'GitHub Actions'],
    null,
    true,
    2
);
