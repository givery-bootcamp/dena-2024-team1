import type { FC, RefObject } from "react";

import { useSketchHandWriter } from "./useSketchHandWriter";

type Props = {
  canvasRef: RefObject<HTMLCanvasElement>
  onCanvasUpdate: (canvas: HTMLCanvasElement) => void;
}

export const SketchHandWriter: FC<Props> = ({ canvasRef, onCanvasUpdate }) => {
  const {
    canvasSetting,
    handleMouseDown,
    handleMouseMove,
    handleDrawingFinish,
  } = useSketchHandWriter({
    canvasRef,
    onCanvasUpdate,
  });

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
