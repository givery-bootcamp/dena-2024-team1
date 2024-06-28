import { atom, useAtom } from "jotai";

const strokeColorAtom = atom<"black" | "red">("black");

export const useColorPicker = () => {
  const [strokeColor, setStrokeColor] = useAtom(strokeColorAtom);

  const toggleStrokeColor = () => {
    setStrokeColor((prev) => (prev === "black" ? "red" : "black"));
  };

  return { strokeColor, toggleStrokeColor };
};
