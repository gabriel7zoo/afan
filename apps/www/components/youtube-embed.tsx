"use client";
import { BorderBeam } from "@/components/border-beam";
import { cn } from "@/lib/utils";
import FsLightbox from "fslightbox-react";
import { useState } from "react";
import { ImageWithBlur } from "./image-with-blur";

export function YoutubeEmbed({ className }: { className?: string }) {
  const [toggler, setToggler] = useState(false);

  return (
    <div className={cn("w-full h-full rounded-[38px] bg-white/5 group", className)}>
      <div
        aria-hidden
        className="rounded-[inherit] absolute inset-0 pointer-events-none -inset-2.5 border border-gray-800/40 bg-white/5"
      />

      <button type="button" className="relative w-full h-full" onClick={() => setToggler(!toggler)}>
        <div className="relative w-full h-full rounded-[inherit]">
          <ImageWithBlur
            className="rounded-[28px]"
            src="/apps/www/images/Screenshot_20250308_140811_YouTube.jpg"
            alt="Thumbnail for 'Arianna by Domy' YouTube video"
            quality={95}
            priority
            fill
          />
        </div>

        <div className="group absolute top-[calc(50%-80px/2)] duration-200 left-[calc(50%-112px/2)] bg-[#ffffff/30] h-[80px] w-[112px] bg-yt-button-gradient group-hover:bg-white transition-all rounded-[20px] flex items-center justify-center">
          <BorderBeam className="duration-200 group-hover:opacity-0" />
          <svg
            className="text-white duration-200 fill-current group-hover:text-black"
            xmlns="http://www.w3.org/2000/svg"
            width="48"
            height="48"
            viewBox="0 0 48 48"
          >
            <path d="M16 38V10L40 24L16 38Z" />
          </svg>
        </div>
      </button>

      <FsLightbox
        toggler={toggler}
        sources={[
          <div className="h-[600px] w-[1200px]">
            <iframe
              width="100%"
              height="100%"
              src="https://www.youtube.com/embed/7dzrKqm3Z2o?si=6V1y3ilhQw4q7v8f&autoplay=1"
              title="Arianna by Domy"
              allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share"
              allowFullScreen
            />
          </div>,
        ]}
      />
    </div>
  );
}
