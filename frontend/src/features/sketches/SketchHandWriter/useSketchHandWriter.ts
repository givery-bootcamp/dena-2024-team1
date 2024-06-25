import { useState, useEffect, type RefObject } from "react";

const canvasSetting = {
  width: 500,
  height: 300,
};

type Args = {
  canvasRef: RefObject<HTMLCanvasElement>
  onCanvasUpdate: (canvas: HTMLCanvasElement) => void;
}

export const useSketchHandWriter = ({ canvasRef, onCanvasUpdate }: Args) => {
  const [isDrawing, setIsDrawing] = useState<boolean>(false);

  useEffect(() => {
    if (canvasRef.current === null) return;

    const ctx = canvasRef.current.getContext("2d");
    if (!ctx) return;

    ctx.clearRect(0, 0, canvasSetting.width, canvasSetting.height);
  }, []);

  const getContext = (): CanvasRenderingContext2D => {
    if (!canvasRef.current) {
      throw new Error("canvasRef is not assigned");
    }

    const ctx = canvasRef.current.getContext("2d");
    if (!ctx) {
      throw new Error("Failed to get 2d context");
    }

    return ctx;
  };

  const handleMouseDown: React.MouseEventHandler = (e) =>  {
    const { offsetX: x, offsetY: y } = e.nativeEvent;
    setIsDrawing(true);
    const ctx = getContext();

    ctx.beginPath();
    ctx.moveTo(x, y);
  };

  const handleMouseMove: React.MouseEventHandler = (e) => {
    if (!isDrawing) return;

    const { offsetX: x ,offsetY: y } = e.nativeEvent;
    const ctx = getContext();

    ctx.lineTo(x, y);
    ctx.stroke();
  }; 

  const handleDrawingFinish = () => {
    setIsDrawing(false);
    if (canvasRef.current === null) return;
    onCanvasUpdate(canvasRef.current);
  };

  return {
    canvasSetting,
    handleMouseDown,
    handleMouseMove,
    handleDrawingFinish,
  };
};
