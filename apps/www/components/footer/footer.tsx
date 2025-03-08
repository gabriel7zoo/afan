import { useEffect } from 'react';
import { twMerge } from 'tailwind-merge';
import { Button } from './Button';
import { sections } from './Navigation';
import Link, { LinkProps } from 'next/link';
import { PropsWithChildren } from 'react';
import { BrandingMenu } from '../Layout/Public/BrandingMenu';

const Footer = ({ wide }: { wide?: boolean }) => {
  return (
    <div className={twMerge('flex w-full flex-col items-center space-y-24 px-4 py-16')}>
      <div
        className={twMerge(
          'dark:md:bg-polar-900 md:rounded-4xl flex w-full flex-col gap-x-32 gap-y-24 md:justify-between md:gap-y-12 md:bg-gray-50 md:p-16 lg:flex-row',
          wide ? 'max-w-7xl' : 'max-w-[970px]',
        )}
      >
        <div className="flex flex-grow flex-col gap-y-6">
          <span className="ml-2 text-black md:ml-0 dark:text-white">
            <BrandingMenu logoVariant="logotype" size={120} />
          </span>
          <span className="dark:text-polar-500 w-full flex-grow text-gray-500">
            &copy; Extrinsic Music Group. {new Date().getFullYear()}
          </span>
        </div>
        
        {/* Navigation Links */}
        <div className="flex flex-col gap-x-12 gap-y-12 text-sm md:flex-row [&>div]:w-36">
          {/* Platform */}
          <div className="flex flex-col gap-y-4">
            <h3 className="text-base dark:text-white">Platform</h3>
            <div className="flex flex-col gap-y-2">
              <FooterLink href="https://apply.extrinsicmusicgroup.com/">Apply</FooterLink>
              <FooterLink href="https://affiliates.extrinsicmusicgroup.com/">Affiliates</FooterLink>
              <FooterLink href="https://artists.extrinsicmusicgroup.com/">Artists</FooterLink>
              <FooterLink href="https://portal.extrinsicmusicgroup.com">Customer Portal</FooterLink>
            </div>
          </div>

          {/* Company */}
          <div className="flex flex-col gap-y-4">
            <h3 className="text-base dark:text-white">Company</h3>
            <div className="flex flex-col gap-y-2">
              <FooterLink href="/about">About</FooterLink>
              <FooterLink href="/blog">Blog</FooterLink>
              <FooterLink href="/changelog">Changelog</FooterLink>
              <FooterLink href="/careers">Careers</FooterLink>
              <FooterLink href="/docs">Docs</FooterLink>
              <FooterLink href="/glossary">Glossary</FooterLink>
              <FooterLink href="https://help.extrinsicmusicgroup.com/docs/paper/whitepaper">
                White Paper
              </FooterLink>
              <FooterLink href="https://extrinsicmusicgroup.com/assets/brand/emg_brand.zip">
                Brand Assets
              </FooterLink>
              <FooterLink href="https://help.extrinsicmusicgroup.com/docs/tos/tos">
                Terms of Service
              </FooterLink>
              <FooterLink href="https://help.extrinsicmusicgroup.com/docs/tos/privacy">
                Privacy Policy
              </FooterLink>
            </div>
          </div>

          {/* Community */}
          <div className="flex flex-col gap-y-4">
            <h3 className="text-lg dark:text-white">Community</h3>
            <div className="flex flex-col gap-y-2">
              <FooterLink href="https://youtube.com/extrinsicmusicgroup">YouTube</FooterLink>
              <FooterLink href="https://Instagram.com/extrinsicmusicgroup">Instagram</FooterLink>
              <FooterLink href="https://x.com/Extrinsicmusica">X / Twitter</FooterLink>
            </div>
          </div>

          {/* Support */}
          <div className="flex flex-col gap-y-4">
            <h3 className="text-base dark:text-white">Support</h3>
            <div className="flex flex-col gap-y-2">
              <FooterLink href="https://help.extrinsicmusicgroup.com/">Docs</FooterLink>
              <FooterLink href="mailto:support@extrinsicmusicgroup.com">Contact</FooterLink>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Footer;

const FooterLink = (props: PropsWithChildren<LinkProps>) => {
  return (
    <Link
      className="dark:text-polar-500 dark:hover:text-polar-50 flex flex-row items-center gap-x-1 text-gray-500 transition-colors hover:text-gray-500"
      {...props}
    >
      {props.children}
    </Link>
  );
};
