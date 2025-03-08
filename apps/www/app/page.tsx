import { Hero } from "@/app/hero";
import { TopLeftShiningLight, TopRightShiningLight } from "@/components/svg/hero";
import Image from "next/image";
import Link from "next/link";
import mainboard from "../images/mainboard.svg";
import { DesktopLogoCloud, MobileLogoCloud } from "./(components)/logo-cloud-content";
import { PrimaryButton, SecondaryButton } from "@/components/button";
import { ChevronRight, LogIn } from "lucide-react";

export const metadata = {
  title: "Unkey",
  description: "Build better APIs faster",
  openGraph: {
    title: "Unkey",
    description: "Build better APIs faster",
    url: "https://unkey.com/",
    siteName: "unkey.com",
    images: [
      {
        url: "https://unkey.com/og.png",
        width: 1200,
        height: 675,
      },
    ],
  },
  twitter: {
    title: "Unkey",
    card: "summary_large_image",
  },
  icons: {
    shortcut: "/unkey.png",
  },
};

export default function Landing() {
  return (
    <>
      <TopRightShiningLight />
      <TopLeftShiningLight />
      <div className="relative w-full pt-6 overflow-hidden">
        <div className="container relative mx-auto">
          <Image
            src={mainboard}
            alt="Animated SVG showing computer circuits lighting up"
            className="absolute inset-x-0 flex xl:hidden -z-10 scale-[2]"
            priority
          />
        </div>
        <div className="container relative flex flex-col mx-auto space-y-16 md:space-y-32">
          {/* Hero Section */}
          <section>
            <Hero />
          </section>

          {/* Cloud Logos */}
          <section className="mt-16 md:mt-32">
            <DesktopLogoCloud />
            <MobileLogoCloud />
          </section>

          {/* Call-to-Action Buttons */}
          <section className="mt-16 md:mt-32 text-center">
            <div className="flex justify-center space-x-6">
              <Link href="https://app.unkey.com" className="group">
                <PrimaryButton shiny IconLeft={LogIn} label="Get Started" className="h-10" />
              </Link>
              <Link href="/docs">
                <SecondaryButton label="Visit the Docs" IconRight={ChevronRight} />
              </Link>
            </div>
          </section>

          {/* Footer */}
          <footer className="mt-16 md:mt-32 text-center text-gray-500">
            <p>© {new Date().getFullYear()} Unkey. All rights reserved.</p>
          </footer>
        </div>
      </div>
    </>
  );
}
