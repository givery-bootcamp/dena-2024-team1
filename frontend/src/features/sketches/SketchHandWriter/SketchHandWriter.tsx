import type { FC, RefObject } from "react";

import { useSketchHandWriter } from "./useSketchHandWriter";

type Props = {
  canvasRef: RefObject<HTMLCanvasElement>
}

export const SketchHandWriter: FC<Props> = ({ canvasRef }) => {
  const {
    canvasSetting,
    handleMouseDown,
    handleMouseMove,
    handleDrawingFinish,
  } = useSketchHandWriter({ canvasRef });

  return (
    <canvas
      className="border"
      ref={canvasRef}
      width={canvasSetting.width}
      height={canvasSetting.height}
      onMouseDown={handleMouseDown}
      onMouseMove={handleMouseMove}
      onMouseUp={handleDrawingFinish}
      onMouseLeave={handleDrawingFinish}
    />
  );
};
