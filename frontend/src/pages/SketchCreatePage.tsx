import axios from "axios";
import { useRef } from "react";

import { SketchHandWriter } from "~/features/sketches/SketchHandWriter";
import { Button } from "~/shared/components/Button";

export const SketchCreatePage = () => {
  const canvasRef = useRef<HTMLCanvasElement>(null);

  const postImageToApi = async () => {
    const canvas = canvasRef.current;
    if (canvas === null) return;
    const image = canvas.toDataURL("image/png");

    // TODO: 画像保存APIに画像を送信する
    await axios.post("http://localhost:9000/images", { image });
  };

  return (
    <div>
      <SketchHandWriter canvasRef={canvasRef} />
      <Button type="submit" onClick={postImageToApi}>APIに送る</Button>
    </div>
  );
};
