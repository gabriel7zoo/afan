import Image from "next/image";
import type { CSSProperties } from "react";
import { cn } from "@/lib/utils";

const logos = [
  {
    name: "EMG",
    url: "https://i.imghippo.com/files/NGlA1344HAI.png",
  },
  {
    name: "EMG",
    url: "https://i.imghippo.com/files/NGlA1344HAI.png",
  },
  {
    name: "EMG",
    url: "https://i.imghippo.com/files/NGlA1344HAI.png",
  },
];

export function DesktopLogoCloud() {
  return (
    <div className="hidden md:flex w-full flex-col items-center">
      <span className="font-mono text-sm md:text-md text-white/50 text-center opacity-0 animate-fade-in-up [animation-delay:1s]">
        Powering
      </span>

      <div className="flex w-full flex-col items-center justify-center px-4 md:px-8">
        <div className="mt-10 grid grid-cols-3 gap-x-6">
          {logos.map((logo, idx) => (
            <div
              key={String(idx)}
              className="relative w-[100px] h-[25px] opacity-0 animate-fade-in-up"
              style={{ "--animation-delay": `calc(1.1s + .12s * ${idx})` } as CSSProperties}
            >
              <Image src={logo.url} alt={logo.name} width={100} height={25} />
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}

export const MobileLogoCloud = () => {
  return (
    <div className="md:hidden w-full flex flex-col items-center">
      <span className="font-mono text-sm md:text-md text-white/50 text-center">Powering</span>

      <div className="w-full px-4 md:px-8">
        <div
          className="group relative mt-6 flex gap-6 overflow-hidden p-2"
          style={{
            maskImage: "linear-gradient(to left, transparent 0%, black 20%, black 80%, transparent 95%)",
          }}
        >
          {Array(5)
            .fill(null)
            .map((_, index) => (
              <div key={index} className="flex shrink-0 animate-logo-cloud flex-row justify-around gap-6">
                {logos.map((logo, key) => (
                  <div key={String(key)} className="relative w-[100px] h-[25px]">
                    <Image src={logo.url} alt={logo.name} width={100} height={25} />
                  </div>
                ))}
              </div>
            ))}
        </div>
      </div>
    </div>
  );
};
