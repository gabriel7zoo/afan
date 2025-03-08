"use client";
import {
  Drawer,
  DrawerContent,
  DrawerFooter,
  DrawerHeader,
  DrawerTrigger,
} from "@/components/ui/drawer";
import { cn } from "@/lib/utils";
import { motion } from "framer-motion";
import { ChevronDown, ChevronRight } from "lucide-react";
import Link from "next/link";
import { useEffect, useState } from "react";
import { PrimaryButton } from "../button";
import { DesktopNavLink, MobileNavLink } from "./link";

export function Navigation() {
  const [scrollPercent, setScrollPercent] = useState(0);

  useEffect(() => {
    const handleScroll = () => {
      const scrollThreshold = 100;
      setScrollPercent(Math.min(window.scrollY / 2 / scrollThreshold, 1));
    };

    handleScroll();
    window.addEventListener("scroll", handleScroll);
    return () => window.removeEventListener("scroll", handleScroll);
  }, []);

  return (
    <motion.nav
      style={{
        backgroundColor: `rgba(0, 0, 0, ${scrollPercent * 0.8})`,
        borderColor: `rgba(255, 255, 255, ${Math.min(scrollPercent / 5, 0.1)})`,
        backdropFilter: "blur(10px)",
      }}
      className="fixed z-[100] top-0 border-b-[.5px] border-white/10 w-full py-3 transition-all"
    >
      <div className="container flex items-center justify-between">
        <div className="flex items-center justify-between w-full sm:w-auto sm:gap-12 lg:gap-20">
          <Link href="/" aria-label="Home">
            <Logo />
          </Link>
          <MobileLinks className="lg:hidden" />
          <DesktopLinks className="hidden lg:flex" />
        </div>
        <div className="hidden sm:flex">
          <Link href="https://extrinsicmusicgroup.com">
            <PrimaryButton
              label="EMG"
              IconRight={ChevronRight}
              className="h-9 px-5 text-sm rounded-full bg-white/10 hover:bg-white/20 transition-all"
            />
          </Link>
        </div>
      </div>
    </motion.nav>
  );
}

function MobileLinks({ className }: { className?: string }) {
  const [isOpen, setIsOpen] = useState(false);
  return (
    <div className={className}>
      <Drawer open={isOpen} onOpenChange={setIsOpen}>
        <DrawerTrigger asChild>
          <button
            type="button"
            onClick={() => setIsOpen(true)}
            className="flex items-center justify-end h-9 px-4 text-sm text-white/70 hover:text-white/90 rounded-full bg-white/10 hover:bg-white/20 transition-all"
          >
            Menu
            <ChevronDown className="w-4 h-4 ml-2" />
          </button>
        </DrawerTrigger>
        <DrawerContent className="bg-black/80 backdrop-blur-lg rounded-lg p-4 z-[110]">
          <DrawerHeader className="flex justify-center">
            <Logo />
          </DrawerHeader>
          <div className="relative w-full mx-auto antialiased z-[110]">
            <ul className="flex flex-col px-8 divide-y divide-white/15">
              <li>
                <MobileNavLink onClick={() => setIsOpen(false)} href="/" label="Home" />
              </li>
              <li>
                <MobileNavLink onClick={() => setIsOpen(false)} href="/about" label="About" />
              </li>
              <li>
                <MobileNavLink onClick={() => setIsOpen(false)} href="/docs" label="Docs" />
              </li>
              <li>
                <MobileNavLink
                  onClick={() => setIsOpen(false)}
                  href="https://extrinsicmusicgroup.com"
                  label="EMG"
                />
              </li>
            </ul>
          </div>
          <DrawerFooter>
            <button
              type="button"
              onClick={() => setIsOpen(false)}
              className={cn(
                "px-5 py-2 text-white/80 hover:text-white transition-all bg-white/10 hover:bg-white/20 rounded-full",
                className,
              )}
            >
              Close
            </button>
          </DrawerFooter>
        </DrawerContent>
      </Drawer>
    </div>
  );
}

const DesktopLinks: React.FC<{ className: string }> = ({ className }) => (
  <ul className={cn("items-center hidden gap-8 lg:flex xl:gap-12", className)}>
    <li>
      <DesktopNavLink href="/" label="Home" />
    </li>
    <li>
      <DesktopNavLink href="/about" label="About" />
    </li>
    <li>
      <DesktopNavLink href="/docs" label="Docs" />
    </li>
    <li>
      <DesktopNavLink href="https://extrinsicmusicgroup.com" label="EMG" external />
    </li>
  </ul>
);

const Logo: React.FC<{ className?: string }> = ({ className }) => (
  <img
    src="https://i.imghippo.com/files/NGlA1344HAI.png"
    alt="Logo"
    className={cn("h-10 w-auto rounded-full", className)}
  />
);
