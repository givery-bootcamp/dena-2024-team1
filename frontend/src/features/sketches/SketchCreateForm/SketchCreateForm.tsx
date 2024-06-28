import { useRef } from "react";

import { useSketchCreateForm } from "./useSketchCreateForm";

import { SketchHandWriter } from "~/features/sketches/SketchHandWriter";
import { Button } from "~/shared/components/Button";
import { useCanvas } from "~/shared/hooks/useCanvas";
import { useColorPicker } from "~/shared/hooks/useColorPicker";

export const SketchCreateForm = () => {
  const canvasRef = useRef<HTMLCanvasElement>(null);
  const { clearCanvas } = useCanvas({ canvasRef });
  const { createSketch } = useSketchCreateForm({ canvasRef });
  const { strokeColor, toggleStrokeColor } = useColorPicker();

  return (
    <div className="flex flex-col gap-8">
      <div>
        <SketchHandWriter canvasRef={canvasRef} />
      </div>

      <div className="flex flex-col gap-2">
        <Button variant="secondary" onClick={clearCanvas}>リセット</Button>
        <Button onClick={createSketch}>投稿する</Button>
        <Button onClick={toggleStrokeColor}>
          {strokeColor === "black" ? "赤" : "黒"}色を使う
        </Button>
      </div>
    </div>
  );
};
