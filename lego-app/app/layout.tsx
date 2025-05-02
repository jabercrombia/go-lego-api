import type { Metadata } from "next";
import "./globals.css";
import "../styles/headings.scss"
import Footer from "../components/footer"
import Header from "../components/header"
import { Lato } from 'next/font/google'

const lato = Lato({
  weight: ['100','300','400','700','900'],
  subsets: ['latin'],
})
 

export const metadata: Metadata = {
  title: "Lego Set API",
  description: "Created using Nextjs and Go",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body>
        <div className={` ${lato.className} min-h-screen flex flex-col`}>
          <Header/>
            <main className="flex-1">
              {children}  
            </main>
          <Footer/>
        </div>
      </body>
    </html>
  );
}
