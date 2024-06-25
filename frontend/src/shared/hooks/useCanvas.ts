import type { RefObject } from "react";

type Args = {
  canvasRef: RefObject<HTMLCanvasElement>;
}

export const useCanvas = ({ canvasRef }: Args) => {
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

  const clearCanvas = () => {
    const canvas = canvasRef.current;
    if (canvas === null) return;

    const ctx = canvas.getContext("2d");
    if (ctx === null) return;

    ctx.clearRect(0, 0, canvas.width, canvas.height);
  };

  return { getContext, clearCanvas };
};
