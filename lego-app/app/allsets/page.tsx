// app/page.tsx or pages/index.tsx
'use client';

import { useEffect, useState } from 'react';

export default function HomePage() {
  const [legoSets, setLegoSets] = useState([]);

  useEffect(() => {
    // Replace with the appropriate URL if Go backend is deployed
    fetch('http://localhost:8080/api/allsets')
      .then((res) => res.json())
      .then((data) => setLegoSets(data))
      .catch((err) => console.error('Error fetching data:', err));
  }, []);

  console.log(legoSets);

  return (
    <div className='container mx-auto'>
      <h1>LEGO Sets</h1>
      <pre className="p-4 rounded overflow-x-auto">
        <code className="text-sm text-gray-800">
          {JSON.stringify(legoSets, null, 2)}
        </code>
      </pre>

    </div>
  );
}
