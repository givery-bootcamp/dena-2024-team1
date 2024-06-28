import type { ChangeEventHandler, FC, ReactNode } from "react";

import { SketchColorPicker } from "./SketchColorPicker";

import { Input } from "~/shared/components/Input";
import { useStrokeSetting } from "~/shared/hooks/useStrokeSetting";

export const SketchTool = () => {
  const { strokeWidth, setStrokeWidth } = useStrokeSetting();

  const handleStorkeWidthChange: ChangeEventHandler<HTMLInputElement> = (event) => {
    // 太さは0未満にはできない
    const value = Math.min(0, Number(event.target.value));

    setStrokeWidth(value);
  };

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
          <Input
            className="w-full"
            type="number"
            value={strokeWidth}
            onChange={handleStorkeWidthChange}
          />
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
