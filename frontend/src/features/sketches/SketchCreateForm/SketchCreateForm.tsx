import { useRef } from "react";

import { useSketchCreateForm } from "./useSketchCreateForm";

import { SketchTool } from "~/features/sketches/SketchTool";
import { SketchHandWriter } from "~/features/sketches/SketchHandWriter";
import { Button } from "~/shared/components/Button";
import { useCanvas } from "~/shared/hooks/useCanvas";

export const SketchCreateForm = () => {
  const canvasRef = useRef<HTMLCanvasElement>(null);
  const { clearCanvas } = useCanvas({ canvasRef });
  const { createSketch } = useSketchCreateForm({ canvasRef });

  return (
    <div className="flex">
      <div className="flex flex-col gap-8">
        <div>
          <SketchHandWriter canvasRef={canvasRef} />
        </div>
        <div className="flex flex-col gap-2">
          <Button variant="secondary" onClick={clearCanvas}>リセット</Button>
          <Button onClick={createSketch}>投稿する</Button>
        </div>
      </div>

      <div className="ml-8">
        <SketchTool />
      </div>
    </div>
  );
};
