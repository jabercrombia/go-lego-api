import type { Metadata } from "next";
import "./globals.css";
import "../styles/headings.scss"
import Footer from "../components/footer"
import { GoogleAnalytics } from '@next/third-parties/google'

import Header from "../components/header"
import { Lato } from 'next/font/google'

const lato = Lato({
  weight: ['100','300','400','700','900'],
  subsets: ['latin'],
})
 

export const metadata: Metadata = {
  title: `Lego Set API`,
  description: `Discover a powerful LEGO API that provides detailed information about LEGO sets, minifigures, parts, themes, and more. Perfect for developers, hobbyists, and collectors looking to build LEGO-powered applications with ease and accuracy.`,
  keywords :[
    "LEGO API",
    "LEGO set data",
    "LEGO minifigures API",
    "LEGO parts database",
    "LEGO themes",
    "LEGO for developers",
    "LEGO collection API",
    "LEGO JSON API",
    "LEGO programming",
    "LEGO data integration",
    "LEGO hobbyist tools",
    "LEGO collectors app"
  ],
  authors: [{ name: 'Justin Abercrombia', url: 'http://www.github.com/jabercrombia' }],
  creator: 'Justin Abercrombia',
  openGraph: {
    images: '/homepage.png',
  },
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body className={lato.className}>
        <div className="min-h-screen flex flex-col">
          <Header/>
            <main className="flex-1">
              {children}  
            </main>
          <Footer/>
        </div>
      </body>
      <GoogleAnalytics gaId={process.env.GOOGLE_TRACKIND_ID || ''} />
    </html>
  );
}
