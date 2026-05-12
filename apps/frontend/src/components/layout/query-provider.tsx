"use client"

/*
 * TanStack Query manages ALL server state in your app.
 *
 * "Server state" = data that lives on your backend/database.
 *   e.g. your projects list, blog posts, visitor counter
 *
 * "Client state" = data that lives only in the browser.
 *   e.g. is the mobile menu open? what tab is selected?
 *
 * Without TanStack Query, every component that needs API data must:
 *   1. Track loading state manually (isLoading)
 *   2. Track error state manually (error)
 *   3. Store the data manually (useState)
 *   4. Handle refetching manually
 *   5. Invalidate cache manually when data changes
 *
 * TanStack Query handles ALL of that automatically.
 * One hook call: const { data, isLoading, error } = useQuery(...)
 */

import { QueryClient, QueryClientProvider } from "@tanstack/react-query"
import { useState } from "react"

export function QueryProvider({ children }: { children: React.ReactNode }) {
  const [queryClient] = useState(
    () =>
      new QueryClient({
        defaultOptions: {
          queries: {
            staleTime: 60 * 1000,      // Data is "fresh" for 60 seconds
            gcTime: 5 * 60 * 1000,     // Unused cache lives for 5 minutes
            retry: 1,                   // Retry failed requests once
            refetchOnWindowFocus: false, // Don't refetch when user tabs back
          },
        },
      })
  )

  return (
    <QueryClientProvider client={queryClient}>{children}</QueryClientProvider>
  )
}
