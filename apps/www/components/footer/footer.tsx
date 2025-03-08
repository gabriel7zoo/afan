"use client";
import { cn } from "@/lib/utils";
import Link from "next/link";
import Image from "next/image"; // Import Image for logos

// Extrinsic Music Group Logo (Replace with actual hosted logo if needed)
const ExtrinsicLogo = () => (
  <svg width="150" height="50" viewBox="0 0 150 50" xmlns="http://www.w3.org/2000/svg">
    <text x="10" y="35" fontSize="24" fill="white" fontFamily="Arial, sans-serif">
    </text>
  </svg>
);

// Proudly European Icon (SVG)
const EuropeanIcon = () => (
  <div className="flex flex-col items-center mt-4">
    <svg width="80" height="80" viewBox="0 0 100 100" xmlns="http://www.w3.org/2000/svg">
      <rect width="100" height="100" rx="15" fill="#003399" />
      <g fill="#FFD700">
        {[...Array(12)].map((_, i) => {
          const angle = (i * 30) * (Math.PI / 180);
          const x = 50 + 30 * Math.cos(angle);
          const y = 50 + 30 * Math.sin(angle);
          return <circle key={i} cx={x} cy={y} r="4" />;
        })}
      </g>
    </svg>
    <span className="text-white mt-2 text-sm font-medium">Proudly European</span>
  </div>
);

type NavLink = {
  title: string;
  href: string;
  external?: boolean;
};

const navigation = [
  {
    title: "Company",
    links: [
      { title: "About", href: "/about" },
      { title: "Blog", href: "/blog" },
      { title: "Careers", href: "/careers" },
      {
        title: "Analytics",
        href: "https://us.posthog.com/shared/HwZNjaKOLtgtpj6djuSo3fgOqrQm0Q?whitelabel",
        external: true,
      },
      {
        title: "Docs",
        href: "/docs",
        external: true,
      },
      {
        title: "Glossary",
        href: "/glossary",
      },
    ],
  },
  {
    title: "Connect",
    links: [
      { title: "X (Twitter)", href: "https://x.com/extrinsicmusicgroup", external: true },
      {
        title: "Book a Call",
        href: "https://cal.com/team/extrinsic/user-interview?utm_source=banner&utm_campaign=oss",
        external: true,
      },
    ],
  },
  {
    title: "Legal",
    links: [
      { title: "Terms of Service", href: "/policies/terms" },
      { title: "Privacy Policy", href: "/policies/privacy" },
    ],
  },
] satisfies Array<{ title: string; links: Array<NavLink> }>;

const Column: React.FC<{ title: string; links: Array<NavLink>; className?: string }> = ({
  title,
  links,
  className,
}) => {
  return (
    <div className={cn("flex flex-col gap-8 text-left", className)}>
      <span className="w-full text-sm font-medium tracking-wider text-white font-display">
        {title}
      </span>
      <ul className="flex flex-col gap-4 md:gap-6">
        {links.map((link) => (
          <li key={link.href}>
            <Link
              href={link.href}
              target={link.external ? "_blank" : undefined}
              rel={link.external ? "noopener noreferrer" : undefined}
              className="text-sm font-normal transition hover:text-white/40 text-white/70"
            >
              {link.title}
            </Link>
          </li>
        ))}
      </ul>
    </div>
  );
};

export function Footer() {
  return (
    <div className="border-t border-white/20 blog-footer-radial-gradient">
      <footer className="container relative grid grid-cols-2 gap-8 pt-8 mx-auto overflow-hidden lg:gap-16 sm:grid-cols-3 xl:grid-cols-5 sm:pt-12 md:pt-16 lg:pt-24 xl:pt-32">
        <div className="flex flex-col items-center col-span-2 sm:items-start sm:col-span-3 xl:col-span-2">
          <ExtrinsicLogo />
          <div className="mt-8 text-sm font-normal leading-6 text-white/60">
            NextGen Label.
          </div>
          <div className="text-sm font-normal leading-6 text-white/40">
            Extrinsic Music Group, Inc. {new Date().getUTCFullYear()}
          </div>
          <EuropeanIcon /> {/* Adds Proudly European icon */}
        </div>

        {navigation.map(({ title, links }) => (
          <Column key={title} title={title} links={links} className="col-span-1" />
        ))}
      </footer>
      <div className="container mt-8 h-[100px]"></div>
    </div>
  );
}
