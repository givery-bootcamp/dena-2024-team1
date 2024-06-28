import type { FC } from "react";
import classNames from "classnames";

import { SketchItem } from "./SketchItem";

import type { Sketch } from "~/generated";

type Props = {
  sketches: Sketch[]
}

export const SketchList: FC<Props> = ({ sketches }) => {
  const sketchCount = sketches.length;

  /**
   * スケッチ総数に対して、レイアウトが正方形に近くなるようグリッド数を計算する
   * Tailwindはデフォルトで12列のグリッドを提供しているため、12列を超える場合は12列に収める
   */
  const gridCols = Math.min(12, Math.ceil(Math.sqrt(sketchCount)));

  const colClassName = gridColClassNames[gridCols];

  return (
    <div className={classNames("grid gap-10", colClassName)}>
      {sketches.map((sketch) => (
        <button key={sketch.id} className="transition hover:scale-105">
          <SketchItem sketch={sketch}/>
        </button>
      ))}
    </div>
  );
};

const gridColClassNames = {
  1: "grid-cols-1",
  2: "grid-cols-2",
  3: "grid-cols-3",
  4: "grid-cols-4",
  5: "grid-cols-5",
  6: "grid-cols-6",
  7: "grid-cols-7",
  8: "grid-cols-8",
  9: "grid-cols-9",
  10: "grid-cols-10",
  11: "grid-cols-11",
  12: "grid-cols-12",
} as { [key: number]: string };
