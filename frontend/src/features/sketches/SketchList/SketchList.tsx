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
        <button key={sketch.id} className="transition hover:scale-105">
          <SketchItem sketch={sketch}/>
        </button>
      ))}
    </div>
  );
};
