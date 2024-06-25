import {
  useState,
  useEffect,
  type RefObject,
  type MouseEventHandler,
} from "react";

const canvasSetting = {
  width: 1020,
  height: 400,
};

type Args = {
  canvasRef: RefObject<HTMLCanvasElement>
}

export const useSketchHandWriter = ({ canvasRef }: Args) => {
  const [isDrawing, setIsDrawing] = useState<boolean>(false);

  useEffect(() => {
    clearCanvas();
  }, []);

  const clearCanvas = () => {
    const ctx = getContext();
    ctx.clearRect(0, 0, canvasSetting.width, canvasSetting.height);
  };

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
