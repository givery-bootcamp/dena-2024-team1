import type { FC } from "react";

import { SketchItem } from "./SketchItem";

import type { Sketch } from "~/generated";

type Props = {
  sketches: Sketch[]
}

export const SketchList: FC<Props> = ({ sketches }) => {
  return (
    <div className='grid grid-cols-3 gap-10'>
      {sketches.map((sketch) => (
        <SketchItem key={sketch.id} sketch={sketch}/>
      ))}
    </div>
  );
};
