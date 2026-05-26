CREATE TABLE IF NOT EXISTS projects (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    slug        VARCHAR(100) UNIQUE NOT NULL,
    title       VARCHAR(255) NOT NULL,
    description TEXT,
    tech_stack  TEXT[] NOT NULL DEFAULT '{}',
    github_url  VARCHAR(500),
    live_url    VARCHAR(500),
    featured    BOOLEAN NOT NULL DEFAULT false,
    sort_order  INTEGER NOT NULL DEFAULT 0,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_projects_featured ON projects(featured);
CREATE INDEX idx_projects_sort ON projects(sort_order);
