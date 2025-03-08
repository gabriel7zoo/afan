'use client';

import { useOutsideClick } from '@/utils/useOutsideClick';
import { MouseEventHandler, useCallback, useRef, useState } from 'react';
import Link from 'next/link';
import { twMerge } from 'tailwind-merge';
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuTrigger,
} from '@polar-sh/ui/components/ui/dropdown-menu';
import { ContentPasteOutlined, ArrowDownwardOutlined } from '@mui/icons-material';

// Placeholder for logo components, update with actual ones
const LogoIcon = ({ size, className }: { size?: number; className?: string }) => (
  <img src="/logo-icon.png" alt="Logo Icon" className={className} style={{ width: size ?? 42 }} />
);

const LogoType = ({ width, className }: { width?: number; className?: string }) => (
  <img src="/logo-type.png" alt="Logo Type" className={className} style={{ width: width ?? 100 }} />
);

export const BrandingMenu = ({
  logoVariant = 'icon',
  size,
  className,
  logoClassName,
}: {
  logoVariant?: 'icon' | 'logotype';
  size?: number;
  className?: string;
  logoClassName?: string;
}) => {
  const brandingMenuRef = useRef<HTMLDivElement>(null);
  const [brandingMenuOpen, setBrandingMenuOpen] = useState(false);

  useOutsideClick([brandingMenuRef], () => setBrandingMenuOpen(false));

  const handleTriggerClick: MouseEventHandler<HTMLButtonElement> = useCallback((e) => {
    e.preventDefault();
    e.stopPropagation();
    setBrandingMenuOpen(true);
  }, []);

  const handleCopyLogoToClipboard = useCallback(() => {
    navigator.clipboard.writeText(logoVariant === 'icon' ? LogoIconLink : LogoTypeLink);
    setBrandingMenuOpen(false);
  }, [logoVariant]);

  return (
    <div className={twMerge('relative flex flex-row items-center', className)}>
      <DropdownMenu open={brandingMenuOpen}>
        <DropdownMenuTrigger onContextMenu={handleTriggerClick}>
          <Link href="/">
            {logoVariant === 'logotype' ? (
              <LogoType className={twMerge('text-black dark:text-white', logoClassName)} width={size ?? 100} />
            ) : (
              <LogoIcon className={twMerge('text-black dark:text-white', logoClassName)} size={size ?? 42} />
            )}
          </Link>
        </DropdownMenuTrigger>
        <DropdownMenuContent ref={brandingMenuRef} align="start">
          <DropdownMenuLabel>Platform</DropdownMenuLabel>
          <DropdownMenuItem className="flex flex-row gap-x-3" onClick={handleCopyLogoToClipboard}>
            <ContentPasteOutlined fontSize="inherit" />
            <span>Copy Logo Link</span>
          </DropdownMenuItem>
          <DropdownMenuItem className="flex flex-row gap-x-3" onClick={() => setBrandingMenuOpen(false)}>
            <ArrowDownwardOutlined fontSize="inherit" />
            <Link href="/assets/brand/branding_assets.zip">Download Branding Assets</Link>
          </DropdownMenuItem>
        </DropdownMenuContent>
      </DropdownMenu>
    </div>
  );
};

// Replace with actual hosted image URLs
const LogoIconLink = 'https://i.imghippo.com/files/NGlA1344HAI.png';
const LogoTypeLink = 'https://i.imghippo.com/files/NGlA1344HAI.png';
