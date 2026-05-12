"use client"

/*
 * WHY "use client" here?
 *
 * Next.js App Router renders components on the SERVER by default.
 * Server-side rendering is great for performance and SEO.
 * BUT — some things only exist in the BROWSER:
 *   - localStorage (where theme preference is saved)
 *   - window object
 *   - React hooks like useState, useEffect
 *
 * "use client" tells Next.js: "this component and everything it renders
 * must run in the browser, not on the server."
 *
 * next-themes uses localStorage and window.matchMedia (for system theme detection)
 * so it MUST be a client component.
 */

import * as React from "react"
import { ThemeProvider as NextThemesProvider } from "next-themes"

type ThemeProviderProps = React.ComponentProps<typeof NextThemesProvider>

export function ThemeProvider({ children, ...props }: ThemeProviderProps) {
  return <NextThemesProvider {...props}>{children}</NextThemesProvider>
}
