"use client";
import { PrimaryButton, SecondaryButton } from "@/components/button";
import { SectionTitle } from "@/components/section";
import type { LangIconProps } from "@/components/svg/lang-icons";
import {
  CurlIcon,
  ElixirIcon,
  GoIcon,
  JavaIcon,
  PythonIcon,
  RustIcon,
  TSIcon,
} from "@/components/svg/lang-icons";
import { CodeEditor } from "@/components/ui/code-editor";
import { CopyCodeSnippetButton } from "@/components/ui/copy-code-button";
import { MeteorLines } from "@/components/ui/meteorLines";
import { cn } from "@/lib/utils";
import * as TabsPrimitive from "@radix-ui/react-tabs";
import { ChevronRight } from "lucide-react";
import Link from "next/link";
import type { PrismTheme } from "prism-react-renderer";
import React, { useEffect, useState } from "react";
const Tabs = TabsPrimitive.Root;

const editorTheme = {
  plain: {
    color: "#F8F8F2",
    backgroundColor: "#282A36",
  },
  styles: [
    { types: ["keyword"], style: { color: "#9D72FF" } },
    { types: ["function"], style: { color: "#FB3186" } },
    { types: ["string"], style: { color: "#3CEEAE" } },
    { types: ["string-property"], style: { color: "#9D72FF" } },
    { types: ["number"], style: { color: "#FB3186" } },
    { types: ["comment"], style: { color: "#4D4D4D" } },
  ],
} satisfies PrismTheme;

type Framework = {
  name: string;
  Icon: React.FC<LangIconProps>;
  editorLanguage: string;
};

const languagesList = {
  Typescript: [
    { name: "Typescript", Icon: TSIcon, editorLanguage: "tsx" },
    { name: "Next.js", Icon: TSIcon, editorLanguage: "tsx" },
    { name: "Nuxt", Icon: TSIcon, editorLanguage: "tsx" },
    { name: "Hono", Icon: TSIcon, editorLanguage: "tsx" },
    { name: "Ratelimiting", Icon: TSIcon, editorLanguage: "tsx" },
  ],
  Python: [
    { name: "Python", Icon: PythonIcon, editorLanguage: "python" },
    { name: "FastAPI", Icon: PythonIcon, editorLanguage: "python" },
  ],
  Golang: [
    { name: "Verify key", Icon: GoIcon, editorLanguage: "go" },
    { name: "Create key", Icon: GoIcon, editorLanguage: "go" },
  ],
  Java: [
    { name: "Verify key", Icon: JavaIcon, editorLanguage: "tsx" },
    { name: "Create key", Icon: JavaIcon, editorLanguage: "tsx" },
  ],
  Elixir: [
    { name: "Verify key", Icon: ElixirIcon, editorLanguage: "tsx" },
  ],
  Rust: [
    { name: "Verify key", Icon: RustIcon, editorLanguage: "rust" },
  ],
  Curl: [
    { name: "Verify key", Icon: CurlIcon, editorLanguage: "tsx" },
    { name: "Create key", Icon: CurlIcon, editorLanguage: "tsx" },
    { name: "Ratelimit", Icon: CurlIcon, editorLanguage: "tsx" },
  ],
} as const;

type Props = { className?: string };
type Language = keyof typeof languagesList;
type LanguagesList = { name: Language; Icon: React.FC<LangIconProps> };

const languages = [
  { name: "Typescript", Icon: TSIcon },
  { name: "Python", Icon: PythonIcon },
  { name: "Rust", Icon: RustIcon },
  { name: "Golang", Icon: GoIcon },
  { name: "Curl", Icon: CurlIcon },
  { name: "Elixir", Icon: ElixirIcon },
  { name: "Java", Icon: JavaIcon },
] as LanguagesList[];

type FrameworkName = (typeof languagesList)[Language][number]["name"];

const LanguageTrigger = React.forwardRef<
  React.ElementRef<typeof TabsPrimitive.Trigger>,
  React.ComponentPropsWithoutRef<typeof TabsPrimitive.Trigger>
>(({ className, value, ...props }, ref) => (
  <TabsPrimitive.Trigger
    ref={ref}
    value={value}
    className={cn(
      "inline-flex items-center gap-1 justify-center whitespace-nowrap rounded-t-lg px-3 py-1.5 text-sm transition-all hover:text-white/80 disabled:pointer-events-none disabled:opacity-50 bg-gradient-to-t from-black to-black data-[state=active]:from-white/10 border border-b-0 text-white/30 data-[state=active]:text-white border-[#454545] font-light",
      className,
    )}
    {...props}
  />
));

LanguageTrigger.displayName = TabsPrimitive.Trigger.displayName;

export const CodeExamples: React.FC<Props> = ({ className }) => {
  const [language, setLanguage] = useState<Language>("Typescript");
  const [framework, setFramework] = useState<FrameworkName>("Typescript");
  const [languageHover, setLanguageHover] = useState("Typescript");

  function getLanguage({ language, framework }: { language: Language; framework: FrameworkName }) {
    return languagesList[language].find((f) => f.name === framework)?.editorLanguage || "tsx";
  }

  useEffect(() => {
    setFramework(languagesList[language][0].name);
  }, [language]);

  return (
    <section className={className}>
      <SectionTitle
        label="Code"
        title="Any language, any framework, always secure"
        text="Add authentication to your APIs in a few lines of code. We provide SDKs for a range of languages and frameworks, and an intuitive REST API with public OpenAPI spec."
        align="center"
        className="relative"
      >
        <div className="absolute bottom-32 left-[-50px]">
          <MeteorLines className="ml-2 fade-in-0" delay={3} number={1} />
          <MeteorLines className="ml-10 fade-in-40" delay={0} number={1} />
          <MeteorLines className="ml-16 fade-in-100" delay={5} number={1} />
        </div>
        <div className="absolute bottom-32 right-[200px]">
          <MeteorLines className="ml-2 fade-in-0" delay={4} number={1} />
          <MeteorLines className="ml-10 fade-in-40" delay={0} number={1} />
          <MeteorLines className="ml-16 fade-in-100" delay={2} number={1} />
        </div>
        <div className="mt-10">
          <div className="flex gap-6 pb-14">
            <Link key="get-started" href="https://app.unkey.com">
              <PrimaryButton shiny label="Get Started" IconRight={ChevronRight} />
            </Link>
            <Link key="docs" href="/docs">
              <SecondaryButton label="Visit the docs" IconRight={ChevronRight} />
            </Link>
          </div>
        </div>
      </SectionTitle>
      <div className="relative w-full rounded-4xl border-[.75px] border-white/10 bg-gradient-to-b from-[#111111] to-black border-t-[.75px] border-t-white/20">
        <div
          aria-hidden
          className="absolute pointer-events-none inset-x-16 h-[432px] bottom-[calc(100%-2rem)] bg-[radial-gradient(94.69%_94.69%_at_50%_100%,rgba(255,255,255,0.20)_0%,rgba(255,255,255,0)_55.45%)]"
        />
        <Tabs
          defaultValue={language}
          onValueChange={(l) => setLanguage(l as Language)}
          className="relative flex items-end h-16 px-4 border rounded-tr-3xl rounded-tl-3xl border-white/10 editor-top-gradient"
        >
          <TabsPrimitive.List className="flex items-end gap-4 overflow-x-auto scrollbar-hidden">
            {languages.map(({ name, Icon }) => (
              <LanguageTrigger key={name} value={name}>
                <Icon active={languageHover === name || language === name} />
                {name}
              </LanguageTrigger>
            ))}
          </TabsPrimitive.List>
        </Tabs>
      </div>
    </section>
  );
};
