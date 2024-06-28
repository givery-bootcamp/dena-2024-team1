import type { ChangeEventHandler } from "react";

import { useStrokeSetting } from "~/shared/hooks/useStrokeSetting";

export const SketchColorPicker = () => {
  const { strokeColor, setStrokeColor } = useStrokeSetting();

  const handleChange: ChangeEventHandler<HTMLInputElement> = (event) => {
    setStrokeColor(event.target.value);
  };

  return <input className="size-8 rounded p-0.5" type="color" value={strokeColor} onChange={handleChange} />;
};
