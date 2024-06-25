import React, { useState, useEffect, RefObject } from "react";

const canvasSetting = {
  width: 500,
  height: 300,
};

type Props = {
  canvas: RefObject<HTMLCanvasElement>
  onCanvasUpdate: (canvas: HTMLCanvasElement) => void;
}

export const SketchHandWriter: React.FC<Props> = ({ canvas, onCanvasUpdate }) => {
  const [drawing, setDrawing] = useState<boolean>(false);

  useEffect(() => {
    if (canvas.current === null) return;

    const ctx = canvas.current.getContext("2d");
    if (!ctx) return;

    ctx.clearRect(0, 0, canvasSetting.width, canvasSetting.height);
  }, []);

  const getContext = () => {
    if (!canvas.current) return;
    const ctx = (canvas.current as HTMLCanvasElement).getContext("2d");
    return ctx;
  };

  const mouseDown: React.MouseEventHandler = (e) =>  {
    const { offsetX: x, offsetY: y } = e.nativeEvent;
    setDrawing(true);
    const ctx = getContext();
    if (!ctx) return;

    ctx.beginPath();
    ctx.moveTo(x, y);
  };

  const mouseMove: React.MouseEventHandler = (e) => {
    if (!drawing) return;

    const { offsetX: x ,offsetY: y } = e.nativeEvent;
    const ctx = getContext();
    if (!ctx) return;
    ctx.lineTo(x, y);
    ctx.stroke();
  }; 

  const endDrawing = () => {
    setDrawing(false);
    if (canvas.current === null) return;
    onCanvasUpdate(canvas.current);
  };

  return (
    <canvas
      className="border"
      ref={canvas}
      width={canvasSetting.width}
      height={canvasSetting.height}
      onMouseDown={mouseDown}
      onMouseMove={mouseMove}
      onMouseUp={endDrawing}
      onMouseLeave={endDrawing}
    />
  );
};
