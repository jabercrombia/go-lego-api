'use client';

import { useEffect, useState } from 'react';

export default function HomePage() {
  const [legoSets, setLegoSets] = useState([]);

  useEffect(() => {
    const baseURL = process.env.NEXT_PUBLIC_API_BASE_URL;
    fetch(`${baseURL}/api/allsets`)
      .then((res) => res.json())
      .then((data) => setLegoSets(data))
      .catch((err) => console.error('Error fetching data:', err));
  }, []);

  return (
    <div className='container mx-auto'>
      <h1 className='text-3xl font-medium pt-[20px]'>All LEGO Sets</h1>
      <p className='mb-5'>This API provides a list of LEGO sets, including details like the set ID, name, theme, and release year.</p>
      <pre className="p-4 rounded overflow-x-auto bg-black">
        <code className="text-sm text-green-400">
          {JSON.stringify(legoSets, null, 2)}
        </code>
      </pre>
    </div>
  );
}
