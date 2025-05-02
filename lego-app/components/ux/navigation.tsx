// app/page.tsx or pages/index.tsx
'use client';

import "../../styles/nav.scss"
import * as React from "react"
import Link from "next/link"

const navigation = [
    {
        'name': 'home',
        'path': '/'
    },
    {
        'name': 'docs',
        'path': '/docs'
    },
    {
        'name': 'endpoints',
        'path': '/api/allsets'
    },
];



export default function HomePage() {
  return (
    <nav className="flex justify-stretch gap-20">
        <div className="justify-self-start">
        Lego Set API
        </div>
        <div className="flex gap-3">
            {
                navigation.map((elem, index) =>{
                    return (
                        <div key={index}>
                            <Link href={elem.path}>{elem.name}</Link>
                        </div>
                    )
                })
            }
        </div>
    </nav>
  );
}
