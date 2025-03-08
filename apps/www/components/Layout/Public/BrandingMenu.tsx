'use client'

import LogoIcon from '@/components/Brand/LogoIcon'
import {
  ArrowDownwardOutlined,
  ContentPasteOutlined,
} from '@mui/icons-material'
import Link from 'next/link'

import LogoType from '@/components/Brand/LogoType'
import { useOutsideClick } from '@/utils/useOutsideClick'
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuTrigger,
} from '@polar-sh/ui/components/ui/dropdown-menu'
import { MouseEventHandler, useCallback, useRef, useState } from 'react'
import { twMerge } from 'tailwind-merge'

export const BrandingMenu = ({
  logoVariant = 'icon',
  size,
  className,
  logoClassName,
}: {
  logoVariant?: 'icon' | 'logotype'
  size?: number
  className?: string
  logoClassName?: string
}) => {
  const brandingMenuRef = useRef<HTMLDivElement>(null)

  useOutsideClick([brandingMenuRef], () => setBrandingMenuOpen(false))

  const [brandingMenuOpen, setBrandingMenuOpen] = useState(false)

  const handleTriggerClick: MouseEventHandler<HTMLButtonElement> = useCallback(
    (e) => {
      e.preventDefault()
      e.stopPropagation()
      setBrandingMenuOpen(true)
    },
    [],
  )

  const handleCopyLogoToClipboard = useCallback(() => {
    navigator.clipboard.writeText(
      logoVariant === 'icon' ? PolarIconLink : PolarLogoLink,
    )
    setBrandingMenuOpen(false)
  }, [logoVariant])

  return (
    <div className={twMerge('relative flex flex-row items-center', className)}>
      <DropdownMenu open={brandingMenuOpen}>
        <DropdownMenuTrigger onContextMenu={handleTriggerClick}>
          <Link href="/">
            {logoVariant === 'logotype' ? (
              <LogoType
                className={twMerge(
                  '-ml-2 text-black md:ml-0 dark:text-white',
                  logoClassName,
                )}
                width={size ?? 100}
              />
            ) : (
              <LogoIcon
                className={twMerge('text-black dark:text-white', logoClassName)}
                size={size ?? 42}
              />
            )}
          </Link>
        </DropdownMenuTrigger>
        <DropdownMenuContent ref={brandingMenuRef} align="start">
          <DropdownMenuLabel>Platform</DropdownMenuLabel>
          <DropdownMenuItem
            className="flex flex-row gap-x-3"
            onClick={handleCopyLogoToClipboard}
          >
            <ContentPasteOutlined fontSize="inherit" />
            <span>Copy Logo Link</span>
          </DropdownMenuItem>
          <DropdownMenuItem
            className="flex flex-row gap-x-3"
            onClick={() => setBrandingMenuOpen(false)}
          >
            <ArrowDownwardOutlined fontSize="inherit" />
            <Link href="/assets/brand/polar_brand.zip">
              Download Branding Assets
            </Link>
          </DropdownMenuItem>
        </DropdownMenuContent>
      </DropdownMenu>
    </div>
  )
}

// Replace SVG string with the image link
const PolarIconLink = "https://i.imghippo.com/files/NGlA1344HAI.png";
const PolarLogoLink = "https://i.imghippo.com/files/NGlA1344HAI.png";
