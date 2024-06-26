import { useRef } from "react";
import { useNavigate } from "react-router-dom";

import { SketchHandWriter } from "./SketchHandWriter";

import { Button } from "~/shared/components/Button";
import { useCanvas } from "~/shared/hooks/useCanvas";
import { sketchApi } from "~/shared/services/API";

export const SketchCreateForm = () => {
  const canvasRef = useRef<HTMLCanvasElement>(null);
  const navigate = useNavigate();
  const { clearCanvas } = useCanvas({ canvasRef });

  const convertCanvasToFile = async (): Promise<File> => {
    const canvas = canvasRef.current;
    if (canvas === null) {
      throw new Error("canvasが存在しません");
    }

    const image = canvas.toDataURL("image/png");

    const blob = await fetch(image).then((res) => res.blob());
    const file = new File([blob], "sketch.png", { type: "image/png" });

    return file;
  };

  const createSketch = async () => {
    const file = await convertCanvasToFile();
    const result = await sketchApi.postSketch(file, {
      withCredentials: true,
    });

    if (result.status !== 201) {
      throw new Error("画像のアップロードに失敗しました");
    }

    navigate("/sketches");
  };
  
  return (
    <div className="flex flex-col gap-8">
      <div>
        <SketchHandWriter canvasRef={canvasRef} />
      </div>

      <div className="flex flex-col gap-2">
        <Button variant="secondary" onClick={clearCanvas}>リセット</Button>
        <Button onClick={createSketch}>投稿する</Button>
      </div>
    </div>
  );
};
