import axios from "axios";
import { useRef } from "react";

import { SketchHandWriter } from "~/features/sketches/SketchHandWriter";
import { Button } from "~/shared/components/Button";
import { Container } from "~/shared/components/Container";
import { useCanvas } from "~/shared/hooks/useCanvas";

export const SketchCreatePage = () => {
  const canvasRef = useRef<HTMLCanvasElement>(null);
  const { clearCanvas } = useCanvas({ canvasRef });

  const postSketchToApi = async () => {
    const canvas = canvasRef.current;
    if (canvas === null) return;
    const image = canvas.toDataURL("image/png");

    // TODO: 画像保存APIに画像を送信する
    await axios.post("http://localhost:9000/images", { image });
  };

  return (
    <Container className="flex flex-col gap-8 pt-10">
      <SketchHandWriter canvasRef={canvasRef} />

      <div className="flex flex-col gap-2">
        <Button variant="secondary" onClick={clearCanvas}>リセット</Button>
        <Button onClick={postSketchToApi}>投稿する</Button>
      </div>
    </Container>
  );
};
