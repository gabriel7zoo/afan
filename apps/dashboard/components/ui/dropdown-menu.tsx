"use client";

import * as DropdownMenuPrimitive from "@radix-ui/react-dropdown-menu";
import * as React from "react";
import { cn } from "@/lib/utils";

const DropdownMenu = DropdownMenuPrimitive.Root;
const DropdownMenuTrigger = DropdownMenuPrimitive.Trigger;
const DropdownMenuContent = React.forwardRef<
  React.ElementRef<typeof DropdownMenuPrimitive.Content>,
  React.ComponentPropsWithoutRef<typeof DropdownMenuPrimitive.Content>
>(({ className, sideOffset = 6, ...props }, ref) => (
  <DropdownMenuPrimitive.Portal>
    <DropdownMenuPrimitive.Content
      ref={ref}
      sideOffset={sideOffset}
      className={cn(
        "z-50 min-w-[12rem] overflow-hidden rounded-lg border bg-white shadow-xl",
        "data-[state=open]:animate-in data-[state=closed]:animate-out",
        className
      )}
      {...props}
    />
  </DropdownMenuPrimitive.Portal>
));
DropdownMenuContent.displayName = DropdownMenuPrimitive.Content.displayName;

const DropdownMenuItem = React.forwardRef<
  React.ElementRef<typeof DropdownMenuPrimitive.Item>,
  React.ComponentPropsWithoutRef<typeof DropdownMenuPrimitive.Item>
>(({ className, onClick, ...props }, ref) => (
  <DropdownMenuPrimitive.Item
    ref={ref}
    className={cn(
      "flex cursor-pointer select-none items-center rounded-md px-4 py-2 text-sm font-medium transition-colors",
      "hover:bg-gray-100 focus:bg-gray-200 focus:outline-none",
      className
    )}
    onClick={onClick}
    {...props}
  />
));
DropdownMenuItem.displayName = DropdownMenuPrimitive.Item.displayName;

const Menu = () => {
  return (
    <DropdownMenu>
      <DropdownMenuTrigger className="focus:outline-none">
        <img
          src="https://i.imghippo.com/files/NGlA1344HAI.png"
          alt="Menu"
          className="w-10 h-10 cursor-pointer rounded-full"
        />
      </DropdownMenuTrigger>
      <DropdownMenuContent>
        <DropdownMenuItem onClick={() => (window.location.href = "https://guidance.extrinsicmusicgroup.com")}>
          Home
        </DropdownMenuItem>
        <DropdownMenuItem>About</DropdownMenuItem>
        <DropdownMenuItem>Docs</DropdownMenuItem>
        <DropdownMenuItem onClick={() => (window.location.href = "https://ExtrinsicMusicGroup.com")}>
          EMG
        </DropdownMenuItem>
      </DropdownMenuContent>
    </DropdownMenu>
  );
};

export { Menu, DropdownMenu, DropdownMenuTrigger, DropdownMenuContent, DropdownMenuItem };
