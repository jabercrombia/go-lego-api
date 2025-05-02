// app/page.tsx or pages/index.tsx
'use client';

import Nav from "../components/ux/navigation"

export default function HomePage() {


  return (
    <div className='w-full bg-black text-white text-center py-4'>
      <h1>LEGO Sets</h1>
      <div className="container mx-auto">
        <Nav/>
      </div>
    </div>
  );
}
