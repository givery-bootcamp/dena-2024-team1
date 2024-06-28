import { atom, useAtom } from "jotai";

const strokeColorAtom = atom<string>("black");
const strokeWidthAtom = atom<number>(1);

export const useColorPicker = () => {
  const [strokeColor, setStrokeColor] = useAtom(strokeColorAtom);
  const [strokeWidth, setStrokeWidth] = useAtom(strokeWidthAtom);

  return { strokeColor, strokeWidth, setStrokeColor, setStrokeWidth };
};
