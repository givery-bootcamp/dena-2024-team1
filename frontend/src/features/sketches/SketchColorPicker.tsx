import type { ChangeEventHandler } from "react";

import { useColorPicker } from "~/shared/hooks/useColorPicker";

export const SketchColorPicker = () => {
  const { setStrokeColor } = useColorPicker();

  const handleChange: ChangeEventHandler<HTMLInputElement> = (event) => {
    setStrokeColor(event.target.value);
  };

  return <input className="size-8 rounded p-0.5" type="color" onChange={handleChange} />;
};
