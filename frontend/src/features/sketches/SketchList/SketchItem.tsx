import { useSetAtom } from "jotai";

import type { Sketch } from "~/generated";
import { selectedSketchAtom } from "~/shared/store/Sketch";

type SketchItemProps = {
  sketch: Sketch
};

export const SketchItem = ({ sketch }: SketchItemProps) => {
  const setSelectedSketch = useSetAtom(selectedSketchAtom);
  return (
    <div onClick={() => setSelectedSketch(sketch)}>
      <img className="h-[200px] w-[300px] border border-gray-200 object-cover" src={sketch.image_url} alt={sketch.image_url} />
    </div>
  );
};
