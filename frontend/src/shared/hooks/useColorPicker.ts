import { atom, useAtom } from "jotai";

const strokeColorAtom = atom<string>("black");

export const useColorPicker = () => {
  const [strokeColor, setStrokeColor] = useAtom(strokeColorAtom);

  return { strokeColor, setStrokeColor };
};
