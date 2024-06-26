import type { Sketch } from "~/generated";

type SketchItemProps = {
  sketch: Sketch
};

export const SketchItem = ({ sketch }: SketchItemProps) => {
  return(
    <img className="h-[200px] w-[300px] border border-gray-200 object-cover" src={sketch.imageUrl} alt={sketch.imageUrl} />
  );
};
