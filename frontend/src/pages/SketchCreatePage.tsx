import { useRef } from "react";
import { useNavigate } from "react-router-dom";

import { SketchHandWriter } from "~/features/sketches/SketchHandWriter";
import { Button } from "~/shared/components/Button";
import { Container } from "~/shared/components/Container";
import { useCanvas } from "~/shared/hooks/useCanvas";
import { sketchApi } from "~/shared/services/API";

export const SketchCreatePage = () => {
  const canvasRef = useRef<HTMLCanvasElement>(null);
  const navigate = useNavigate();
  const { clearCanvas } = useCanvas({ canvasRef });

  const postSketchToApi = async () => {
    const canvas = canvasRef.current;
    if (canvas === null) return;
    const image = canvas.toDataURL("image/png");

    const blob = await fetch(image).then((res) => res.blob());
    const file = new File([blob], "sketch.png", { type: "image/png" });
    const result = await sketchApi.postSketch(file, {
      withCredentials: true,
    });

    if (result.status !== 201) {
      throw new Error("画像のアップロードに失敗しました");
    }

    navigate("/sketches");
  };

  return (
    <Container className="flex max-w-[900px] flex-col gap-8 pt-10">
      <div>
        <SketchHandWriter canvasRef={canvasRef} />
      </div>

      <div className="flex flex-col gap-2">
        <Button variant="secondary" onClick={clearCanvas}>リセット</Button>
        <Button onClick={postSketchToApi}>投稿する</Button>
      </div>
    </Container>
  );
};
