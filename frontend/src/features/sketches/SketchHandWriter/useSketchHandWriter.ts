import {
  useState,
  useEffect,
  type RefObject,
  type MouseEventHandler,
} from "react";

import { useCanvas } from "~/shared/hooks/useCanvas";

const canvasSetting = {
  width: 900,
  height: 600,
};

type Args = {
  canvasRef: RefObject<HTMLCanvasElement>
}

export const useSketchHandWriter = ({ canvasRef }: Args) => {
  const [isDrawing, setIsDrawing] = useState<boolean>(false);
  const { getContext, clearCanvas } = useCanvas({ canvasRef });

  useEffect(() => {
    clearCanvas();
  }, []);

  const handleMouseDown: MouseEventHandler = (event) =>  {
    setIsDrawing(true);

    const { offsetX: x, offsetY: y } = event.nativeEvent;

    const ctx = getContext();
    ctx.beginPath();
    ctx.moveTo(x, y);
  };

  const handleMouseMove: MouseEventHandler = (event) => {
    if (!isDrawing) return;

    const { offsetX: x ,offsetY: y } = event.nativeEvent;

    const ctx = getContext();
    ctx.lineTo(x, y);
    ctx.stroke();
  }; 

  const handleDrawingFinish = () => {
    setIsDrawing(false);
  };

  return {
    canvasSetting,
    handleMouseDown,
    handleMouseMove,
    handleDrawingFinish,
  };
};
