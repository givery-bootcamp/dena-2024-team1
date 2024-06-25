import React, { useState, useEffect, RefObject } from "react";

const canvasSetting = {
  width: 500,
  height: 300,
};

type Props = {
  canvasRef: RefObject<HTMLCanvasElement>
  onCanvasUpdate: (canvas: HTMLCanvasElement) => void;
}

export const SketchHandWriter: React.FC<Props> = ({ canvasRef, onCanvasUpdate }) => {
  const [drawing, setDrawing] = useState<boolean>(false);

  useEffect(() => {
    if (canvasRef.current === null) return;

    const ctx = canvasRef.current.getContext("2d");
    if (!ctx) return;

    ctx.clearRect(0, 0, canvasSetting.width, canvasSetting.height);
  }, []);

  const getContext = () => {
    if (!canvasRef.current) return;
    const ctx = (canvasRef.current as HTMLCanvasElement).getContext("2d");
    return ctx;
  };

  const handleMouseDown: React.MouseEventHandler = (e) =>  {
    const { offsetX: x, offsetY: y } = e.nativeEvent;
    setDrawing(true);
    const ctx = getContext();
    if (!ctx) return;

    ctx.beginPath();
    ctx.moveTo(x, y);
  };

  const handleMouseMove: React.MouseEventHandler = (e) => {
    if (!drawing) return;

    const { offsetX: x ,offsetY: y } = e.nativeEvent;
    const ctx = getContext();
    if (!ctx) return;
    ctx.lineTo(x, y);
    ctx.stroke();
  }; 

  const handleDrawingFinish = () => {
    setDrawing(false);
    if (canvasRef.current === null) return;
    onCanvasUpdate(canvasRef.current);
  };

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
