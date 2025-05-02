import Link from "next/link"


export default function Home() {
  return (
    <div className="grid grid-cols-2 container mx-auto pt-10">
      <div>
          <h1 className="text-5xl">The Ultimate LEGO Set API for Developers</h1>
          <p className="text-xl font-light pt-4">Welcome to my Lego sets API, your gateway to discovering the world of LEGO® sets. This project leverages Next.js for a fast, responsive frontend and Go (Golang) for a powerful backend API, 
            enabling efficient data handling and seamless user experience across pages.</p>
            <Link href="/docs/">
            <button className="primary">View Documentation</button>
            </Link>
      </div>
      <div>
      </div>
    </div>
  );
}
