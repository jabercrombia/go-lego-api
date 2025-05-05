// app/page.tsx or pages/index.tsx
'use client';

import "../../styles/nav.scss"
import * as React from "react"
import Link from "next/link"

import {
    NavigationMenu,
    NavigationMenuContent,
    NavigationMenuItem,
    NavigationMenuLink,
    NavigationMenuList,
    NavigationMenuTrigger,
    navigationMenuTriggerStyle,
  } from "@/components/ui/navigation-menu"
  

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
    // <nav className="flex flex-row gap-2">
    //     <div className="flex items-center">
    //         <div className="h-[20px] w-[20px] bg-yellow-300 rounded-md"></div>
    //         <div className="w-[100px]">
    //             <Link href="/">
    //                 Lego Set API
    //             </Link>
    //         </div>
    //     </div>

    //     <div className="flex pl-20 gap-7">
    //         {
    //             navigation.map((elem, index) =>{
    //                 return (
    //                     <div key={index}>
    //                         <Link href={elem.path}>{elem.name}</Link>
    //                     </div>
    //                 )
    //             })
    //         }
    //     </div>
    // </nav>
    <NavigationMenu>
  <NavigationMenuList>
    <NavigationMenuItem>
      <NavigationMenuTrigger>Item One</NavigationMenuTrigger>
      <NavigationMenuContent>
        <NavigationMenuLink>Link</NavigationMenuLink>
      </NavigationMenuContent>
    </NavigationMenuItem>
    <Link href="/docs" legacyBehavior passHref>
    <NavigationMenuLink className={navigationMenuTriggerStyle()}>
      Documentation
    </NavigationMenuLink>
  </Link>
  </NavigationMenuList>
</NavigationMenu>

  );
}
