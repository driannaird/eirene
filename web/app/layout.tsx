import type { Metadata } from "next";
import { JetBrains_Mono } from "next/font/google";
import "./globals.css";

const jet = JetBrains_Mono({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "Eirene",
  description: "Tools base website for automate install package on your OS, manage docker, and storage server",
  authors: [
    {
      name: "Pradana",
      url: "https://github.com/rulanugrh/eirene"
    }
  ]
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body className={jet.className}>{children}</body>
    </html>
  );
}
