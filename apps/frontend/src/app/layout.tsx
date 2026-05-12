import type { Metadata } from "next"
import { Geist, Geist_Mono } from "next/font/google"
import { ThemeProvider } from "@/components/layout/theme-provider"
import { QueryProvider } from "@/components/layout/query-provider"
import { Navbar } from "@/components/layout/navbar"
import "./globals.css"

const geistSans = Geist({
  variable: "--font-geist-sans",
  subsets: ["latin"],
})

const geistMono = Geist_Mono({
  variable: "--font-geist-mono",
  subsets: ["latin"],
})

export const metadata: Metadata = {
  title: {
    default: "Sameer Malik RP — DevOps Engineer",
    template: "%s | Sameer DevOps Platform",
  },
  description: "Enterprise-grade DevOps Portfolio Platform. Kubernetes, GitOps, CI/CD, Observability, Security Engineering.",
  keywords: ["DevOps", "Kubernetes", "GitOps", "ArgoCD", "Terraform", "RHCSA", "Platform Engineering", "SRE"],
  authors: [{ name: "Sameer Malik RP" }],
  robots: { index: true, follow: true },
}

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode
}>) {
  return (
    <html lang="en" suppressHydrationWarning>
      <body className={`${geistSans.variable} ${geistMono.variable} font-sans antialiased min-h-screen`}>
        <ThemeProvider attribute="class" defaultTheme="dark" enableSystem disableTransitionOnChange>
          <QueryProvider>
            <Navbar />
            <main className="relative">{children}</main>
          </QueryProvider>
        </ThemeProvider>
      </body>
    </html>
  )
}
