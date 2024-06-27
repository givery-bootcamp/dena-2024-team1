import { useAtom } from "jotai";

import type { Sketch } from "~/generated";
import { SelectedSketchUrl } from "~/shared/store/Sketch";

type SketchItemProps = {
  sketch: Sketch
};

export const SketchItem = ({ sketch }: SketchItemProps) => {
  const [, setSelectedSketchUrl] = useAtom(SelectedSketchUrl);
  return (
    <div onClick={() => setSelectedSketchUrl(sketch.image_url)}>
      <img className="h-[200px] w-[300px] border border-gray-200 object-cover" src={sketch.image_url} alt={sketch.image_url} />
    </div>
  );
};
