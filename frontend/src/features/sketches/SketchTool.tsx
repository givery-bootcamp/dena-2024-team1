import type { FC, ReactNode } from "react";

import { SketchColorPicker } from "./SketchColorPicker";

import { useStrokeSetting } from "~/shared/hooks/useStrokeSetting";
import { Slider } from "~/shared/components/Slider";

export const SketchTool = () => {
  const { strokeWidth, setStrokeWidth } = useStrokeSetting();

  return (
    <div className="flex w-56 flex-col gap-8">
      <h5 className="text-lg font-bold">お絵描きツール</h5>
      <div className="flex flex-col gap-6">
        <Section>
          <SectionTitle>線の色</SectionTitle>
          <SketchColorPicker />
        </Section>
        <Section>
          <SectionTitle>線の太さ</SectionTitle>
          <div className="flex w-full items-center justify-center gap-2">
            <Slider
              min={1}
              max={10}
              value={[strokeWidth]}
              onValueChange={(value) => setStrokeWidth(value[0])}
            />
            {/* レイアウト崩れを防ぐためにw-1を指定する */}
            <p className="w-1 text-gray-200">{strokeWidth}</p>
          </div>
        </Section>
      </div>
    </div>
  );
};

const Section: FC<{children: ReactNode}> = ({ children }) => {
  return (
    <div className="flex flex-col items-center gap-2">{children}</div>
  );
};
const SectionTitle: FC<{children: ReactNode}> = ({ children }) => {
  return (
    <p className="w-full text-sm font-semibold text-gray-200">{children}</p>
  );
};
